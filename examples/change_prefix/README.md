# Using a prefix in the environment variables

Usually your program is not running alone, it works along with other parts of the system and it is much better to have a prefix in the environment variables so you can differentiate between your programs.

To do this just add the following line before calling the parser

```go
config.PrefixEnv = "EXAMPLE"
```
