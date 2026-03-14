# boo

A CLI for [Ghostty](https://ghostty.org)'s AppleScript API on macOS.

Uses [DarwinKit](https://github.com/progrium/darwinkit) to execute AppleScript.

## Install

```
go install code.selman.me/boo@latest
```

## Usage

```
Usage:
  boo [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  info        Print Ghostty app info
  quit        Quit Ghostty
  session     Manage sessions
  tab         Manage Ghostty tabs
  terminal    Manage Ghostty terminals
  tree        Show hierarchy
  version     Print Ghostty version
  window      Manage Ghostty windows

Flags:
  -h, --help   help for boo
```

```
boo tree                          # show window/tab/terminal hierarchy
boo window list                   # list windows
boo window front                  # get frontmost window
boo tab list --window <id>         # list tabs in a window
boo tab focused-terminal           # focused terminal of selected tab in front window
boo tab focused-terminal --id <id> # focused terminal of a specific tab
boo tab set-title "prod"           # set title of selected tab in front window
boo terminal list                  # list all terminals
boo terminal set-title "logs"      # set title of focused terminal
boo terminal find --cwd ghostty   # find terminals by working directory
boo terminal find --name nvim      # find terminals by name
boo version                        # ghostty version
```

Omit the title argument to clear a tab or terminal title.

### Session save/restore

```
boo session save > session.json
boo session restore < session.json
boo session restore --dry-run < session.json   # preview the generated AppleScript
```

Restore recreates windows, tabs, and splits. Each terminal gets a hint pasted into its prompt showing what was previously running, not executed, just sitting there for reference.


## Limitations

- Ghostty's AppleScript API exposes terminals per tab but not their spatial arrangement, with no position, size, or split tree structure. Restore uses balanced binary space partitioning to approximate the layout.
- Ghostty's AppleScript API does not expose running processes or PIDs. Terminal names often reflect the running command, e.g. `nvim ~/project`. These are saved in the session file and pasted as hints into restored terminals so you know what to restart.
- Requires Ghostty 1.3.0+; AppleScript support was introduced in that release.
