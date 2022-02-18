Hi, Curious Being! 👋

The program you see in this repo is a simple command line tool(CLI) built in Golang. I heard Go makes it super easy to create custom CLIs. Gave it a shot.

## What can I do with StringWizard CLI?

You can directly input a text as a command in the terminal window and evaluate it on the basis of certain metrics. 

Example:
```
$ stringparse list --text "Hackers beware." --metric words --unique
textPtr: Hackers beware., metricPtr: words, uniquePtr: true

$ stringparse count --text "She sells sea shells by the sea shore" --metric substring --substringList se,sh,ea,he
textPtr: She sells sea shells by the sea shore, metricPtr: substring, substringPtr: , substringListPtr: [se sh ea he], uniquePtr: false

$ stringparse count --text "Komand Security"
textPtr: Komand Security, metricPtr: chars, substringPtr: , substringListPtr: [], uniquePtr: false
```