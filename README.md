# asimov-backend
Backend of the Asimov 3D printer

# Install
1. Clone repo
    ```bash
    git clone git@github.com:3DDiamantes/asimov-backend.git
    ```
2. Go to the root folder and download the dependencies
    ```bash
    cd asimov-backend
    go mod download
    ```

# Test
`go test -v ./...`

# Coverage report
```
go test -coverpkg=./... -coverprofile=cov.out -covermode=atomic ./... && go tool cover -html=cov.out && rm cov.out
```
or run the bash file
```
./RunCoverage
```

# Deploy
## On local machine
1. Compile
    ```
    GOOS=linux GOARCH=arm GOARM=7 go build -o builds/x.x.x-asimov-backend cmd/main.go
    ```
2. Upload
    ```
    scp builds/x.x.x-asimov-backend [user]@[server_ip]:/builds
    ```
# On target machine
3. Make executable and run
    ```
    sudo chmod 545 /builds/x.x.x-asimov-backend
    /builds/x.x.x-asimov-backend
    ```
