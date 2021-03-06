## kpt pkg sync

Fetch and update packages declaratively

<link rel="stylesheet" type="text/css" href="/kpt/gifs/asciinema-player.css" />
<asciinema-player src="/kpt/gifs/pkg-sync.cast" speed="1" theme="solarized-dark" cols="100" rows="26" font-size="medium" idle-time-limit="1"></asciinema-player>
<script src="/kpt/gifs/asciinema-player.js"></script>

    # run the tutorial from the cli
    kpt tutorial pkg sync

[tutorial-script]

### Synopsis

Sync uses a manifest to manage a collection of dependencies.

The manifest declares *all* direct dependencies of a package in a Kptfile.
When `sync` is run, it will ensure each dependency has been fetched at the
specified ref.

This is an alternative to managing package dependencies individually using
the `get` and `update` commands.

    kpt pkg sync LOCAL_PKG_DIR [flags]

  LOCAL_PKG_DIR:

    Local package with dependencies to sync.  Directory must exist and contain a Kptfile.

#### Env Vars

  KPT_CACHE_DIR:

    Controls where to cache remote packages during updates.
    Defaults to ~/.kpt/repos/

#### Dependencies

For each dependency in the Kptfile, `sync` will ensure that it exists locally with the
matching repo and ref.

Dependencies are specified in the `Kptfile` `dependencies` field and can be added or updated
with `kpt pkg sync set`.  e.g.

    kpt pkg sync set https://github.com/GoogleContainerTools/kpt.git/package-examples/helloworld-set \
        hello-world

Or edit the Kptfile directly:

    apiVersion: kpt.dev/v1alpha1
    kind: Kptfile
    dependencies:
    - name: hello-world
      git:
        repo: "https://github.com/GoogleContainerTools/kpt.git"
        directory: "/package-examples/helloworld-set"
        ref: "master"

Dependencies have following schema:

    name: <local path (relative to the Kptfile) to fetch the dependency to>
    git:
      repo: <git repository>
      directory: <sub-directory under the git repository>
      ref: <git reference -- e.g. tag, branch, commit, etc>
    updateStrategy: <strategy to use when updating the dependency -- see kpt help update for more details>
    ensureNotExists: <remove the dependency, mutually exclusive with git>

Dependencies maybe be updated by updating their `git.ref` field and running `kpt pkg sync`
against the directory.

### Examples

  Example Kptfile to sync:

    # file: my-package-dir/Kptfile

    apiVersion: kpt.dev/v1alpha1
    kind: Kptfile
    # list of dependencies to sync
    dependencies:
    - name: local/destination/dir
      git:
        # repo is the git respository
        repo: "https://github.com/pwittrock/examples"
        # directory is the git subdirectory
        directory: "staging/cockroachdb"
        # ref is the ref to fetch
        ref: "v1.0.0"
    - name: local/destination/dir1
      git:
        repo: "https://github.com/pwittrock/examples"
        directory: "staging/javaee"
        ref: "v1.0.0"
      # set the strategy for applying package updates
      updateStrategy: "resource-merge"
    - name: app2
      path: local/destination/dir2
      # declaratively delete this dependency
      ensureNotExists: true

  Example invocation:

    # print the dependencies that would be modified
    kpt pkg sync my-package-dir/ --dry-run

    # sync the dependencies
    kpt pkg sync my-package-dir/

[tutorial-script]: ../gifs/pkg-sync.sh
