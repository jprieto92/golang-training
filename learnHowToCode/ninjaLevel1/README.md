# golang-training
golang exercises of the course " Learn How To Code: Google's Go (golang) Programming Language"

## Ninja level 1

### EX1
1. Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”

   a. 42

   b. “James Bond”

   c. true
2. Now print the values stored in those variables using

   a. a single print statement

   b. multiple print statements

### EX2
1. Use var to DECLARE three VARIABLES. The variables should have package level  scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the  variables and make sure the variables are of the following TYPE (meaning they can  store VALUES of that TYPE).

   a. identifier “x” type int

   b. identifier “y” type string

   c. identifier “z” type bool
2. in func main

   a. print out the values for each identifier

   b. The compiler assigned values to the variables. What are these values called?

### EX3
Using the code from the previous exercise,
1. At the package level scope, assign the following values to the three variables

   a. for x assign 42

   b. for y assign “James Bond”

   c. for z assign true
2. in func main
   
   a. use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”

   b. print out the value stored by variable “s”

### EX4
● FYI - nice documentation and new terminology “underlying type”

○ https://golang.org/ref/spec#Types

For this exercise
1. Create your own type. Have the underlying type be an int.
2. create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
3. in func main

   a. print out the value of the variable “x”

   b. print out the type of the variable “x”

   c. assign 42 to the VARIABLE “x” using the “=” OPERATOR

   d. print out the value of the variable “x”

### Ex5
Building on the code from the previous example
1. at the package level scope, using the “var” keyword, create a VARIABLE with the IDENTIFIER “y”. The variable should be of the UNDERLYING TYPE of your custom TYPE “x”

type hotdog int
var x hotdog
var y int

2. in func main

   a. this should already be done

      i. print out the value of the variable “x”

      ii. print out the type of the variable “x”

      iii. assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR

      iv. print out the value of the variable “x”

   b. now do this

      i. now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
         
         - then use the “=” operator to ASSIGN that value to “y” 

         - print out the value stored in “y”

         -print out the type of “y”


