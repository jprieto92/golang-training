# golang-training
golang exercises of the course " Learn How To Code: Google's Go (golang) Programming Language"

## Ninja level 7

### EX1
- Create a value and assign it to a variable.
- Print the address of that value.

### EX2
1. create a person struct
2. create a func called “changeMe” with a *person as a parameter
- change the value stored at the *person address
  - important:
    - to dereference a struct, use (*value).field
      - p1.first
      - (*p1).first
    - “As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method), x.f is shorthand for (*x).f.”
      - https://golang.org/ref/spec#Selectors

3. create a value of type person
- print out the value
4. in func main
- call “changeMe”
5. in func main
- print out the value