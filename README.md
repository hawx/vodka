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
>> quit
$
```

[cc]: http://github.com/hawx/catcon
