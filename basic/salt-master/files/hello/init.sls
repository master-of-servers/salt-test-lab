hello_world:
  file.managed:
    - name: /tmp/hello.txt
    - source: salt://hello/files/hello.txt
    - mode: 755
    - user: root
    - group: root
