# cli_search
golang app with cmd support and clean architecture
Dependency
- Viper - for read config
- Cobra - support cli

This version only support cli, but easily support REST because of clean architecture
## Build
You can build from the source
```bash
go build main.go
```
or you just use the binary "main" file (mac,linux)
# usage
This app support 2 command [describe|search] and two object [user|ticket|organization]

Command describe
```bash
  ./main describe [user|ticket|organization]
```
Example: 
```bash
$ ./main describe user
$ ./main describe ticket
```
Command search
```bash
  ./main search [user|ticket|organization] field-name search-text
```
Example
```bash
$ ./main search organization name Enthaze
$ ./main search organization tag Fulton
$ ./main search user name "Loraine Pittman"
$ ./main search ticket subject "A Nuisance in Poland"
```

## Unsupported features:
Search empty data
Unit test
Mock data



