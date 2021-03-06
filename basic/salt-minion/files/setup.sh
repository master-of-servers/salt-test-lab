#!/bin/bash
salt_minion_check () {
  if [ ! -f /var/log/salt/minion ]; then
    echo "File not found!" && \
      touch /var/log/salt/minion && \
      salt_minion_check
        else
          tail -f /var/log/salt/minion
  fi
}

echo "master: salt.demo.com" >> /etc/salt/minion && \
  echo "id: $(hostname)" >> /etc/salt/minion
  service salt-minion start && \
    salt_minion_check
