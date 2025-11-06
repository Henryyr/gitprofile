class Gitprofile < Formula
  desc "A CLI to manage multiple git profiles"
  homepage "https://github.com/henryyr/gitprofile"
  url "https://github.com/henryyr/gitprofile/archive/refs/tags/v0.1.0.tar.gz" # This will be updated by goreleaser
  sha256 "..." # This will be updated by goreleaser
  license "Apache-2.0"

  depends_on "go" => :build

  def install
    system "go", "build", "-o", "#{bin}/gitprofile", "."
  end

  test do
    system "#{bin}/gitprofile", "--version"
  end
end
