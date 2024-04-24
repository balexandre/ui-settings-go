# Just a way to learn Go

Don't mind me, this is just an exercise to learn how to create a REST API using Golang

## TODO

- [x] create model
- [x] create a repository for handling DB
- [x] create repository tests
- [ ] create API (mux vs gin?)
- [ ] handle GET, POST, PATCH and DELETE
- [ ] create open api specs

### Tests output

```bash
❯ go test . -v
=== RUN   TestMongoOperations
2024/04/24 08:45:27 mongoDB connected!
2024/04/24 08:45:27 mongoDB pinged successfully!
=== RUN   TestMongoOperations/Insert_entry_#1
    repository_test.go:74: Insert entry #1 💪 ObjectID("6628aa878c3adbc650fe5ffc")
=== RUN   TestMongoOperations/Insert_entry_#2
    repository_test.go:91: Insert entry #2 💪 ObjectID("6628aa878c3adbc650fe5ffd")
=== RUN   TestMongoOperations/Get_by_merchant_code
    repository_test.go:101: Get by merchant code 💪 [{Alex 123  Alex123 0} {Bob 123  Alex123 1234567890}]
=== RUN   TestMongoOperations/Get_by_user_id
    repository_test.go:113: Get by user id 💪 [{Bob 123  Alex123 1234567890}]
=== NAME  TestMongoOperations
    repository_test.go:45: Database cleared 🤌
--- PASS: TestMongoOperations (0.01s)
    --- PASS: TestMongoOperations/Insert_entry_#1 (0.00s)
    --- PASS: TestMongoOperations/Insert_entry_#2 (0.00s)
    --- PASS: TestMongoOperations/Get_by_merchant_code (0.00s)
    --- PASS: TestMongoOperations/Get_by_user_id (0.00s)
PASS
ok      ui-settings     0.010s
```