# Create directory for certificates
mkdir certs
cd certs

# Generate CA private key and certificate
openssl genrsa -out ca.key 4096
openssl req -new -x509 -days 365 -key ca.key -out ca.crt -subj "/C=US/ST=State/L=City/O=Organization/OU=CA/CN=Root CA"

# Generate server private key
openssl genrsa -out server.key 4096

# Generate server CSR using the configuration file
openssl req -new -key server.key -out server.csr -config ../server.conf

# Generate server certificate
openssl x509 -req -days 365 -in server.csr \
    -CA ca.crt -CAkey ca.key -CAcreateserial \
    -out server.crt \
    -extensions v3_req \
    -extfile ../server.conf

# Generate client private key and CSR
openssl genrsa -out client.key 4096
openssl req -new -key client.key -out client.csr -subj "/C=US/ST=State/L=City/O=Organization/OU=Client/CN=client"

# Generate client certificate
openssl x509 -req -days 365 -in client.csr \
    -CA ca.crt -CAkey ca.key -CAcreateserial \
    -out client.crt