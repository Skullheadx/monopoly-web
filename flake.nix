{
  description = "Monopoly web game";
  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
  };

  outputs =
    { self, nixpkgs }:
    let
      version = "0.1";
      system = "x86_64-linux";
      pkgs = import nixpkgs {inherit system;};
    in
    {
      packages.${system}.default = pkgs.buildGoModule {
        pname = "monopoly-web";
        inherit version;
        src = ./.;

        # This hash locks the dependencies of this package. It is
        # necessary because of how Go requires network access to resolve
        # VCS.  See https://www.tweag.io/blog/2021-03-04-gomod2nix/ for
        # details. Normally one can build with a fake hash and rely on native Go
        # mechanisms to tell you what the hash should be or determine what
        # it should be "out-of-band" with other tooling (eg. gomod2nix).
        # To begin with it is recommended to set this, but one must
        # remember to bump this hash when your dependencies change.
        # vendorHash = pkgs.lib.fakeHash;

        vendorHash = null; 
      };

      # Add dependencies that are only needed for development
      devShells.${system} = {
        default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            gopls
            gotools
            go-tools
          ];
        };
      };
    };

}
