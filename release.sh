#!/bin/bash
set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo_error() { echo -e "${RED}❌ $1${NC}"; }
echo_success() { echo -e "${GREEN}✅ $1${NC}"; }
echo_info() { echo -e "${YELLOW}ℹ️  $1${NC}"; }

# Check if version provided
if [ -z "$1" ]; then
  echo_error "Version required. Usage: ./release.sh 0.1.0"
  exit 1
fi

VERSION=$1
TAG="v${VERSION}"

# Check if gh CLI is installed
if ! command -v gh &> /dev/null; then
  echo_error "GitHub CLI (gh) not installed. Install with: brew install gh"
  exit 1
fi

# Check if authenticated
if ! gh auth status &> /dev/null; then
  echo_error "Not authenticated with GitHub. Run: gh auth login"
  exit 1
fi

echo_info "Starting release process for version ${VERSION}"

# Update version in main.go
echo_info "Updating version in main.go..."
sed -i '' "s/var version = \".*\"/var version = \"${VERSION}\"/" main.go
echo_success "Version updated in main.go"

# Build release tarballs
echo_info "Building release tarballs..."
make release VERSION=${VERSION}
echo_success "Release tarballs built"

# Calculate SHA256 hashes
echo_info "Calculating SHA256 hashes..."
AMD64_SHA=$(shasum -a 256 vamos-${VERSION}-darwin-amd64.tar.gz | awk '{print $1}')
ARM64_SHA=$(shasum -a 256 vamos-${VERSION}-darwin-arm64.tar.gz | awk '{print $1}')
echo_success "SHA256 hashes calculated"

# Update Formula
echo_info "Updating Formula/vamos.rb..."
sed -i '' "s/version \".*\"/version \"${VERSION}\"/" Formula/vamos.rb
sed -i '' "s|download/v.*/vamos-.*-darwin-arm64.tar.gz|download/${TAG}/vamos-${VERSION}-darwin-arm64.tar.gz|" Formula/vamos.rb
sed -i '' "s|download/v.*/vamos-.*-darwin-amd64.tar.gz|download/${TAG}/vamos-${VERSION}-darwin-amd64.tar.gz|" Formula/vamos.rb
sed -i '' "s/PLACEHOLDER_ARM64_SHA256/${ARM64_SHA}/" Formula/vamos.rb
sed -i '' "s/PLACEHOLDER_AMD64_SHA256/${AMD64_SHA}/" Formula/vamos.rb
sed -i '' "s/sha256 \".*\" # arm64/sha256 \"${ARM64_SHA}\"/" Formula/vamos.rb
sed -i '' "s/sha256 \".*\" # amd64/sha256 \"${AMD64_SHA}\"/" Formula/vamos.rb
echo_success "Formula updated"

# Create git tag
echo_info "Creating git tag ${TAG}..."
git add main.go Formula/vamos.rb
git commit -m "chore: release ${TAG}"
git tag ${TAG}
echo_success "Git tag created"

# Push changes
echo_info "Pushing changes..."
git push origin main
git push origin ${TAG}
echo_success "Changes pushed"

# Create GitHub release
echo_info "Creating GitHub release..."
gh release create ${TAG} \
  vamos-${VERSION}-darwin-amd64.tar.gz \
  vamos-${VERSION}-darwin-arm64.tar.gz \
  --title "${TAG}" \
  --notes "Release ${TAG}"
echo_success "GitHub release created"

# Cleanup
echo_info "Cleaning up tarballs..."
rm vamos-${VERSION}-darwin-amd64.tar.gz vamos-${VERSION}-darwin-arm64.tar.gz
echo_success "Cleanup complete"

echo ""
echo_success "Release ${TAG} complete!"
echo_info "Users can now install with: brew install kwaimind/vamos/vamos https://github.com/kwaimind/vamos"
echo_info "Test with: brew uninstall vamos && brew install kwaimind/vamos/vamos https://github.com/kwaimind/vamos"
