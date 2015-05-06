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
