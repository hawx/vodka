# Vodka

A functional, stack-based, concatenative programming language of little
practical use. It is based on my earlier [Catcon][cc], but this is written in Go
and features a few syntax changes.


## Building

To get up and running make sure you have Go 1 installed (`brew install go` with
homebrew). Then to start a REPL run,

``` bash
$ cd place-downloaded-to
$ go build
$ ./vodka
Vodka REPL, CTRL+C, or type 'quit' to quit
>>
```

You can now type commands, for example

``` bash
>> 1 2 +
[3] => 3
>>
```


## Quick Tutorial

The syntax is very simple, mainly due to the lack of types. In no particular
order,

### Strings

``` bash
>> "Hello world!"
["Hello world!"] => nil
```

### Integers

``` bash
>> 5 6 1000
[5 6 1000] => nil
```

### Boolean & Nil

``` bash
>> true false nil
[true false nil] => nil
```

### Lists

``` bash
>> ('a' 'b' 'c')
[('a' 'b' 'c')] => nil
```

### Blocks

``` bash
>> [ dup * ]
[[ dup * ]] => nil
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
..   dup dup 1 eq? swap 0 eq? or not [
..     dup 1 sub fib
..     swap 2 sub fib
..     add
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
