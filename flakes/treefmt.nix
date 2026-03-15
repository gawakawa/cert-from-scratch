_: {
  perSystem = _: {
    treefmt = {
      programs = {
        gofmt = {
          enable = true;
          includes = [ "*.go" ];
        };
        goimports = {
          enable = true;
          includes = [ "*.go" ];
        };
        golines = {
          enable = true;
          maxLength = 70;
          includes = [ "*.go" ];
        };
        nixfmt = {
          enable = true;
          includes = [ "*.nix" ];
        };
      };
    };
  };
}
