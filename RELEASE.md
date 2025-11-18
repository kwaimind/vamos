# release process

## automated release (recommended)

Run the release script with the version number:

```bash
./release.sh 0.1.0
```

This will:
1. Update version in `main.go`
2. Build release tarballs for darwin-amd64 and darwin-arm64
3. Calculate SHA256 hashes
4. Update `Formula/vamos.rb` with new version and hashes
5. Create git tag and commit
6. Push changes to GitHub
7. Create GitHub release with tarballs
8. Clean up local tarballs

**Prerequisites:**
- GitHub CLI installed: `brew install gh`
- Authenticated with GitHub: `gh auth login`

## manual release

1. Update version in `main.go`:
   ```go
   var version = "0.1.0"
   ```

2. Build release tarballs:
   ```bash
   make release VERSION=0.1.0
   ```

3. Create GitHub release:
   - Go to https://github.com/kwaimind/vamos/releases/new
   - Tag: `v0.1.0`
   - Title: `v0.1.0`
   - Upload both `.tar.gz` files generated

4. Update Homebrew formula with SHA256 hashes:
   ```bash
   shasum -a 256 vamos-0.1.0-darwin-amd64.tar.gz
   shasum -a 256 vamos-0.1.0-darwin-arm64.tar.gz
   ```

   Update `Formula/vamos.rb`:
   - Replace `PLACEHOLDER_AMD64_SHA256` with amd64 hash
   - Replace `PLACEHOLDER_ARM64_SHA256` with arm64 hash
   - Update version number if needed

5. Commit and push formula changes:
   ```bash
   git add Formula/vamos.rb
   git commit -m "chore: update formula for v0.1.0"
   git push
   ```

6. Test installation:
   ```bash
   brew uninstall vamos  # if previously installed
   brew install kwaimind/vamos/vamos https://github.com/kwaimind/vamos
   vamos --version
   ```
