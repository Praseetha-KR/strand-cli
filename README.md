# strand-cli

CLI tool to generate random strings using [strand](https://github.com/Praseetha-KR/strand) package


Supported random types:
- String
- Alpha Mumeric
- Alpha
- Numeric
- URL-safe
- Hexadecimal
- Binary Digits
- Password
- URL-safe Password
- String from custom character list


## Installation

### Binaries:

For installation instructions for binaries, please visit the [Releases Page](https://github.com/Praseetha-KR/strand-cli/releases).

### Build & install:

#### Build:

```
$ make
```

#### Install:

```
$ make install_<os>_<platform>
```

> eg: For OSX
>
> ```
> $ make install_darwin_amd64
> ```

Supported OS-Platform values are listed below:

<table>
    <thead>
        <tr>
            <th>OS</th>
            <th>Platform</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>darwin</td>
            <td>amd64</td>
        </tr>
        <tr>
            <td rowspan="4">linux</td>
            <td>amd64</td>
        </tr>
        <tr>
            <td>386</td>
        </tr>
        <tr>
            <td>arm64</td>
        </tr>
        <tr>
            <td>arm</td>
        </tr>
    </tbody>
</table>


## Usage

```bash
$ strand

            $$\                                        $$\
            $$ |                                       $$ |
 $$$$$$$\ $$$$$$\    $$$$$$\  $$$$$$\  $$$$$$$\   $$$$$$$ |
$$  _____|\_$$  _|  $$  __$$\ \____$$\ $$  __$$\ $$  __$$ |
\$$$$$$\    $$ |    $$ |  \__|$$$$$$$ |$$ |  $$ |$$ /  $$ |
 \____$$\   $$ |$$\ $$ |     $$  __$$ |$$ |  $$ |$$ |  $$ |
$$$$$$$  |  \$$$$  |$$ |     \$$$$$$$ |$$ |  $$ |\$$$$$$$ |
\_______/    \____/ \__|      \_______|\__|  \__| \_______|


NAME:
   strand - Random String Generator

USAGE:
   strand_darwin_amd64 [global options] command [command options] [arguments...]

VERSION:
   v0.0.3

AUTHOR:
   Praseetha KR <praseetha04@gmail.com>

COMMANDS:
   string, s         Generates a random string
   alpha, a          Generates a random string of letters
   alphanumeric, an  Generates a random alphanumeric string
   numeric, n        Generates a random number
   urlsafe, u        Generates a random URL safe string
   hex, h            Generates a random hexadecimal string
   binary, b         Generates a random string of binary digits
   password, p       Generates a password of random length
   from, f           Generates a random string from the given character list
   help, h           Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

---
### Examples



```bash
# --- Random string ---

$ strand string -l 25
xiZ3pOLN}JY!XaB).4`AQ3,[=


# --- Alpha-numeric with uppercase letters ---

$ strand alphanumeric -l 40 --upper
UW7LVTFEMT0BH2EVXCYW0ZB1AOP1QE4QQ2CUULSH


# --- Simple password ---

$ strand password --simple
6?U+msF8


# --- URL-safe password ---

$ strand password --urlsafe
ESQv7nP5H7w-_~_5L7CNF
```
