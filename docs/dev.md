# Developer Documentation

The rest of this document is for people that want to be able to hack on this project.
If that's not you, please feel free to ignore this entire document.

## Table of Contents

- [Dependencies](#dependencies)
- [Developer Environment Setup](#developer-environment-setup)

---

## Dependencies

- [Install pre-commit](https://pre-commit.com/):

  ```bash
  brew install pre-commit
  ```

- Set up `go.mod` for development:

  ```bash
  REPO=github.com/master-of-puppets/salt-test-lab
  FORK="${PWD}"

  echo -e "\nreplace ${REPO} => ${FORK}" >> go.mod
  ```

---

## Developer Environment Setup

1. Run the following command to set up `go.mod` for development with your fork:

   ```bash
   REPO=github.com/master-of-puppets/salt-test-lab
   FORK="${PWD}"

   echo -e "\nreplace ${REPO} => ${FORK}" >> go.mod
   ```

2. Install pre-commit hooks locally

   ```bash
   # If your host OS is Linux
   salt-test-lab-linux-amd64 preCommit
   # If your host OS is MacOS
   salt-test-lab-darwin-amd64 preCommit
   ```

3. Run pre-commit hooks locally

   ```bash
   pre-commit run --all-files
   ```
