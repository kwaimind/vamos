# vamos improvement plan

## Phase 1: Critical bug fixes ✅

**Goal**: Fix bugs that break core functionality

- [x] Fix stderr redirect in main.go:65 (uses `os.Stdin` instead of `os.Stderr`)
- [x] Fix arg filtering - remove `-f` flag and value from args before passing to package manager
- [x] Fix ignore logic in SetupIgnore - check paths relative to project, not current dir

**Success criteria**: Tool works correctly in all documented scenarios

---

## Phase 2: Test coverage ✅

**Goal**: Ensure existing code is properly tested

- [x] Add tests for select.go (package manager selection logic)
- [x] Add tests for strings.go (CapitalizeFirst)
- [x] Add tests for ignore.go (SetupIgnore function)
- [x] Add integration test for main flow
- [x] Increase coverage to 62.9% (70% target requires refactoring main())

**Success criteria**: All core functions tested ✅
**Note**: Coverage is 62.9% (up from 30.3%). The 70% target is difficult to achieve because main() (76 lines, 0% coverage) is not easily testable without refactoring. All business logic functions have 87%+ coverage.

---

## Phase 3: Core features ✅

**Goal**: Add missing package manager support and smarter detection

- [x] Add bun support to Engines struct and Select function
- [x] Add lockfile detection fallback
  - Detect from package-lock.json → npm
  - Detect from pnpm-lock.yaml → pnpm
  - Detect from yarn.lock → yarn
  - Detect from bun.lockb → bun
- [x] Use lockfile detection when engines field is missing

**Success criteria**: Works with all 4 major package managers (npm, pnpm, yarn, bun) ✅
**Note**: Coverage increased to 66.1% (up from 62.9%)

---

## Phase 4: Code quality ✅

**Goal**: Improve maintainability and reduce technical debt

- [x] Replace panic in ParseArgs with proper error return
- [x] Standardize error handling across all functions
- [x] Return errors from SetupIgnore instead of printing
- [x] Pass config as parameter instead of reinitializing (select.go:6, dir.go:14)
- [x] Move style out of global scope in main.go
- [x] Consistent naming (packageJSON vs PackagejsonName)
- [x] Consider using fs.WalkDir instead of filepath.Walk for better performance

**Success criteria**: No panics, consistent error handling, better performance ✅

---

## Phase 5: UX improvements ✅

**Goal**: Better developer experience

- [x] Add urfave/cli library for proper flag parsing
  - Zero external dependencies (beyond stdlib)
  - Lightweight and fast
  - Declarative flag definitions
- [x] Refactor main.go to use urfave/cli structure
  - Move business logic to `run()` function
  - Define flags in Command struct
  - Remove manual arg parsing (PickFilter no longer needed)
- [x] Add `--help` flag with usage documentation (auto-generated)
- [x] Add `--version` flag (auto-generated)
- [x] Add verbose mode (`--verbose`) to show which package manager was selected and why
- [ ] Better directory structure (cmd/ and internal/)

**Success criteria**: Professional CLI with help, version, and verbose output ✅
**Note**: All core features implemented. Directory restructuring deferred as optional improvement.

---

## Phase 6: Distribution

**Goal**: Easy installation and automated releases

- [ ] Add goreleaser config for multi-platform builds
- [ ] Setup GitHub Actions workflow
  - Run tests on PR
  - Run linter on PR
  - Auto-release on tag push
- [ ] Create homebrew tap
- [ ] Update README with installation instructions (homebrew, manual)
- [ ] Add badge for build status and test coverage

**Success criteria**: `brew install <user>/vamos/vamos` works, automated releases on GitHub
