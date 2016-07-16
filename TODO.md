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
