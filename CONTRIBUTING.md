# Contributing Guidelines :boom:

Firstly, thank you for considering contributing to `grole`!

When developing, please follow the below guidelines.

## Devcontainer :zap:

A devcontainer is provided for development.

This is intended to be used with VSCode, and should be used to develop
locally.

You will just need the `ms-vscode-remote.remote-containers` VSCode extension
before opening the `grole` directory. Once, you have this, VSCode will automatically
ask if you would like to build the Devcontainer image and reopen the directory
in a container.

Developing in here will ensure your code passes some required checks before opening
a PR.

## Branches & PRs :deciduous_tree:

Please use the following convention when creating a new branch:

`fb/issue-number-brief-description`

Of course, you can leave the issue number out if it isn't applicable.

When raising a PR, please be descriptive in both the Title and Description.

If the PR closes an issue, please indicate this in the PR title with
`Closes #issuenumber`.

You should add a reviewer to each PR you raise, which can be `jay-kinder` for
now.

## Workflows :sunglasses:

There are some GitHub Workflows which run on the raising of every PR, and also
when changes are merged to the `main` branch.

On PR:

1. `pre-commit`: this runs a series of checks against the pushed code, to ensure
that it adheres to the guildelines put in place.
2. `test-coverage`: this checks that test coverage remains over 90%. If you have
coded a new feature, you should ensure there is a test(s) covering this.

On push to `main` branch:

1. `update-cache`: always runs, regardless of what is being merged. Used to update
the cache used in the `pre-commit` workflow.
2. `create-push-tag`: only runs if `*.go` file(s) have been merged. Will create and
push a new tag.
3. `create-release`: if a new tag (`v*`) is uploaded, this will create a new release.
4. `release`: this is the release pipeline. When a new release is created, this
workflow will release a new version of the code.

**Tags and Releases will be managed by Maintainers.**

## Contacts :postbox:

Maintainer: `jay-kinder`

## Thank You :star:

Once again, thank you for taking the time to contribute to `grole`.

Have fun!!
