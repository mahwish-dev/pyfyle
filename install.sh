#!/usr/bin/env bash
REQUIRED_PKGS=("uv" "git" "hugo" "gum")

for pkg in "${REQUIRED_PKGS[@]}"; do
  if ! command -v "$pkg" >/dev/null 2>&1; then
    echo "Error: $pkg is required but not installed."
    exit 1
  fi
done
echo "All dependencies met!"
if ! git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
  echo "Not currently in a git repo"
  if gum confirm "Do you want to initialise a repository" </dev/tty; then
    git init
  fi
fi
if git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
  echo "Pyfyle can be installed as a git submodule or simply cloned"
  if gum confirm "Do you want to install Pyfyle as a git submodule?" </dev/tty; then
    git submodule add -f https://github.com/mahwish-dev/pyfyle.git pyfyle
  else
    git clone https://github.com/mahwish-dev/pyfyle.git
    if gum confirm "Do you want to add Pyfyle to your .gitignore" </dev/tty; then
      if [ -f ".gitignore" ]; then
        echo "pyfyle/" >>.gitignore
      else
        echo ".gitignore is missing."
        if gum confirm "Do you want to create one?" </dev/tty; then
          touch .gitignore
          echo "pyfyle/" >>.gitignore
        fi
      fi
    fi
  fi
else
  git clone https://github.com/mahwish-dev/pyfyle.git
fi

if gum confirm "Do you want to setup dashboard" </dev/tty; then
  cd pyfyle || exit
  echo 'DashboardEnabled = true' >>pyfyle.toml
  git submodule add https://github.com/sid314/pyfyle-hugo-site.git site

fi
cd pyfyle || exit
go build .
cd ..

echo "Pyfyle is installed, have a nice day"
