# GOreplace
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/andrejjovanovic/goreplace/Go)
![GitHub Releases](https://img.shields.io/github/downloads/andrejjovanovic/goreplace/0.2.0/total)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/andrejjovanovic/goreplace)

## What is GOreplace

GOreplace is a CLI application written in GO that allows you to replace placeholder tokens in configuration files. It is meant to be used with legacy applications that depend on having multiple or single configuration file with sensitive data. Commiting such an applications to git repositories and basing them on single branch aproaches includes a lot of seds in your deploy stages. GOreplace allows you to commit JSON file with your application and call it just once, efectively replacing all of your strings based on environment.

## Command line options

GOreplace expects three flags to be passed on call

-m Manifest file in JSON structure that will hold your configuration key:value\
-c Configuration file that you want to replace tokens on\
-e Named environment that exists within manifest file

## Manifest structure

```
{
  "develop":
  {
    "database":"wp-develop",
    "database_user":"devuser"
  },
  "integration":
  {
    "database":"wp-integration",
    "database_user":"intuser"
  }
}
```

Keys `develop` and `integration` marks the environment variable that needs to be passed to GOreplace. Key-values that are under the specific environment will be the only one to get replaced in destination file. Environments can be named however you want.

## Configuration files

Configuration files that you pass do not have any restrictions in type or format. Only thing that needs to be correctly done is that tokens for replacment has to be embedded in curly brackets and match the key that you have in manifest. `{database}`

## Executing the GOreplace

Following upper docs, GOreplace will work by simply executing 
`goreplace -c file.config -m manifest.json -e develop`
