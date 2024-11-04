# go-repository-in-transaction

Repository の同一メソッドにて、Transaction の有無で処理を切り替えるサンプル。
（DDD 的な良し悪しは一旦置いておく）

## クエリログ

### 非トランザクション時

```
LOG:  execute <unnamed>: INSERT INTO users (id, name, address) VALUES ($1, $2, $3)
DETAIL:  parameters: $1 = '100', $2 = 'Suzuki', $3 = 'Saitama'
LOG:  execute <unnamed>: UPDATE users SET name = $1, address = $2 WHERE id = $3
DETAIL:  parameters: $1 = 'Suzuki', $2 = 'Saitama', $3 = '100'
LOG:  execute <unnamed>: DELETE FROM users WHERE id = $1
DETAIL:  parameters: $1 = '100'
```

### トランザクション時

```
LOG:  statement: BEGIN ISOLATION LEVEL REPEATABLE READ READ WRITE
LOG:  execute <unnamed>: INSERT INTO users (id, name, address) VALUES ($1, $2, $3)
DETAIL:  parameters: $1 = '200', $2 = 'Takahashi', $3 = 'Osaka'
LOG:  execute <unnamed>: UPDATE users SET name = $1, address = $2 WHERE id = $3
DETAIL:  parameters: $1 = 'Takahashi', $2 = 'Osaka', $3 = '200'
LOG:  execute <unnamed>: DELETE FROM users WHERE id = $1
DETAIL:  parameters: $1 = '200'
LOG:  statement: COMMIT
```
