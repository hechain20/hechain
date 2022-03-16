<!---
 File generated by help_docs.sh. DO NOT EDIT.
 Please make changes to preamble and postscript wrappers as appropriate.
 --->

# cryptogen

`cryptogen` is an utility for generating Hechain key material.
It is provided as a means of preconfiguring a network for testing purposes.
It would normally not be used in the operation of a production network.

## Syntax

The ``cryptogen`` command has five subcommands, as follows:

  * help
  * generate
  * showtemplate
  * extend
  * version

## cryptogen help
```
usage: cryptogen [<flags>] <command> [<args> ...]

Utility for generating Hechain key material

Flags:
  --help  Show context-sensitive help (also try --help-long and --help-man).

Commands:
  help [<command>...]
    Show help.

  generate [<flags>]
    Generate key material

  showtemplate
    Show the default configuration template

  version
    Show version information

  extend [<flags>]
    Extend existing network
```


## cryptogen generate
```
usage: cryptogen generate [<flags>]

Generate key material

Flags:
  --help                    Show context-sensitive help (also try --help-long
                            and --help-man).
  --output="crypto-config"  The output directory in which to place artifacts
  --config=CONFIG           The configuration template to use
```


## cryptogen showtemplate
```
usage: cryptogen showtemplate

Show the default configuration template

Flags:
  --help  Show context-sensitive help (also try --help-long and --help-man).
```


## cryptogen extend
```
usage: cryptogen extend [<flags>]

Extend existing network

Flags:
  --help                   Show context-sensitive help (also try --help-long and
                           --help-man).
  --input="crypto-config"  The input directory in which existing network place
  --config=CONFIG          The configuration template to use
```


## cryptogen version
```
usage: cryptogen version

Show version information

Flags:
  --help  Show context-sensitive help (also try --help-long and --help-man).
```

## Usage

Here's an example using the different available flags on the ``cryptogen extend``
command.

```
    cryptogen extend --input="crypto-config" --config=config.yaml

    org3.example.com
```

Where config.yaml adds a new peer organization called ``org3.example.com``

<a rel="license" href="http://creativecommons.org/licenses/by/4.0/"><img alt="Creative Commons License" style="border-width:0" src="https://i.creativecommons.org/l/by/4.0/88x31.png" /></a><br />This work is licensed under a <a rel="license" href="http://creativecommons.org/licenses/by/4.0/">Creative Commons Attribution 4.0 International License</a>.