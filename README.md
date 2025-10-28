# vamos ⚡

> Run the package manager you need, when you need it

`vamos` is a simple tool to run the correct package manager when jumping between projects or in monorepo environments.

It looks for the closest `package.json` file and detects the package manager using:
1. The `engines` field in package.json (preferred)
2. Lockfile detection (package-lock.json, pnpm-lock.yaml, yarn.lock, bun.lockb)
3. Falls back to npm if neither is found

**Supported package managers:** npm, pnpm, yarn, bun

### using vamos

`vamos test`

```json
{
  "name": "my-project",
  "engines": {
    "node": "20.18.3",
    "npm": "10.9.2"
  },
  "scripts": {
    "test": "echo 'running npm'"
  }
}
```

in this example, `vamos` will take your command line arguments and pass them to `npm`.

this also works in a workspace environment by running vamos from the root of the workspace

`vamos -f frontend test`

```json
// ./app/frontend
{
  "name": "frontend",
  "engines": {
    "node": "20.18.3",
    "pnpm": ">=8.9.0"
  },
  "scripts": {
    "test": "echo 'running pnpm'"
  }
}
```

now your command arguments will be called using `pnpm`.

## arguments

### `-f, --filter <package-name>`
Filter to a specific package in a workspace/monorepo. vamos will find the matching package's `package.json` and use its configured package manager.

Example:
```bash
vamos -f frontend test
vamos --filter frontend test
```

### `--verbose`
Show which package manager was selected and why. Useful for debugging or understanding package manager detection.

Example:
```bash
vamos --verbose test
# Output: 🚀 Using pnpm (detected from engines.pnpm in package.json)
```

### `--help, -h`
Display usage information and available options.

### `--version, -v`
Display the current version of vamos.

## install

right now you need to build the script locally with `make build`. i'm planning on moving this to homebrew later.