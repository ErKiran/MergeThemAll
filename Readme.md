# Merge 'em All

![MergeThemAll](https://github.com/ErKiran/MergeThemAll/blob/master/mergethemall.png)

### Merge all the Pending PR 

* You need to install `gh` cli and authenticate it

Using the power of Github cli `gh` to merge all the pending pr in user workspace.

Commands Used:  
* `gh repo list --limit=n --json=name,owner`

* `gh pr list --repo=reponame --json=number,author`

* `gh pr merge n -m --repo=reponame`

Merging all the pending PR can be destructive. I have used it to merge all the PR opened by [@dependabot](https://github.com/dependabot) only.

> Dependabot helps you keep your dependencies up to date. Every day, it checks your dependency files for outdated requirements and opens individual PRs for any it finds. You review, merge, and get to work on the latest, most secure releases. Dependabot is a tool in the Dependency Monitoring category of a tech stack.