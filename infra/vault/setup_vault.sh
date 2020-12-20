#!/bin/bash

export VAULT_TOKEN=$(cat cluster-keys.json | jq -r ".root_token")
export VAULT_ADDR=http://localhost:8200
kubens harpokratos
#kubectl exec harpokratos-vault-0 -- vault policy write harpokratos $PWD/harpokratos-acl.hcl

curl --header "X-Vault-Token: $VAULT_TOKEN" \
       --request PUT \
       --data @$PWD/acl.json \
       $VAULT_ADDR/v1/sys/policies/acl/harpokratos