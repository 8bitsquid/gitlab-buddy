# Gitlab Buddy

This currently has limited functionality, but was designed so that development of new features
can be made with relative ease.

> :warning: Only branch migration is currently supported. See [Features](#features-implemented-and-planned) for details.


## Getting Started

Gitlab Buddy requires `API Tokens` be stored in a config file. Each token is treated as a separate `client`. Multiple tokens, with different permissions, can be
configured for the same Gitlab host. `Clients` registered in Gitlab Buddy will only _read_ from eachother, they will never _write_ to eachother. This lets you use of tokens with write access to a minimum, and prevents potential cross-chatter.

### 1. Download the binary

[Click here](#) to download the latest binary.

### 2. Run the config wizard

The config wizard is very minimal for the time being. More features are planned, but you can always make your own config file [see here](#the-config-file)

```bash
# From your command terminal and go to the directory with the 'gitlab-biddy' binary
$ cd /path/to/dir

# Make sure it's executable
$ chmod u+x gitlab-buddy

# Run the config wizard 
# (the "./" is required, since 'gitlab-buddy' isn't in you $PATH env variable)
$ ./gitlab-buddy config init
```

### 3. Check available features

Use `--help` flag for details on `gitlab-buddy` commands
```bash
# General help info
$ ./gitlab-buddy --help

# Migrate Branch feature info
$ ./gitlab-buddy migrate branch --help
```

## The Config file

The config file is required, and stores your Gitlab Tokens. This prevents API Tokens from ending up
in your terminal history. Remote hosts (e.g., https://gitlab.com) require a _human friendly_ name associated with them.
This name is used to identify a host when executing `gitlab-buddy` commands. 

*The config file must have a `yml` or `yaml` extension*

```yaml
# ~/my-glb-config.yml

hosts:
    myHostName1:
        token: mY$Up3rs3cre770k3N
        url: https://myGitlabHost.com # Optional (default: https://gitlab.com)
        apiPath: /path/to/api/version  # Optional (default: /api/v4/)
    myHostName2:
        token: h4cK7H3g18$0n
```

With the config above, I would run commands for `myhostName1` like so:
```bash
$ ./gitlab-buddy migrate branch master main --host myHostName1
```

## Features: Implemented and Planned

- [x] Migration
    - [x] Git branches (e.g., `[master]` -> `[main]`)
        - [x] Migrate all `projects` in a `group`
        - [x] Migrate individual `projects`
            - [x] `Merge Requests` will target new branch
            - [x] Archives old branch in a protected git `tag` (e.g., `archived-glb-master-to-main`)
            - [x] Checks project files for references to old branch
                - Looks for string variable declarations that match the old branch (e.g., `"oldBranchName"` and `'oldBranchName'`)
                - :warning: Project will be skipped if detected
            - [ ] Migrate `submodules`
                - :warning: Not yet supported. Projects with `submodules` are skipped during migration

    - [ ] Projects
    - [ ] Groups
- [ ] Access & deploy tokens
- [ ] Membership management
- [ ] Local clone sync
  - [ ] Migrations
