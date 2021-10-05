```
 _   _ _
| \ | (_)_ __   __ _
|  \| | | '_ \ / _` |
| |\  | | | | | (_| | - A cli for the Noko Time Tracking Software
|_| \_|_|_| |_|\__,_|
```

[![Go Report Card](https://goreportcard.com/badge/github.com/Raphexion/nina)](https://goreportcard.com/report/github.com/Raphexion/nina)
[![codecov.io](https://codecov.io/gh/Raphexion/nina/coverage.svg?branch=master)](https://codecov.io/gh/Raphexion/nina?branch=master)

## Noko (previously Freckle)

[Noko](https://nokotime.com/) is a time tracking software tool.
Nina is a command-line tool to directly interact with Noko through the Noko API.

In Noko, each project has a timer. However, it is tied to specific user.
For example, if your organization has two projects: Sales and R&D and three employees: Anna, Bengt and Carolina.
Then are six timers.

| Person   | Project |
|----------|---------|
| Anna     | Sales   |
| Anna     | R&D     |
| Bengt    | Sales   |
| Bengt    | R&D     |
| Carolina | Sales   |
| Carolina | R&D     |


That means, that if you are Anna. When you want to start/pause/log a timer, you actually only need the project name.
This timer will not conflict with Bengt's and Carolina's timers for the same project.

## Download

[Latest binaries and packages](https://github.com/Raphexion/nina/releases/latest)

## Configure

Please create a Personal Access Token in Noko (web page).
On the right hand side, open *Connected Apps & API*.
Then look for *NOKO API* and *Personal Access Tokens*.

Either create an environmental variable

```
export NOKO_ACCESS_TOKEN="my-key-1234"
```

Or, create nina.yaml in your home folder:

```
# ~/nina.yaml
access_token: my-key-1234
```

## Usage

Please use the internal help messages at the command-line for an up-to-date usage. Example:

```sh
$ nina
A commandline client written in golang to help interact with Noko time tracker

Usage:
  nina [command]

Available Commands:
  entries
  help        Help about any command
  projects
  timers

Flags:
      --config string   config file (default is $HOME/nina.yaml)
  -h, --help            help for nina

Use "nina [command] --help" for more information about a command.
```
