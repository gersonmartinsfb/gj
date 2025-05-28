# gj

## Pre Requisitos

### Token Jira
Criar token do Jira


### Variáveis de Ambiente

```bash
GJ_TOKEN="1234asdf12345"
GJ_USER=email@foxbit.com.br
GJ_DOMAIN=foxbit.atlassian.net
GJ_ISSUE_PREFIX=QUANT
```

## Compilação e instalação

```bash
make build
mkdir -p ~/bin
copy bin/gj ~/bin 
export PATH=$PATH:$HOME/bin # <- colocar esse no .bashrc ou .zhrc ou equivalentes
```

## Uso

```bash
Usage: gj [--disable-creation] <type> <issue-number>
Options:
  type: The type of issue (bug, feat, hotfix, enhance)
  issue-number: The issue number to create a branch for. Just the number. The suffix QUANT- will be added automatically.
  -disable-creation
        Disable branch creation and just print the command that would be run
```        