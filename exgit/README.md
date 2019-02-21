EXGIT v1.0.1

A script to do some git commands on several repositories.
<hr />

```
$ exgit --help
-s      git status
-P      git pull
-c      git-clean
-b      git branch
```
<hr />

`exgit [option] [path]`

Option: please do `--help` to get the options list.

Path: the path of the directory where there are repositories.
<hr />

`git-clean` is an alias for:

`git remote prune origin && git branch -vv | grep \"origin/.*: gone]\" | awk \"{print }\" | xargs git branch -D 2>/dev/null`

This command checks and removes all deleted branches
