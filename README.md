## Arbitrary Code Execution Service In Go

You'll need:

- [Go](https://golang.org/) version 1.19 or better installed.
- Docker 

Start the server:

```bash
go run main.go run standalone
```

Execute a code snippet. Example

```bash
curl \
  -s \
  -X POST \
  -H "content-type:application/json" \
  -d '{"language":"python","code":"print(\"hello world\")"}' \
  http://localhost:8000/execute
```

Should output:

```bash
hello world
```

You can try changing the `language` to `go` or `bash`.

## Frontend 

```shell
cd frontend
npm install
npm run dev
```


