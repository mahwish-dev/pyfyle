#!/usr/bin/env bash
REQUIRED_PKGS=("uv" "git" "hugo" "gum")

for pkg in "${REQUIRED_PKGS[@]}"; do
  if ! command -v "$pkg" >/dev/null 2>&1; then
    echo "Error: $pkg is required but not installed."
    exit 1
  fi
done
gum log "All dependencies met!" --structured --level info
if ! git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
  gum log "Not currently in a git repo" --structured --level info

  if gum confirm "Do you want to initialise a repository" </dev/tty; then
    git init
  fi
fi
if git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
  gum log "Pyfyle can be installed as a git submodule or simply cloned" --structured --level info

  if gum confirm "Do you want to install Pyfyle as a git submodule?" </dev/tty; then
    git submodule add -f https://github.com/mahwish-dev/pyfyle.git pyfyle
  else
    git clone https://github.com/mahwish-dev/pyfyle.git
    if gum confirm "Do you want to add Pyfyle to your .gitignore" </dev/tty; then
      if [ -f ".gitignore" ]; then
        echo "pyfyle/" >>.gitignore
      else
        gum log ".gitignore is missing." --structured --level info

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
cd pyfyle || exit
git pull
cd ..
if gum confirm "Do you want to setup dashboard" </dev/tty; then
  cd pyfyle || exit
  echo 'DashboardEnabled = true' >>pyfyle.toml
  git submodule add https://github.com/sid314/pyfyle-hugo-site.git site
  cd site || exit
  rm -rf themes/re-terminal
  git submodule add -f https://github.com/mirus-ua/hugo-theme-re-terminal.git themes/re-terminal
  cd ..
  cd ..

else
  echo 'DashboardEnabled = false' >>pyfyle.toml
fi
cd pyfyle || exit
chmod +x bin/pyfyle
chmod +x bin/pyfyle-viewer
chmod +x bin/pyfyle-dashboard
go build .
mv pyfyle bin/pyfyle
cd ..

gum log "Pyfyle is installed, have a nice day" --structured --level info
