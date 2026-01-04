# Serein Change Directory

Serein is a supercharged cd wrapper with aliases and TUI.

## Installation

### For NixOS/Home Manager Configurations

If you manage your system or user environment with NixOS or Home Manager flakes, you can add `srn-cd` as an input to your configuration.

1.  **Add `srn-cd` as an input in your `flake.nix`:**

    ```nix
    {
      description = "Your personal NixOS/Home Manager configuration";

      inputs = {
        nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
        home-manager.url = "github:nix-community/home-manager";
        home-manager.inputs.nixpkgs.follows = "nixpkgs";

        srn-cd = {
          url = "github:infraflakes/srn-cd";
          inputs.nixpkgs.follows = "nixpkgs";
        };
      };

      outputs = { self, nixpkgs, home-manager, srn-cd, ... } @ inputs: {
        # ... your existing outputs
      };
    }
    ```

2.  **Install the CLI in your NixOS or Home Manager configuration:**

    The flake provides a `default` package.

    **Option A: Install System-Wide (NixOS Configuration)**

    ```nix
    # In your configuration.nix (or a NixOS module)
    { config, pkgs, lib, ... }:

    {
      environment.systemPackages = with pkgs; [
        # Reference it from the srn-cd flake input
        inputs.srn-cd.packages.${pkgs.stdenv.hostPlatform.system}.default

      ];

      # ... other system configurations
    }
    ```

    **Option B: Install via Home Manager (User-Specific)**

    ```nix
    # In your Home Manager configuration (e.g., ~/.config/home-manager/home.nix)
    { config, pkgs, ... }:

    {
      home.packages = [
        # Reference it from the srn-cd flake input
        inputs.srn-cd.packages.${pkgs.stdenv.hostPlatform.system}.default

      ];

      # ... other Home Manager options
    }
    ```

### Binary Distribution (For Non-Nix Users)

For users not using Nix, the CLI can be downloaded as a single executable binary.

1.  **Download the latest release:**
    Visit the [GitHub Releases page](https://github.com/infraflakes/srn-cd/releases) and download the wanted binary.

2.  **Make the binary executable:**
    ```bash
    chmod +x scd
    ```

3.  **Move the binary to your PATH (optional but recommended):**
    ```bash
    sudo mv scd /usr/local/bin/
    ```

### Manual Installation (from source)

If you have a Go environment set up, you can build from source.

1.  **Clone the repo:**
    ```bash
    git clone https://github.com/infraflakes/srn-cd
    cd srn-cd
    ```

2.  **Build the binary:**
    The included `Makefile` provides an easy way to build the application:
    ```bash
    make build
    ```
    Alternatively, you can use the standard Go command:
    ```bash
    CGO_ENABLED=0 go build -o scd .
    ```

#### Caution!

**In order for `scd` to work we need to generate shell functions for it to able to change directory:**
Add these to your shell config:

Bash:

    eval "$(scd init bash)"

Zsh:

    eval "$(scd init zsh)"

Fish:

    scd init fish | source

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.
