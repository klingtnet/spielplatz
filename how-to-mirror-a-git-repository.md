# How to mirror a git repository

Most git providers allow to mirror a remote repository but the biggest one of all—github.com—does not.
So, how do you mirror a repository to github then?
I think the easiest and most straightforward way is to specify a git-hook to do exactly this.
First specify both remotes, in my chase those are `git remote add <origin> <repo-URL>` and `git remote add <mirror> <repo-URL>`.
Now create the hook: `$ echo 'git push --force --mirror <mirror>' > .git/hooks/post-commit'.
The next time you push your changes into `<origin>` the changes will be automatically mirrored into `<mirror>`. Voila!
