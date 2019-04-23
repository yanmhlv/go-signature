### Sign and check signature


```bash
openssl genrsa -out private.pem 2048 # generate priv key
openssl rsa -in private.pem -pubout > public.pub # extract pub key
```
