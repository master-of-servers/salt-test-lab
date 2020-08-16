FROM ubuntu:18.04

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update && \
      apt-get install -y \
      curl \
      mysql-client \
      nginx \
      php7.2 php7.2-fpm php7.2-cli php7.2-curl php7.2-mysql php7.2-curl php7.2-gd php7.2-mbstring php-pear \
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
      sed -i 's/listen = \/run\/php\/php.*.sock/listen = localhost:8080/' /etc/php/7.2/fpm/pool.d/www.conf && \
      curl -L https://bootstrap.saltstack.com -o install_salt.sh && \
      sudo sh install_salt.sh -P && \
      rm -rf /var/lib/apt/lists/

RUN cd /var/www/html/ && \
      # Download pages from bsidesboulder.org for us to host
      wget \
      --mirror            \
     --convert-links     \
     --html-extension    \
     --wait=2            \
     -o log              \
      https://bsidesboulder.org/home; exit 0

RUN cd /var/www/html/ && \
      rm index.nginx-debian.html && \
      mv bsidesboulder.org/ home/

COPY files/default /etc/nginx/sites-available/default
COPY files/nginx.conf /etc/nginx/nginx.conf
COPY files/setup.sh /opt/setup.sh
COPY files/index.html /var/www/html
COPY files/start.sh /start.sh
RUN chmod +x /start.sh

EXPOSE 22
ENTRYPOINT ["/bin/bash", "/start.sh"]