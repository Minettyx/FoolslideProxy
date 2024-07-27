{
  description = "FoolslideProxy";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    nixpkgs,
    flake-utils,
    ...
  }:
    flake-utils.lib.eachDefaultSystem (
      system: let
        pkgs = import nixpkgs {inherit system;};

        nativeBuildInputs = with pkgs; [
          go
          gopls
        ];
        buildInputs = with pkgs; [];

        goModule = pkgs.buildGoModule rec {
          name = "foolslideproxy";
          src = ./.;
          inherit buildInputs;
          vendorHash = "sha256-hEyV9aSe5jypKvZMt1doejQd8ARAL1hZEl0LTmQUPlo=";
        };

        dockerImage = pkgs.dockerTools.buildLayeredImage {
          name = "foolslideproxy";
          contents = with pkgs; [
            busybox
            dockerTools.binSh
          ];
          config = {
            Cmd = ["${goModule}/bin/foolslideproxy"];
          };
        };
      in {
        devShells.default = pkgs.mkShell {inherit nativeBuildInputs buildInputs;};

        packages = {
          default = goModule;
          docker = dockerImage;
        };
      }
    );
}
