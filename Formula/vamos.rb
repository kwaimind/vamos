class Vamos < Formula
  desc "Smart package manager runner for monorepos"
  homepage "https://github.com/kwaimind/vamos"
  version "0.3.0"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/kwaimind/vamos/releases/download/v0.3.0/vamos-0.3.0-darwin-arm64.tar.gz"
      sha256 "9a9e6e7b665f06c434e9914e3f4b348a1fc07919e5a53505a183ea39d0e2b73c"
    else
      url "https://github.com/kwaimind/vamos/releases/download/v0.3.0/vamos-0.3.0-darwin-amd64.tar.gz"
      sha256 "eeb04bb9d6573c141541e4cfd27e94f8537353eb5991ee4a2554b0b57b9ba5f5"
    end
  end

  def install
    bin.install "vamos"
  end

  test do
    assert_match "vamos", shell_output("#{bin}/vamos --version")
  end
end
