# Vodka

A functional, stack-based, concatenative programming language of little
practical use. It is based on my earlier [Catcon][cc], but this is written in Go
and features a few syntax changes.


## Building

To get up and running with a REPL run,

``` bash
$ cd place-download-to
$ gomake && 6l _go_.6 && mv 6.out vodka
$ ./vodka
Vodka REPL, CTRL+C or type 'quit' to quit
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
>>
```

### Integers

``` bash
>> 5 6 1000
[5 6 1000] => nil
>>
```

### Boolean & Nil

``` bash
>> true false nil
[true false nil] => nil
>>
```

### Blocks

Note the spaces around the square brackets are necessary,

``` bash
>> [ dup * ]
[[dup *]] => nil
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
