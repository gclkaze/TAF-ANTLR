cwd=$(PWD)
all:
	java -jar antlr-4.13.1-complete.jar -Dlanguage=Go -o parser grammar/tafexpr.g4