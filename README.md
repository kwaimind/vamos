# vamos ⚡

> Run the package manager you need, when you need it

`vamos` is a simple tool to run the correct package manager when jumping between projects or in monorepo environments.

It’s a Go script, that looks for the closest `package.json` file, and relies on the `engines` schema to pick the right package for the current project

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

`vamos -F frontend test`

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

`-F <package-name>` - Filter to a specific package in a workspace/monorepo

When working in a monorepo, use `-F` to target a specific package by name. vamos will find the matching package's `package.json` and use its configured package manager.

Example:
```bash
vamos -F frontend test
```

This finds the package named "frontend" and runs the test script using that package's configured manager (e.g., pnpm, npm, yarn).

## install

right now you need to build the script locally with `make build`. i'm planning on moving this to homebrew later.