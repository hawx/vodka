include $(GOROOT)/src/Make.inc

TARG=vodka
GOFILES=\
	parser.go\
	table.go\
	stack.go\
	interpreter.go\
	vodka.go\

include $(GOROOT)/src/Make.pkg
