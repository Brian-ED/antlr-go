{
  buildPythonPackage,
  fetchFromGitHub,
  lib,
  python,
  setuptools,
}:

let
  version = "v0.2.0";
in
buildPythonPackage {
  pname = "tangled-up-in-unicode";
  inherit version;
  pyproject = true;

  src = fetchFromGitHub {
    owner = "dylan-profiler";
    repo = "tangled-up-in-unicode";
    tag = version;
    hash = "sha256-fp2DNWNUQzJ5EtgtSXbzXfUkgwowe3saPdZZb2UFuaE=";
  };

  build-system = [ setuptools ];

  pythonImportsCheck = [ "tangled_up_in_unicode" ];

  meta = {
    description = "Provides access to character properties for all Unicode characters, from the Unicode Character Database";
    homepage = "https://github.com/dylan-profiler/tangled-up-in-unicode";
    license = lib.licenses.bsdOriginal;
  };
}
