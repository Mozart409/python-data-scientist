# Use an official Python runtime as a parent image
FROM python:3.8

# Install Go
RUN apt-get update && \
    apt-get install -y wget zsh curl git zip && \
    wget https://dl.google.com/go/go1.21.5.linux-amd64.tar.gz && \
    tar -xvf go1.21.5.linux-amd64.tar.gz && \
    mv go /usr/local && \
    rm go1.21.5.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN go install -v golang.org/x/tools/gopls@latest

RUN curl -o ~/.zshrc -L https://raw.githubusercontent.com/Mozart409/.dotfiles/main/zsh/.zshrc.sample
RUN curl -L git.io/antigen > ~/antigen.zsh

RUN curl -fsSL https://dprint.dev/install.sh | sh

RUN chsh -s $(which zsh)

RUN zsh -c "source ~/.zshrc"

RUN pip install notebook

# Set up the working directory
WORKDIR /workspace
