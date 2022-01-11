import * as React from 'react';
import { useTranslation } from 'react-i18next';
import { Dropdown, DropdownToggle, DropdownItem } from '@patternfly/react-core';
import { ThIcon } from '@patternfly/react-icons';
import * as _ from 'lodash';

type Props = {
  setSize: (v: string) => void;
  id?: string;
};

export const DisplayDropdown: React.FC<Props> = ({ id, setSize }) => {
  const [isOpen, setIsOpen] = React.useState<boolean>(false);
  const { t } = useTranslation('plugin__network-observability-plugin');

  const sizeOptions = {
    s: t('Compact'),
    m: t('Normal'),
    l: t('Large')
  };

  return (
    <Dropdown
      id={id}
      isPlain
      dropdownItems={_.map(sizeOptions, (name, key) => (
        <DropdownItem id={key} component="button" key={key} onClick={() => setSize(key)}>
          {name}
        </DropdownItem>
      ))}
      isOpen={isOpen}
      onSelect={() => setIsOpen(false)}
      toggle={
        <DropdownToggle id={`${id}-dropdown`} toggleIndicator={null} onToggle={() => setIsOpen(!isOpen)}>
          <ThIcon />
        </DropdownToggle>
      }
    />
  );
};

export default DisplayDropdown;
