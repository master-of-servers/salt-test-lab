# salt-test-lab

[![License](http://img.shields.io/:license-mit-blue.svg)](https://github.com/master-of-servers/salt-test-lab/blob/master/LICENSE)
[![Run Pre-Commit hooks](https://github.com/master-of-servers/salt-test-lab/actions/workflows/pre-commit.yaml/badge.svg)](https://github.com/master-of-servers/salt-test-lab/actions/workflows/pre-commit.yaml)
[![Test salt-test-lab](https://github.com/master-of-servers/salt-test-lab/actions/workflows/test-build.yaml/badge.svg)](https://github.com/master-of-servers/salt-test-lab/actions/workflows/test-build.yaml)
[![Release builder for salt-test-lab](https://github.com/master-of-servers/salt-test-lab/actions/workflows/release.yaml/badge.svg)](https://github.com/master-of-servers/salt-test-lab/actions/workflows/release.yaml)

Used to create a test lab that can be used to work with MOSE and Salt.

**Warning, take heed:**
This lab should be run in a controlled environment; it contains vulnerable assets.

## Table of Contents

- [Dependencies](#dependencies)
- [Basic Lab Build Instructions](#basic-lab-build-instructions)
- [Extended Lab Build Instructions](#extended-lab-build-instructions)
- [Credits](#credits)
- [Developer Documentation](docs/dev.md)

---

## Dependencies

You must download and install the following for this environment to work:

- [Docker](https://docs.docker.com/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Latest release of the salt-test-lab binary](https://github.com/master-of-servers/salt-test-lab/releases)

## Basic Lab Build Instructions

To create an environment with a salt master that controls a single minion
with a simple hello world module, run the following command:

```bash
# If your host OS is Linux
salt-test-lab-linux-amd64 buildBasic
# If your host OS is MacOS
salt-test-lab-darwin-amd64 buildBasic
```

To tear down the lab, run the following command:

```bash
# If your host OS is Linux
salt-test-lab-linux-amd64 destroyBasic
# If your host OS is MacOS
salt-test-lab-darwin-amd64 destroyBasic
```

## Extended Lab Build Instructions

To create an environment with a salt master that controls multiple minions with
multiple environments and includes a database, web server, and development system:

```bash
# If your host OS is Linux
salt-test-lab-linux-amd64 buildExtended
# If your host OS is MacOS
salt-test-lab-darwin-amd64 buildExtended
```

**Please note:**
The docker labs do not support transferring payloads to the target via the web server.
You can however generate a payload with the `-f` parameter and transfer
it via `docker cp`.

To tear down the lab, run the following command:

```bash
# If your host OS is Linux
salt-test-lab-linux-amd64 destroyExtended
# If your host OS is MacOS
salt-test-lab-darwin-amd64 destroyExtended
```

## Credits

[This article](https://medium.com/@jmarhee/creating-a-saltstack-test-environment-with-docker-compose-7dfa79837712)
was a critical component in standing up this lab.
