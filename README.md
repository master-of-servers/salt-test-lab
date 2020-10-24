# salt-test-lab
[![License](http://img.shields.io/:license-mit-blue.svg)](https://github.com/master-of-servers/salt-test-lab/blob/master/LICENSE)
[![Build Status](https://dev.azure.com/jaysonegrace/salt-test-lab/_apis/build/status/master-of-servers.salt-test-lab?branchName=master)](https://dev.azure.com/jaysonegrace/salt-test-lab/_build/latest?definitionId=35&branchName=master)

Used to create a test lab that can be used to work with MOSE and Salt.

**Warning, take heed: This lab should be run in a controlled environment; it contains vulnerable assets.**

## Dependencies
You must download and install the following for this environment to work:
* [Docker](https://docs.docker.com/install/)
* [Docker Compose](https://docs.docker.com/compose/install/)

## Basic Lab Build Instructions
To create an environment with a salt master that controls a single minion with a simple hello world module, run the following command:
```
cd basic && make run
```

To tear down the lab, run the following command:
```
make destroy
```

## Extended Lab Build Instructions
To create an environment with a salt master that controls multiple minions with multiple environments and includes a database, web server, and development system, run the following command:
```
cd extended && make run
```

**Please note:**
The docker labs do not support transferring payloads to the target via the web server. You can however generate a payload with the `-f` parameter and transfer it via `docker cp`.

To tear down the lab, run the following command:
```
make destroy
```

## Credits
This article was a critical component in standing up this lab:
https://medium.com/@jmarhee/creating-a-saltstack-test-environment-with-docker-compose-7dfa79837712