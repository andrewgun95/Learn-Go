## Question ?

1. What happen if there are packages in sub-directories under root module (directory with go.mod in it) ?
It's a nested package and use a relative path from the root package name to import it
Example :
imports parent/child

2. What is idiomatic go ? has three design principle
    * Orthogonality
        In programming terms "orthogonality" means that pieces are independent from each other. Changes to one part, a type, package, program, etc, have minimal to no effect on other parts.
    * Simplicity
        Also often referred to as "reduced complexity", Go forgoes many of the complicated "features" found in other languages.
    * Readability
        Go reduces clutter and noise. There are no header files, unsurprising syntax, and everything can only be declared once per block.

3. Is realy go command is not working outside %GO_PATH% if no go.mod in it ?

4. Is Go an object-oriented language ?
    Go is Object Oriented
    1. Encapsulation
        1. state ("fields")
        2. behavior ("methods")
        3. exported & unexported; viewable & not viewable
    2. Reusability
        1. inheritance ("embedded types")
    3. Polymorphism
        1. interfaces
    4. Overriding
        1. "promotion"

    Traditional OOP
    1. Classes
        1. data structure describing a type of object
        2. you can then create "instances"/"objects" from the class / blueprint
        3. classes hold both:
            1. state / data / fields
            2. behavior / methods
        4. public / private
    2. Inheritance

    In Go:
    you don't create classes, you create a **TYPE**
    you don't instantiate, you create a **VALUE** of a **TYPE**

## Coding Convention

1. Convention: logically organize your fields together. 
2. Readability & clarity trump performance as a design concern. 
"Go will be performant. Go for readability first." 
However, if you are in a situation where you 
need to prioritize performance: lay the fields out from largest to smallest, eg, int 64, int64, float32, bool

## Others
Useful reads
- https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc
- https://about.sourcegraph.com/go/idiomatic-go/
- https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827
- https://golang.org/doc/articles/wiki/