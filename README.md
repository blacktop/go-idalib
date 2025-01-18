# go-idalib

> Go bindings for the IDA 9.0 SDK via [idalib](https://docs.hex-rays.com/user-guide/idalib)

## Why? 🤔

Everyone deserves bindings, even gophers!

> [!WARNING]
> **NEED HELP**  
> With fleshing out the rest of the objects like Function etc and library C++ wrappers to get to a MVP.


## Getting Started

### Requirements

- IDA Pro 9.0 or newer installation (with a valid license)

```bash
go get github.com/blacktop/go-idalib
```

### `ida` demo tool

Install

```bash
go install github.com/blacktop/go-idalib/cmd/ida@latest
```

Usage

```bash
❯ ida path/to/i64
```
```bash
   • Function count: 22140
   • 1st Function: 0x15a87c2f0
```
```cpp
Decompiled Function:

int __fastcall memset_s(void *s, size_t smax, int c, size_t n)
{
  size_t v4; // x8
  int v5; // w19

  if ( !s )
    return 22;
  if ( (smax & 0x8000000000000000LL) != 0 )
    return 7;
  if ( n >= smax )
    v4 = smax;
  else
    v4 = n;
  if ( n <= smax )
    v5 = 0;
  else
    v5 = 84;
  memset(s, c, v4);
  return v5;
}
```

Print Version

```bash
❯ ida version
   • IDA library version: 9.0.241217
```

## Known Issues

> [!WARNING]
> [idalib](https://docs.hex-rays.com/user-guide/idalib) cannot function correctly when not running on the same thread that you initialized the library on.  Right now I am not doing anything to ensure Go doesn't create/switch threads so it will 'freeze' if you do anything too long/complex.

## Development

### Requirements

- [swig](https://www.swig.org)

Unzip IDA 9.0 SDK in root directory to build

```bash
make build
```

## License

MIT Copyright (c) 2025 **blacktop**