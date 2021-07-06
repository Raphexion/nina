```
 _   _ _
| \ | (_)_ __   __ _
|  \| | | '_ \ / _` |
| |\  | | | | | (_| | - A cli for the Noko Time Tracking Software
|_| \_|_|_| |_|\__,_|
```

[![Go Report Card](https://goreportcard.com/badge/github.com/Raphexion/nina)](https://goreportcard.com/report/github.com/Raphexion/nina)

## Noko (previously Freckle)

[Noko](https://nokotime.com/) is a time tracking software tool.
Nina is a command-line tool to directly interact with Noko through the Noko API.

In Noko, each project has a timer per user.
Which means that if your organization has two projects: Sales and Development and three employees: Anna, Bengt and Carolina.
Then there will be six potential timers.

That means, that if you are Anna. When you want to start/pause/log a timer, you actually only the project name.
This timer will not conflict with Bengt's and Carolina's timers for the same project.

## Download

[Latest binaries and packages](https://github.com/Raphexion/nina/releases/latest)

## Configure

Please create a Personal Access Token in Noko (web page).

Either create an environmental variable

```
export NOKO_ACCESS_TOKEN="my-key-1234"
```

Or, create nina.yaml in your home folder:

```
# ~/nina.yaml
access_token: my-key-1234
```

