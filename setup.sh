#!/usr/bin/bash

if ! command -v go &> /dev/null; then
  echo "Go is not installed on your system."
  echo "Please install Go before running this setup."
  exit 1
fi

uname -o
if [[ $(uname -o) == "Android" ]] then
cp -rf $HOME/passcrax $HOME/.passcrax
cd $HOME/.passcrax
go build
ln -sf passcrax $HOME/.local/bin
echo "export PATH=$PATH:$HOME/.passcrax" >> ~/.bashrc
sleep 2
source ~/.bashrc
echo "  Process Completed"
echo "  Enter 'passcrax'"
echo "  If nothing worked, enter 'source ~/.bashrc' "


elif [[ $(uname -o) == "GNU/Linux" ]] then
cp -rf $HOME/passcrax $HOME/.passcrax
cd $HOME/.passcrax
go build
ln -sf passcrax $HOME/.local/bin
echo "export PATH=$PATH:$HOME/.passcrax" >> ~/.bashrc
sleep 2
source ~/.bashrc
echo "  Process Completed"
echo "  Enter 'passcrax'"
echo "  If nothing worked, enter 'source ~/.bashrc' "
fi
