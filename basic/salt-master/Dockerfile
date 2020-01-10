FROM ubuntu:18.04

# Setup master
RUN apt-get update && \
      apt-get install -y curl \
      sudo \
      tree \
      vim && \
      mkdir /srv/salt && \
      curl -L https://bootstrap.saltstack.com -o install_salt.sh && \
      sudo sh install_salt.sh -P -M

# Generate secret
COPY genkey-batch /srv/pillar/genkey-batch
COPY data.sls /srv/pillar/data.sls

RUN mkdir -p /srv/pillar && \
      mkdir -p /etc/salt/gpgkeys && \
      chmod 0700 /etc/salt/gpgkeys && \
      gpg --batch --gen-key --homedir /etc/salt/gpgkeys /srv/pillar/genkey-batch && \
      gpg --homedir /etc/salt/gpgkeys --armor --export "Salt Key" > /etc/salt/gpgkeys/exported_pubkey.gpg && \
      gpg --import /etc/salt/gpgkeys/exported_pubkey.gpg && \
      echo -n "supersecret" | gpg --armor --batch --trust-model always --encrypt -r "Salt Key" | sed -E 's/^/  /' >> /srv/pillar/data.sls

COPY hello /srv/salt/hello
COPY top.sls /srv/salt/top.sls
COPY pillartop.sls /srv/pillar/top.sls
COPY setup.sh /opt/setup.sh
ENTRYPOINT ["sh","/opt/setup.sh"]
