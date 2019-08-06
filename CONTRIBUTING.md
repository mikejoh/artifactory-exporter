# Contributing

You are more than welcome to contribute with issues and PRs!

### Quick how-to for new GitHub users
If you're unsure on how to contribute, with let's say a simple pull request (e.g. to fix a typo) using the GitHub workflow then here's a quick guide on how you can get started. You need to have `git` installed to proceed:

1. Fork this repository in the GitHub UI, click Fork in the right corner in this repository.
2. Clone your fork to your local computer: `git clone <your fork url here>`
3. Make sure you have `user.name` and `user.email` configured in your `~/.gitconfig`:
```
git config --global user.name "Your name here"
git config --global user.email "Your email address here"
```
4. Checkout a new branch with a descriptive name to fix that typo: `git checkout -b fix-typo`
5. Fix the typo(s) in the file(s).
6. Run `git add .` to stage your changes.
7. Run `git commit -m "Fix typo(s)"` to commit your changes.
8. Run `git push --set-origin origin fix-typo` to push the changes to your fork in the new branch.
9. Create an pull request to merge your fork branch with master in this repository.
10. You're now done and i can do a review of the pull request before it's merged.

This guide covers a very basic scenario, there's alot more to this work flow that can be explained and added but let's keep it simple. If your unsure, always create an issue first!