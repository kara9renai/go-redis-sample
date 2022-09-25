# GolangでRedisのPub/Subを体験

## 手順
1. docker composeを立ち上げます
```
docker compose up -d
```
2. subの方をローカルで立ち上げます
```
go run sub/main.go
```
3. Userの構造体に合うようにエンドポイントにPOSTします
