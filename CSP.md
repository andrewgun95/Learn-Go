# Intro to CSP (Communicating Sequential Processes)

## Test Case
Two function - A and B
<pre>
A >>>   >>> Finish A
B    >>>
</pre>

A will do something that need for B to do the work
B required for the A to continue the work

**Think about two things - Communiction and Synchronization**
1. Data to pass back and forth
2. Run at specific orders

## Brute Force

<pre>
A(data) {
    value <- wait B(data)
    // show the data
}

B(data) {
    // calculate the data
    return result
}
</pre>

### Minus
1. Direct call of B inside A - tight couple - can end up giant net of calls which difficult to maintain and extend

## The Mediator Pattern

<pre>A | B</pre>
which is | as mediator

<pre>
Mediator(data) {
    value <- wait B(data)
    A(data)
}

A(data) {
    // show the data
}

B(data) {
    // calculate the data
    return result
}
</pre>

### Minus
1. A lot of things inside mediator, become complex and mediator need know a lot of things in the system

## Pub Sub Pattern
Gives no direct reference to entities in the system

<pre>
              BUS
publishers -> [][][] -> subscribers
</pre>

<pre>
A(data) {
    BUS.subscribe(1, () -> {
        // show the data
    })
    BUS.publish(0, data)
}

B() {
    BUS.subscribe(0, data -> {
        // calculate the data
        BUS.publish(1, result)
    })
}

A(data)
B()
</pre>

### Minus
1. Both solution only **solve the communication problem** but **no the synchronize problem**

<pre>
IDEA : Wouldn't be nice if we just run A and B in a random order and there is a mechanism which synchronizes them. 

Well, there is 👉 communicating sequential processes.
</pre>

## Channel Pattern - using CSP
PubSub pattern but with many buses - call them as a Channel
<pre>
       channel 1
put -> [][][] -> take
       channel 2
put -> [][][] -> take
</pre>
Rules :
1. Can't put the channel if there is no one to take it
2. Can't take it unless there is someone to put it

## Articles :
1. https://krasimirtsonev.com/blog/article/we-need-channels-intro-to-csp#communicating-sequential-processes-csp
2. https://katcipis.github.io/blog/mux-channels-go/