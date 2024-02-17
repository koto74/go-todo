# go-todo

## 開発環境の起動
```
docker compose up
```

## 作成したMySQLコンテナの確認
```
コンテナにログイン
docker exec it [CONTAINER ID] bash
```
```
# ログインユーザの確認
id

# Linuxカーネルの確認
uname -a

# Linuxディストリビューションの確認
cat /etc/os-release

# MySQLのバージョンの確認
mysql --version

# MySQLにログイン
mysql -u root -p
# パスワードを聞かれるので、rootと入力

-- データベース一覧
show databases;

-- データベース切り替え
use demo

-- 接続中のデータベースを確認
select database();

-- 文字コードの確認
show variables like '%char%';

-- 照合順序の確認
show variables like '%collation%';
```
