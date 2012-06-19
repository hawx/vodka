# Range

A range should be creatable like in ruby, only have inclusive ranges to begin
with, it's less confusing,

    1..5

with nice things like,

    1..5 max   ;=> 5
    1..5 min   ;=> 1

    1..5 list  ;=> (1 2 3 4 5)

    1 5 range  ;=> (1 2 3 4 5)


# Dictionaries and Relations

A relation looks like

    {'a' -> 1}

A dictionary is just multiple relations, or a relation is a dictionary with one
entry?

    {'a' -> 1, 'b' -> 2}

Commas are optional allowing

    {
      'a' -> 1
      'b' -> 2
      'c' -> 3
    }

Should be able to

    {'a' -> 1} {'b' -> 2} merge          ;=> {'a' -> 1, 'b' -> 2}

    1 'a' relate                         ;=> {'a' -> 1}
    (1 2 3) ('a' 'b' 'c') relate-pairs   ;=> {'a' -> 1, 'b' -> 2, 'c' -> 3}

    {'a' -> 1, 'b' -> 2} 'a' get         ;=> 1


# Better list implementation

Maybe based on a gap list?


# Fluid Number Type

Number should be a generic wrapper for; integer, quotient and float/bignum
(undecided which to use at moment) and complex. So I can do

    2            ; integer
    2/3          ; quotient
    2.3          ; real
    c(2 3i)      ; complex

    2 2/3 +      ;=> 8/3
    2 2.3 +      ;=> 4.3
    c(2 3i) 5 +  ;=> c(7 3i)

I should be able to use very large numbers

    ; googol
    100 10 ^        ;=> 100000000000000000...
    ; googolplex
    100 10 ^ 10 ^   ;=> Infinity

    ; scientific notation
    10e100          ;=> 100000000000000000...

    ; zero over zero
    0 0 /           ;=> NaN

These won't be no IEEEE complient numbers!
