FROM ubuntu:18.04

RUN apt-get update && \
      apt-get install -y curl \
      sudo \
      vim && \
      curl -L https://bootstrap.saltstack.com -o install_salt.sh && \
      sudo sh install_salt.sh -P

COPY files/setup.sh /opt/setup.sh
ENTRYPOINT ["sh","/opt/setup.sh"]
