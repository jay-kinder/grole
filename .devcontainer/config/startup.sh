#!/bin/bash

# pip installs
pip install --break-system-packages pre-commit==3.8.0 detect-secrets==1.4.0

# install pre-commit in umbrella
pre-commit install

# clone zsh highlighting, zsh autosuggestions and p10k repos
git clone "${ZSH_SUGGESTIONS_REPO}" "${HOME}"/.oh-my-zsh/custom/plugins/zsh-autosuggestions
git clone "${ZSH_HIGHLIGHTING_REPO}" "${HOME}"/.oh-my-zsh/custom/plugins/zsh-syntax-highlighting

# set alias commands
{
    echo "alias cdc=\"clear && cd \${WORKSPACE}\""
    echo "alias gproj=\"gcloud config set project\""
    echo "alias gadl=\"gcloud auth application-default login\""
} >>"${HOME}/.zshrc"

# add zsh plugins
sed -i 's/plugins=(git)/plugins=(git gcloud golang zsh-autosuggestions zsh-syntax-highlighting)/g' "${HOME}"/.zshrc

# go installs
go install golang.org/x/tools/cmd/goimports@latest
