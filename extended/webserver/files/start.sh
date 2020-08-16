set -x
service nginx start
service php7.2-fpm start
/bin/bash /opt/setup.sh &
/usr/sbin/sshd -D