# Vodka

A functional, stack-based, concatenative programming language of little
practical use. It is based on my earlier [Catcon][cc], but this is written in Go
and features a few syntax changes.


## Building

Install with `go get`,

``` bash
$ go get hawx.me/code/vokda
```

Then to start a REPL simply run,

``` bash
$ vodka
Vodka REPL, CTRL+C or type 'quit' to quit
>>
```

You can now type some commands, for example,

``` bash
>> 1 2 +
[3] => nil
>> 5 *
[15] => nil
>> pop
[] => 15
>>
```


## Quick Tutorial

### Strings

Strings can be created using either single (`'`) or double (`"`) quote marks.

``` bash
>> "Hello world!"
['Hello world!'"] => nil
>> " Yes." concat
['Hellow world! Yes.'] => nil
```

### Integers

Integers are the only numeric types currently available.

``` bash
>> 5 6 1000 -123
[5 6 1000 -123] => nil
>> + * -
[-5257] => nil
```

### Boolean & Nil

The boolean values are `true` and `false`. There is also the `nil` value.

``` bash
>> true false
[true false] => nil
>> and
[false] => nil
>> true
[false true] => nil
>> or
[true] => nil
>> nil
[true nil] => nil
```

### Lists

Lists are created by enclosing a list of values between parentheses. Spaces or
newlines are used to delimit values (you can use multiple if you wish).

``` bash
>> ('a' 'b' 'c')
[('a' 'b' 'c')] => nil
>> 'd' cons
[('a' 'b' 'c' 'd')] => nil
>> head
['a'] => nil
```

### Blocks

Blocks can contain any value or function. They are pushed to the stack
unevaluated so can be used in a similar way to anonymous functions in other
languages.

``` bash
>> [ dup * ]
[[ dup * ]] => nil
>> 5 swap
[5 [ dup * ]] => nil
>> call
[25] => nil
```

And we're done. So how about a slightly useful calculation showing some features
off,

``` bash
>> 'sq' [ dup * ] define
[] => nil
>> 4 sq
[16] => 16
>> sq sq
[65536] => 65536
>> drop
[] => nil
>> 'cb' [ dup sq * ] define
[] => nil
>> 4 cb
[64] => 64
>> pop
[] => 64
>> 'fib' [
..   ; first check whether 1 is less than top of the stack
..   dup 1 lt? [
..     ; F(n) = F(n-1) + F(n-2)
..     dup  1 - fib
..     swap 2 - fib
..     +
..     ] if
..   ] define
>> 1 fib
[1] => nil
>> 2 fib
[1 1] => nil
>> 3 fib
[1 1 2] => nil
>> 10 fib
[1 1 2 55] => nil
>> (1 1) [ dup2 + ] apply
[1 1 2 55 (1 1 2)] => nil
>> [[ dup2 + ] apply ] 5 times
[1 1 2 55 (1 1 2 3 5 8 13 21)] => nil
>> 'n-fib' [
..   ; check if the number at the top of the stack is less than 2
..   dup 2 lt?
..   ; if it is just return a list of 1s
..   [ () swap [ 1 cons ] swap times ]
..   ; otherwise calculate the fibonacci sequence
..   [
..     2 -
..     (1 1) swap
..     [[ dup2 + ] apply ] swap
..     times
..     ] if-else
..   ] define
>> drop 5 n-fib
[(1 1 2 3 5)] => nil
>> drop 20 n-fib
[(1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 1597 2584 4181 6765)] => nil
>> quit
$
```

[cc]: http://github.com/hawx/catcon
