# golang-training
golang exercises of the course " Learn How To Code: Google's Go (golang) Programming Language"

## Ninja level 11

### EX1
Start with this code (). Instead of using the blank identifier, make sure the code is checking and handling the error.

### EX2
Start with this code (https://go.dev/play/p/9a1IAWy5E6). Create a custom error message using “fmt.Errorf”.

### EX3
Create a struct “customErr” which implements the builtin.error interface. Create a func “foo” that has a value of type error as a parameter. Create a value of type “customErr” and pass it into “foo”. If you need a hint, here is one (https://go.dev/play/p/L5QWV8-p11).

### EX4
Starting with this code (https://go.dev/play/p/wlEM1tgfQD), use the sqrt.Error struct as a value of type error. If you would like, use these numbers for your
- lat "50.2289 N"
- long "99.4656 W"

### EX5
We are going to learn about testing next. For this exercise, take a moment and see how much you can figure out about testing by reading the testing documentation (https://pkg.go.dev/testing) & also Caleb Doxsey’s article on testing (https://www.golang-book.com/books/intro/12). See if you can get a basic example of testing working.