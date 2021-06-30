```
 _   _ _
| \ | (_)_ __   __ _
|  \| | | '_ \ / _` | - A cli for the NOKO Time Tracking Software
| |\  | | | | | (_| |
|_| \_|_|_| |_|\__,_|
```

## Download

[Latest binaries and packages](https://github.com/Raphexion/nina/releases/latest)

## Configure

Please create a personal api token in Noko (web page).

Either create an environmental variable

```
export NOKO_API_KEY="my-key-1234"
```

Or, create nina.yaml in your home folder:

```
# ~/nina.yaml
api_key: my-key-1234
```

## Usage

### Projects

List all projects:

```
nina projects list
```

### Timers

List all timers:

```
nina timers list
```

Pause the currently running timer:

```
nina timer pause
```

Start a paused timer:

```
nina timer start my-cool-project
```

Todo:

- [ ] Create a timer and start it
