{
  description = "A very basic flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    treefmt-nix = {
      url = "github:numtide/treefmt-nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs = { self, nixpkgs, flake-utils, treefmt-nix }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };

        treefmtEval = treefmt-nix.lib.evalModule pkgs {
          projectRootFile = "flake.nix";
          programs.nixpkgs-fmt.enable = true;
          programs.prettier = {
            enable = true;
            includes = [ "*.md" "*.yaml" "*.yml" "*.html" "*.css" "*.js" ];
          };
          programs.gofmt.enable = true;
        };

        #goModVendorHash = pkgs.lib.fakeHash;
        goModVendorHash = "sha256-OOT7sgYRgkpXKYXpFe4anMj3G+PA4tiS0FaIoHDWPpQ=";

        todo-cli = pkgs.buildGoModule {
          name = "todo-cli";
          version = "0.0.1";
          vendorHash = goModVendorHash;
          src = ./.;
        };
      in
      {
        formatter = treefmtEval.config.build.wrapper;
        checks.formatter = treefmtEval.config.build.check self;

        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            go
          ];
        };

        packages = flake-utils.lib.flattenTree {
          default = todo-cli;
        };
      }
    );
}
