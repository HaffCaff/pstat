# pstat

`pstat` is a simple cli tool that will iterate through a directory and let you know which project are stale, and which projects are active. This is determined by directory mod time (date last modified). Currently, if it the project mod date is more than 30 days old, it will be labeled as stale.

## Future Improvements:

- Allow customization - Allow user to change modTime date threshold
- Prompt to stale projects to new `stale` directory within current dir
