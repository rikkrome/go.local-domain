## prereqs

### Install [mkcert](https://github.com/FiloSottile/mkcert)
```
brew install mkcert
mkcert -install

```

## Steps

Create certificate for local testing
```
mkcert go.rikkrome
```

Add name resolution just for local testing
```
sudo vim /etc/hosts # With following entry: 127.0.0.1 go.rikkrome
```

```
go mod init go.rikkrome/local-domain-server
```

Start server in the background
```
PORT=443 go run main.go &
```

test
```
curl https://go.rikkrome/
curl https://go.rikkrome/tokyo/
```