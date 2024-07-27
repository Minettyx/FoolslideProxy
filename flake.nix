{
  description = "Proxy that converts manga reading websites to foolslide (for Tachiyomi) ";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = {nixpkgs, ...}: let
    lib = nixpkgs.lib;
    systems = [
      "x86_64-linux"
      "x86_64-darwin"
      "aarch64-darwin"
      "aarch64-linux"
    ];
    forEachSystem = f: lib.genAttrs systems (system: f pkgsFor.${system});
    pkgsFor = lib.genAttrs systems (system:
      import nixpkgs {
        inherit system;
        config.allowUnfree = true;
      });
  in {
    devShells = forEachSystem (pkgs: {
      default = pkgs.mkShell {
        packages = with pkgs; [
          go
          gopls
          templ
        ];
      };
    });

    packages = forEachSystem (pkgs: {
      default = pkgs.buildGoModule {
        name = "foolslideproxy";
        src = ./.;
        vendorHash = "sha256-HxqmyBF03zZSIqo6g8lLDO9+Au0Zpt6jgAzZGg8p/tw=";
        nativeBuildInputs = [];
      };
    });
  };
}
