let
  _pkgs = import <nixpkgs> {};
in
{ pkgs ? import (_pkgs.fetchFromGitHub { owner = "NixOS";
                                         repo = "nixpkgs-channels";
                                         rev = "fd3d2b1a8dca6a4a90093a54cecbdb3e66f163fb";
                                         sha256 = "06cj963b59l29vzsk4d0g89v4kjg4ycgvbgz135yyqfrnn5hlhzn";
                                       }) {}
}:

pkgs.stdenv.mkDerivation rec {
    name = "docker-volume-ipfs-dev";
    env = pkgs.buildEnv { name = name; paths = buildInputs; };
    buildInputs = [
	pkgs.python35
    	pkgs.python35Packages.virtualenv
    	pkgs.python35Packages.pip
        pkgs.vndr
        pkgs.go_1_7
        pkgs.gnumake
        pkgs.ipfs
        pkgs.docker
    ];
}
