# telegram-bot-scripts
A telegram bot to run predefined shell scripts

simplely put your scripts (WHITOUT .sh extension) to ./scripts, and then run:

`telegram-bot-scripts -key aaa -users 111:222`


```
Usage of telegram-bot-scripts
  -key string
        telegram bot api key
  -shell string
        shell to run script (default "sh")
  -users string
        telegram users ids to use this bot, like: 123:456:789
  -verbose
        verbose output (default true)

```
