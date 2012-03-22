include $(GOROOT)/src/Make.inc

TARG=vodka
GOFILES=\
	parser.go\
	stack.go\
	vodka.go\

include $(GOROOT)/src/Make.pkg
