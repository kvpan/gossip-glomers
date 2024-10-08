{
  description = "Fly.io distributed systems challenge";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            gopls
            golangci-lint
            delve
            go-tools

            # Maelstrom
            maelstrom-clj
          ];

          shellHook = ''
            echo "Go development environment"
            echo "Go version: $(go version)"
            export GOPATH=$PWD/.env/go
            export PATH=$PATH:$GOPATH/bin
          '';
        };
      }
    );
}
