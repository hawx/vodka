package table

const BOOT string = `

'__document__' . define

; Evaluates the string at the top of the stack as vodka source code.
; sig: string ->
; example: '1 2 +' eval ;=> [3]
'eval' __document__

; Defines an alias which when used will call the aliased function.
; sig: string string ->
; example: 'add' '+' alias
'alias' __document__

; Defines a new function with the name given, when called the block passed will
; be called.
; sig: string block ->
; example: 'sq' [ dup * ] define
'define' __document__

; Pushes a string containing all defined functions on to the stack.
; sig: -> string
'defined' __document__

; Replaces the element at the top of the stack with a string describing it's
; type.
; sig: 'a -> string
; example: 123 type ;=> ['integer']
'type' __document__

; Checks whether there are no elements on the stack.
; sig: -> bool
'empty?' [
  size 0 eq?
] define

; Checks whether the stack contains one or less elements.
; sig: -> bool
'small?' [
  2 size lte?
] define

; Checks whether the top element is nil.
; sig: 'a -> bool
'nil?'     [ type 'nil'     = ] define

; Checks whether the top element is boolean.
; sig: 'a -> bool
'boolean?' [ type 'boolean' = ] define

; Checks whether the top element is a block.
; sig: 'a -> bool
'block?'   [ type 'block'   = ] define

; Checks whether the top element is a string.
; sig: 'a -> bool
'string?'  [ type 'string'  = ] define

; Checks whether the top element is an integer.
; sig: 'a -> bool
'integer?' [ type 'integer' = ] define

; Includes the vodka file at the top of the stack.
; sig: string ->
; example: 'test' include ; Will load the file 'test.vk'
'include' [
  '.vk' concat read eval
] define


; group: Types

; Replaces the string at the top of the stack with an integer of the value
; represented by the string.
; sig: string -> int
; example: '5' integer ;=> [5]
'integer' __document__

; Converts the element at the top of the stack into a string.
; sig: 'a -> string
; example: 5 string ;=> ['5']
'string' __document__

; Replaces the element at the top of the stack with a list containing the
; element.
; sig: 'a -> list
; example: 5 list ;=> [(5)]
'list' __document__

; Creates a list with all the numbers from the second element to the top
; element.
; sig: int int -> list
; example: 1 5 range ;=> [(1 2 3 4 5)]
'range' __document__


; group: Basic I/O

; Prints the string at the top of the stack to stdout.
; sig: string ->
; example: 'Hello world' print ;=> Hello World
'print' __document__

; Prints the string representation of the element at the top of the stack.
; sig: 'a ->
; example: 500 p ;=> 500
'p' __document__

; Reads the file at the top of the stack.
; sig: string -> string
; example: 'lorem.txt' read ;=> ["Lorem ..."]
'read' __document__

; group: Stack operations

; Removes the element from the top of the stack.
; sig: 'a ->
'pop' __document__

; Pushes the size of the stack on to the top of the stack.
; sig: -> int
'size' __document__

; Pushes a copy of the top element to the stack.
; sig: 'a -> 'a 'a
'dup' __document__

; Swaps the top two elements on the stack.
; sig: 'a 'b -> 'b 'a
'swap' __document__

; Removes all elements from the stack.
; sig: 'A ->
'drop' __document__

; Composes the two blocks at the top of the stack, forming one single block.
; sig: block block -> block
; example: [ dup dec ] :mult compose ;=> [[ dup dec mult ]]
'compose' __document__

; Wraps the block in another block.
; sig: block -> ( -> block)
; example: :inc wrap [[inc]] eq? ;=> [true]
'wrap' __document__

; Duplicates the top two elements of the stack.
; sig: 'a 'b -> 'a 'b 'a 'b
'dup2' [ swap dup swapp swap dup swapp ] define

; Swaps the second and third elements on the stack.
; sig: 'a 'b 'c -> 'b 'a 'c
'swapp' [ :swap under ] define

; Calls the block on the top of the stack for each element on the stack,
; should be used with a reductive function.
; sig: 'A block -> 'a
'stack' [ 3 size - times ] define

; group: Arithmetic operations

; Adds the top two elements of the stack together.
; sig: int int -> int
; example: 5 2 add ;=> [7]
'add' __document__

; Multiples the top two elements of the stack together.
; sig: int int -> int
; example: 5 2 mult ;=> [10]
'mult' __document__

; Subtracts the top element of the stack from the second.
; sig: int int -> int
; example: 5 2 sub ;=> [3]
'sub' __document__

; Divides the second element of the stack by the second. Does integer division
; only at the moment, this does not round to the nearest integer.
; sig: int int -> int
; example: 5 2 div ;=> [2]
'div' __document__

; Negates the number at the top of the stack.
; sig: int -> int
; example: 5 neg ;=> [-5]
'neg' __document__

'add'  '+' alias
'mult' '*' alias
'sub'  '-' alias
'div'  '/' alias

; Sums all values on the stack.
; sig: 'A -> int
'sum' [
  :+ stack
] define

; Multiplies all values on the stack.
; sig: 'A -> int
'prod' [
  :* stack
] define

; Increments the number at the top of the stack.
; sig: int -> int
'inc' [ 1 add ] define

; Decrements the number at the top of the stack.
; sig: int -> int
'dec' [ 1 swap sub ] define

'inc' '++' alias
'dec' '--' alias

; Checks whether the top element is 0.
; sig: 'a -> bool
'zero?' [
  0 eq?
] define

; Replaces the number at the top of the stack with it's absolute value.
; sig: int -> int
'abs' [
  dup 0 lt?
  :neg unless
] define

; Calculates the factorial of the number at the top of the stack.
; sig: int -> int
; example: 5 fact ;=> [120]
'fact' [
  1 swap range :* reduce
] define


; group: Logical operations

; Calculates the value of a boolean or operation applied to the top two elements
; of the stack.
; sig: bool bool -> bool
; example: true false or ;=> [true]
'or' __document__

; Calculates the value of a boolean and operation applied to the top two
; elements of the stack.
; sig: bool bool -> bool
; example: true false and ;=> [false]
'and' __document__

; Compares the top two elements of the stack, pushes -1 if the top element is
; less than the second, 0 if the elements are equal or 1 if the top element is
; greater.
; sig: 'a 'b -> int
; example: 1 2 compare ;=> [1]
'compare' __document__

; Checks if the top two elements are equal.
; sig: 'a 'b -> bool
; example: 1 1 eq? ;=> [true]
'eq?' __document__

; Returns the opposite of the boolean value at the top of the stack.
; sig: bool -> bool
'not' [
  :true :false if-else
] define

; Checks whether the top element is greater than the second.
; sig: 'a 'b -> bool
; example: 1 2 gt? ;=> [true]
'gt?'  [ compare  1 eq? ] define

; Checks whether the top element is less than the second.
; sig: 'a 'b -> bool
; example: 1 2 lt? ;=> [false]
'lt?'  [ compare -1 eq? ] define

; Checks whether the top element is not equal to the second. Note: this has a
; different meaning to 'eq? not'. This will return false if the elements are
; lt? or gt? or eq?. Make sure you understand this properly before using, in
; most cases you will want to use 'eq? not' instead.
;
; sig: 'a 'b -> bool
; example: 'a' 4 neq? ;=> [true]
'neq?' [ compare -2 eq? ] define

'eq?'  '=' alias
'gt?'  '>' alias
'lt?'  '<' alias
'not'  '!' alias

; Checks whether the top element is less than or equal to the second.
; sig: 'a 'b -> bool
'lte?' [
  gt? !
] define

; Checks whether the top element is greater than or equal to the second.
; sig: 'a 'b -> bool
'gte?' [
  lt? !
] define

'lte?' '<=' alias
'gte?' '>=' alias


; group: Control flow

; Calls the block at the top of the stack if the condition is true, or the
; second if the condition is false.
; sig: bool block block -> 'a
; example: 300 300 eq? [ "yes" ] [ "no" ] if-else ;=> ["yes"]
'if-else' __document__

; Calls the block at the top of the stack.
; sig: block -> 'a
; example: 300 :inc call ;=> [301]
'call' __document__

; Calls the second element while temporarily removing the top element from the
; stack.
; sig: ( -> 'A) 'a -> 'A 'a
; example: 5 :dec 'In the way' without ;=> [4 'In the way']
'without' __document__

; Calls the third element while temporarily removing the top two elements from
; the stack.
; sig: ( -> 'A) 'a 'b -> 'A 'a 'b
'without2' __document__

; Calls the top element while temporarily removing the second element.
; sig: 'a ( -> 'A) -> 'a 'A
; example: 5 'In the way' :dec under ;=> [4 'In the way']
'under' [
  swap without
] define

; Calls the top element while temporarily removing the second and third
; elements.
; sig: 'a 'b ( -> 'A) -> 'a 'b 'A
'under2' [
  swap swapp without2
] define

; Evaluates the block at the top of the stack if the value below it is true.
; sig: bool block -> 'a
; example: 300 300 eq? [ "300 equals 300, crazy!" print ] if
'if' [
  . swap if-else
] define

; Evaluates the block at the top of the stack if the value below it is false.
; sig: bool block -> 'a
'unless' [
  swap not swap if
] define

; Evalutes the second block on the third element of the stack while the
; condition at the top of the stack is true.
; sig: block (block -> bool) -> 'a
; example: 5 :dec [ zero? ! ] while ;=> [0]
'while' __document__

; Evaluates the second block while the condition at the top of the stack is
; false.
; sig: block ( -> bool) -> 'a
; example: 5 :dec :zero? until ;=> 0
'until' [
  :not compose while
] define

; Calls the second element of the stack n times, where n is the top element.
; sig: block int -> 'a
; example: 300 :inc 5 times ;=> [305]
'times' [
  swap
  wrap :dec swap compose :under compose
  [ dup zero? ]
  ; 300 5 [ dec :inc under ] [ dup zero? ] until
  until

  ; Cleanup counter
  pop nil pop
] define


; group: Strings

; Joins the top two strings on the stack forming one string.
; sig: string string -> string
; example: 'Hello' ' World' concat ;=> ['Hello World']
'concat' __document__

; Joins all strings in a list together using the string at the top.
; sig: list string -> string
; example: ('John' 'Dave' 'Luke') ', ' join ;=> ['John, Dave, Luke']
'join' [
  wrap [ swap concat concat ] compose reduce
] define


; group: Lists

; Replaces the list with it's first element.
; sig: list -> 'a
; example: (1 2 3 4) head ;=> [1]
'head' __document__

; Replaces the list with a list containing all but the first of it's elements.
; sig: list -> list
; example: (1 2 3 4) tail ;=> [(2 3 4)]
'tail' __document__

; Adds the element at the top of the stack to the end of the list below.
; sig: list 'a -> list
; example: (1 2 3) 4 cons ;=> [(1 2 3 4)]
'cons' __document__

; Adds the second list on the stack to the back of the first.
; sig: list list -> list
; example: (1 2 3) (4 5 6) append ;=> [(4 5 6 1 2 3)]
'append' __document__

; Adds the second list on the stack to the front of the first.
; sig: list list -> list
; example: (1 2 3) (4 5 6) prepend ;=> [(1 2 3 4 5 6)]
'prepend' [ swap append ] define

; Applies the block at the top of the stack to the elements in the list below.
; sig: list block -> list
; example: (1 2 3) :swap apply ;=> [(1 3 2)]
'apply' __document__

; Reverses the elements in the list.
; sig: list -> list
; example: (1 2 3) reverse ;=> [(3 2 1)]
'reverse' __document__

; Applies the function to each pair of elements in the list.
; sig: list block -> 'a
; example: (1 2 3 4) :+ reduce ;=> [10]
'reduce' [
  wrap :stack compose apply head
] define

; Applies the function to each element in the list.
; sig: list block -> list
'map' [
  ()
  swap swapp

  [
    [ dup head swap ] under
    :tail under
    dup
    swapp
    [ call cons ] under2
  ] [
    () swap swapp [ dup := under ] under swapp swap
  ] until

  pop pop
] define
`
