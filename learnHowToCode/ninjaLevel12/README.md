# golang-training
golang exercises of the course " Learn How To Code: Google's Go (golang) Programming Language"

## Ninja level 12

### EX1
Start with this code (https://github.com/GoesToEleven/go-programming/tree/master/code_samples/010-ninja-level-thirteen/01/starting-code). Get the code ready to BET on the code (add benchmarks, examples, tests).
Run the following in this order:
- tests
- benchmarks
- coverage
- coverage shown in web browser
- examples shown in documentation in a web browser

```bash
godoc -http=:8080

go test
go test -bench .
go test -cover
go test -coverprofile c.out
go tool cover -html=c.out
```

### EX2
Start with this code (https://github.com/GoesToEleven/go-programming/tree/master/code_samples/010-ninja-level-thirteen/02/01-code-starting). Get the code ready to BET on (add benchmarks, examples, tests) however do not write an example for the func that returns a map; and only write a test for it as an extra challenge. Add documentation to the code. Run the following in this order:
- tests
- benchmarks
- coverage
- coverage shown in web browser
- examples shown in documentation in a web browser

### EX3
Start with this code (). Get the code ready to BET on (add benchmarks, examples, tests). Write a table test. Add documentation to the code. Run the following in this order:
- tests
- benchmarks
- coverage
- coverage shown in web browser
- examples shown in documentation in a web browser
helpful to know:
1. https://play.golang.org/p/4GUqs1HMpp
2. https://play.golang.org/p/P9unTIFeOq
