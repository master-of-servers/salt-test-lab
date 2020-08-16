FROM ubuntu:18.04

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update && \
      apt-get install -y \
      curl \
      iputils-ping \
      mysql-client \
      python \
      ssh \
      sudo \
      vim && \
      mkdir /var/run/sshd && \
      mkdir /root/.ssh/ && \
      chmod 0755 /var/run/sshd && \
      echo "root:toor" | chpasswd && \
      sed -ri 's/^#?PasswordAuthentication\s+.*/PasswordAuthentication yes/' /etc/ssh/sshd_config && \
      sed -ri 's/^#?PermitRootLogin\s+.*/PermitRootLogin yes/' /etc/ssh/sshd_config && \
      sed -ri 's/UsePAM\syes/#UsePAM yes/g' /etc/ssh/sshd_config && \
      sed -i 's@session\s*required\s*pam_loginuid.so@session optional pam_loginuid.so@g' /etc/pam.d/sshd && \
      curl -L https://bootstrap.saltstack.com -o install_salt.sh && \
      sudo sh install_salt.sh -P && \
      rm -rf /var/lib/apt/lists/

COPY files/setup.sh /opt/setup.sh
COPY files/start.sh /start.sh
RUN chmod +x /start.sh

EXPOSE 22
ENTRYPOINT ["/bin/bash", "/start.sh"]
