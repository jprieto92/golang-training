# golang-training
golang exercises of the course " Learn How To Code: Google's Go (golang) Programming Language"

## Ninja level 6

### EX1
1. Create a func with the identifier foo that returns an int
2. Create a func with the identifier bar that returns an int and a string
3. Call both funcs
4. Print out their results

### EX2
1. create a func with the identifier foo that
- takes in a variadic parameter of type int
- pass in a value of type []int into your func (unfurl the []int)
- returns the sum of all values of type int passed in
2. create a func with the identifier bar that
- takes in a parameter of type []int
- returns the sum of all values of type int passed in

### EX3
Use the “defer” keyword to show that a deferred func runs after the func containing it exits.

### EX4
1. Create a user defined struct with
- the identifier “person”
- the fields:
  - first
  - last
  - age
2. attach a method to type person with
- the identifier “speak”
- the method should have the person say their name and age
3. create a value of type person
4. call the method from the value of type person

### EX5
1. create a type SQUARE
2. create a type CIRCLE
3. attach a method to each that calculates AREA and returns it
- circle area= π r 2
- square area = L * W
4. create a type SHAPE that defines an interface as anything that has the AREA method
5. create a func INFO which takes type shape and then prints the area
6. create a value of type square
7. create a value of type circle
8. use func info to print the area of square
9. use func info to print the area of circle

### EX6
Build and use an anonymous func

### EX7
Assign a func to a variable, then call that func

### EX8
- Create a func which returns a func
- assign the returned func to a variable
- call the returned func

### EX9
A “callback” is when we pass a func into a func as an argument. For this exercise,
- pass a func into a func as an argument

### EX10
Closure is when we have “enclosed” the scope of a variable in some code block. For this hands-on exercise, create a func which “encloses” the scope of a variable.