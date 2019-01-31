# base58

This is a CMD is a Go language for encoding/decoding base58 strings.

## Usage

```bash
$ echo 123 | base58
2FwFmj
$ echo 123 | base58 | base58 -d
123
```
