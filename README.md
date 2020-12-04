# Harpokratos

Harpocrates (Ancient Greek: Ἁρποκράτης) was the god of silence, 
secrets and confidentiality in the Hellenistic religion developed 
in Ptolemaic Alexandria.

![harpo](https://upload.wikimedia.org/wikipedia/commons/9/93/Harpocrates_gulb_082006.JPG)


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
