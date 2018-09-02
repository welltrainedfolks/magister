# MAGISTER

MAGISTER stands for:

**M**AGISTER (is an) **A**dvanced **G**olang **I**mport **S**erver **T**hat **E**nforces right **R**outing for packages.

In other words, MAGISTER replies to ``go`` binary (or your package manager) where it should go to obtain sources.

## Abilities

MAGISTER is able to:

* Be a HTTP server.
* Show easy to use web interface which able to:
  * Login/logout administrators.
  * Control which packages are served.

### ToDo

* Full configuration thru web interface.
* Packages mirrors round-robin.
* ...maybe more :)

## Installation

*TBW*

## Configuration

*TBW*

## Usage

### Controlling users

There is separate application that does (for now only) users controlling called ``magisterctl``. This should be enough to get started:

```bash
go get -u -v github.com/welltrainedfolks/magister/cmd/magisterctl/...
```

Read help (``magisterctl -h``) to understand how to use it.