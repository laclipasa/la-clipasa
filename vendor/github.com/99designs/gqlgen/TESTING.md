How to write tests for gqlgen
===

Testing generated code is a little tricky, here's how its currently set up.

### Testing responses from a server

There is a server in `codegen/testserver` that is generated as part
of `go generate ./...`, and tests written against it.

There are also a bunch of tests in against the examples, feel free to take examples from there.


### Testing the errors generated by the binary

These tests are **really** slow, because they need to run the whole codegen step. Use them very sparingly. If you can, find a way to unit test it instead.

Take a look at `codegen/testserver/input_test.go` for an example.

### Testing introspection

Introspection is tested by diffing the output of `graphql get-schema` against an expected output.

Setting up the integration environment is a little tricky:
```bash
cd integration
go generate ./...
go run ./server/cmd/integration/server.go
```
in another terminal
```bash
cd integration
npm install
./node_modules/.bin/graphql-codegen
```

will write the schema to `integration/schema-fetched.graphql`, compare that with `schema-expected.graphql`

CI will run this and fail the build if the two files don't match.