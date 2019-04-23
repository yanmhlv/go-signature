### Sign and check signature

[![CircleCI](https://circleci.com/gh/yanmhlv/go-signature.svg?style=svg)](https://circleci.com/gh/yanmhlv/go-signature)


```bash
openssl genrsa -out private.pem 2048 # generate priv key
openssl rsa -in private.pem -pubout > public.pub # extract pub key
```
