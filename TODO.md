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


# Some kind of type system

I don't really like type systems that much, but an optional type system could be
useful. And it shouldn't be too difficult to infer types since it is pretty
simple for a stack.

I also need some kind of polymorphic/multi-dispatch system, for instance

    'add5' [
      5 +
    ] Number define-for

    'add5' [
      '5' concat
    ] String define-for

    'add5' [
      5 cons
    ] List define-for

Or even nicer

    'add5' [
      [ 5 + ]
      [ 5 'concat' ]
      [ 5 'cons' ]
    ] multi-define

And use type inference on the stack and functions to determine which to use.


# VSpec

Some kind of test framework.

    '+' [
      'add positive numbers' [
        1 2 +
        3 must-equal
      ] can

      'add negative numbers' [
        -1 -2
        -4 must-equal
      ] can
    ] describe

    $ vodka adder_spec.vk

    +
    . add positive numbers
    x add negative numbers

    FAIL: 1 pass / 1 fail

Then I can make sure vodka actually works against what I say it does... I should
probably write go tests though.
