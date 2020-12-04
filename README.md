# harpokratos

### creating a user in vault

http://localhost:8200/ui/vault/policies/acl

\+ create policy

harpokratos -> 
```hcl
path "config/*" {
                   capabilities = ["read", "list"]
               }
```

create user

http://localhost:8200/ui/vault/access/userpass/item/user

\+ create user

username: harpokratos
password: harpokratos
generated token's type: harpokratos
