{
  description = "A supercharged cd wrapper with aliases and TUI.";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
    ...
  }:
    flake-utils.lib.eachDefaultSystem (
      system: let
        pkgs = nixpkgs.legacyPackages.${system};

        buildScd = {
          src,
          version,
        }:
          pkgs.buildGoModule {
            pname = "scd";
            inherit version src;
            vendorHash = "sha256-NgIc1yRVP74hyE/Bfsr+Cl3MRgylgO+CzTdWRjjRGEg="; # Update if source changes
            ldflags = [
              "-s"
              "-w"
              "-X main.version=${version}"
            ];
            nativeBuildInputs = [pkgs.installShellFiles];
            postInstall = ''
              mv $out/bin/srn-cd $out/bin/scd
            '';
            postFixup = ''
              installShellCompletion --fish ${src}/completions/scd.fish
              installShellCompletion --zsh ${src}/completions/scd.zsh
              installShellCompletion --bash ${src}/completions/scd.bash
            '';
          };

        cleanedSource = pkgs.lib.cleanSourceWith {
          src = ./.;
          filter = path: type: let
            baseName = baseNameOf path;
          in
            baseName == ".version" || pkgs.lib.cleanSourceFilter path type;
        };
      in {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            golangci-lint
            cmake
            goreleaser
          ];
        };

        packages.default = buildScd {
          src = cleanedSource;
          version = let
            versionFile = "${cleanedSource}/.version";
          in
            pkgs.lib.escapeShellArg (
              if builtins.pathExists versionFile
              then builtins.readFile versionFile
              else self.shortRev or "dev"
            );
        };

        apps.default = flake-utils.lib.mkApp {drv = self.packages.${system}.default;};
      }
    );
}
