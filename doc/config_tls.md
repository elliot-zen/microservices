# TLS 配置

## 生成证书

### Privatekey and self-signed certificate

```sh
openssl req -x509 \
-sha256 \
-nodes \
-newkey rsa:4096 \
-days 365 \
-keyout ca.key \
-out ca.crt \
-subj "/C=CN/ST=beijing/L=beijing/O=MyOrg/OU=Personal/CN=*.foo.top" \
-addext "subjectAltName=DNS:*.foo.top,IP:127.0.0.1" 
```



### Server Request Sign

#### Sign Request

```sh
openssl req \
-sha256 -nodes \
-newkey rsa:4096 \
-keyout server.key \
-out server.csr \
-subj "/C=CN/ST=Beijing/L=Beijing/O=MyOrg/OU=Personal/CN=*.foo.top" \
-addext "subjectAltName=DNS:*.foo.top,IP:127.0.0.1"
```



#### Use CA's private key to sign the request 

```sh
openssl x509 \
-req -in server.csr \
-days 365 \
-CA ca.crt \
-CAkey ca.key \
-CAcreateserial \
-out server.crt \
-sha256 \
-extfile <(printf "subjectAltName=DNS:*.foo.top,IP:127.0.0.1")
```
```



### Client Request Sign

#### Sign Request

```sh
openssl req \
-sha256 -nodes \
-newkey rsa:4096 \
-keyout client.key \
-out client.csr \
-subj "/C=CN/ST=Beijing/L=Beijing/O=MyOrg/OU=Personal/CN=*.foo.top"	 \
-addext "subjectAltName = DNS:*.foo.top,IP:127.0.0.1"
```



#### Use CA's private key to sign the request 

```sh
openssl x509 \
-req -in client.csr \
-days 365 \
-CA ca.crt \
-CAkey ca.key \
-CAcreateserial \
-out client.crt \
-sha256 \
-extfile <(printf "subjectAltName=DNS:*.foo.top,IP:127.0.0.1")
```


### Verify

```sh
### Verify
openssl verify -CAfile ca.crt server.crt
openssl verify -CAfile ca.crt client.crt
openssl x509 -in client.crt -noout -text | grep -A1 "Subject Alternative Name"
```

