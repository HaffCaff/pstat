# pstat

![Screenshot 2024-02-08 at 11 30 04 PM](https://github.com/HaffCaff/pstat/assets/31331413/15b02e82-3b7f-492c-8e90-b84afb17118d)


`pstat` (project status) is a simple Go cli tool that will iterate through a directory and let you know which project are stale, and which projects are active. This is determined by directory mod time (date last modified). Currently, if it the project mod date is more than 30 days old, it will be labeled as stale.

The inspiration for this came from opening my dev directory on a teams call to look for something that I knew I worked on recently, but there were so many files I got lost...on camera 😭

## Future Improvements:

- Allow customization - Allow user to change modTime date threshold
- Prompt to stale projects to new `stale` directory within current dir
- Make public and allow for contributions!
- Publish to homebrew
- add CI/CD pipeline implementation
