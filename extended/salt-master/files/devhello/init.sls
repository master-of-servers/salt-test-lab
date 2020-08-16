hello_dev:
  file.managed:
    - name: /tmp/hellodev.txt
    - source: salt://devhello/files/devhello.txt
    - mode: 755
    - user: root
    - group: root
