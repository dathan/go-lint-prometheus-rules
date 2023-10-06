class GoProjectTemplate < Formula
  desc "A generic brew formula to install the project build"
  homepage "https://go-lint-prometheus-rules.gihub.io"
  url "https://github.com/dathan/go-lint-prometheus-rules.git", using: :git #download strategy
  revision 1
  version 'master'
  head "https://github.com/dathan/go-lint-prometheus-rules.git"
  
  depends_on "make" => :build
  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath
    path = buildpath/"src/github.com/dathan/go-lint-prometheus-rules"
    path.install Dir["*"]
    cd path do
      system "make", "build"
      system "ls -ltarh"
    end

    bin.install path/"bin/example1" => "example1"
  end

  test do
    system "true"
  end
end 
