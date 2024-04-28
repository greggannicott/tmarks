# tmarks: Bookmarks, for tmux

## IMPORTANT

This is still a work in progress with no official first version. Frankly I'm amazed you've found it!

Use it at your own risk.

## About

`tmarks` allows you to create bookmarks for tmux sessions. Think [harpoon](https://github.com/ThePrimeagen/harpoon), but dedicated to tmux!

If you are working with a specific set of sessions, rather than having to choose a session from a potentially large number using `choose-session`, select from a subset of sessions using `tmarks`.

You choose which sessions do and do not appear in the list.

## Usage

### Adding Current Session

To add your current session to the list, run:

```bash
tmarks add
```

### Opening a Session

To display a list of saved sessions, run:

```bash
tmarks list
```

Navigate to the session you're after and press Enter to open it.
