# BankRate
Plurk bot 取得 銀行匯率
# Window 32 environment
CGO_ENABLED=0 GOOS=windows GOARCH=386 
# Window 64 environment
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 
# Linux 32 enviromen
CGO_ENABLED=0 GOOS=linux GOARCH=386
# Linux 64 enviroment
CGO_ENABLED=0 GOOS=linux GOARCH=amd64
# Docker-Compose Run Command
docker-compose up --build -d
