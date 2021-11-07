FROM ubuntu:18.04

RUN apt-get update && \
      apt-get install -y curl \
      vim && \
      curl -L https://bootstrap.saltstack.com -o install_salt.sh && \
      sh install_salt.sh -P

COPY files/setup.sh /opt/setup.sh
ENTRYPOINT ["sh","/opt/setup.sh"]
