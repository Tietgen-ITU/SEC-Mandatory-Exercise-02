rm *.crt *.key *.csr

# ca-key.key is the private key and the ca-cert.crt is the signed CA certificate
openssl req -x509 -sha256 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.key -out ca-cert.crt -subj "C=DK/ST=Copenhagen/L=Copenhagen/O=IT University/OU=Students/CN=bob/emailAddress=anti@itu.dk"

# Now we generate the server certificates
# We start of by creating the private key for the server
openssl genrsa -out server-key.key 2048

# In order to create the server signed public key certificate then we need to create a
# CSR (Certificate Signing Request)
openssl req -new -key server-key.key -out server.csr -config csr.conf

# Now we want to create the signed public key certificate for the server
openssl x509 -req -in server.csr -CA ca-cert.crt -CAkey ca-key.key -CAcreateserial -out server.crt -days 365 -sha256 -extfile server.conf