package handler

import (
	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v3"
)

type QuickFilter struct {
	Name    string            `yaml:"name,omitempty" json:"name"`
	Filter  map[string]string `yaml:"filter,omitempty" json:"filter"`
	Default bool              `yaml:"default,omitempty" json:"default"`
}

type frontendConfig struct {
	PortNaming struct {
		Enable    bool              `yaml:"enable,omitempty" json:"enable"`
		PortNames map[string]string `yaml:"portNames,omitempty" json:"portNames"`
	} `yaml:"portNaming,omitempty" json:"portNaming"`
	QuickFilters []QuickFilter `yaml:"quickFilters" json:"quickFilters"`
}

func readConfigFile(filename string) (*frontendConfig, error) {
	cfg := frontendConfig{
		QuickFilters: []QuickFilter{},
	}
	if len(filename) == 0 {
		return &cfg, nil
	}
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &cfg)
	return &cfg, err
}

func GetConfig(filename string) func(w http.ResponseWriter, r *http.Request) {
	resp, err := readConfigFile(filename)
	if err != nil {
		hlog.Errorf("Could not read config file: %v", err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			resp, err = readConfigFile(filename)
			if err != nil {
				writeError(w, http.StatusInternalServerError, err.Error())
			} else {
				writeJSON(w, http.StatusOK, resp)
			}
		} else {
			writeJSON(w, http.StatusOK, resp)
		}
	}
}
