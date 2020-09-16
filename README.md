# Zi

_A key value database._

Zi is a key value database originaly designed for storing coins and id's for Vitecoin. It can also be used for general storage with key value.

Zi should be put on a server or a local machine and in a directory with a file called `dump.zi`.

## Install

First things first your going to want to clone the repository.

```bash
git clone https://github.com/Vitecoin/Zi.git
```

Then all of install process is taken care of in the install file. You can just run:

```bash
chmod +x install
./install
```

This will install the CLI and get all of the dependencies.

## Getting started

The whole server and the database can be started with one simple command.

```bash
    zi serve #by default the server will run on port 9090
```

You can specify a port with the following.

```bash
zi serve 5000
```

This will serve the server on port 5000.

You can add the `--background` flag to run the server as a daemon proccess.

```bash
zi serve 3000 --background
```
