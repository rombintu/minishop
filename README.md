# Minishop

## Prechecks
```
git clone https://github.com/rombintu/minishop.git
cd minishop
cp config/config.toml.bak config/config.toml
vim config/config.toml
```
## Migrate database (sqlite3)
```
sqlite3 dev.db < config/migrate_sqlite.sql
```
## Run or Build
```
make run
make build
```

### For debug
```
cp dev.db server/cmd/dev.db
```
## Docker build
```
docker build -t minishop docker/
docker run -d -p <PORT>:<PORT> minishop
```

## Push to git
`make push`