{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = import nixpkgs {inherit system;};
      lastModifiedDate = self.lastModifiedDate or self.lastModified or "19700101";
      version = builtins.substring 0 8 lastModifiedDate;
    in {
      packages = {
        default = pkgs.buildGoModule {
          inherit version;
          pname = "pocketbase-template";
          src = ./.;
          vendorHash = "sha256-KIwUG9dTaaQD/Fb1fap+7jbRfYMsi0KS7VfHTdAqN2U=";
        };
      };

      devShells = {
        default = pkgs.mkShell {
          packages = with pkgs; [
            go
          ];
        };
      };
    });
}
