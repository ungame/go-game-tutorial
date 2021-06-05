# Golang SDL2

https://github.com/veandco/go-sdl2

## Instalar Mingw-W64

- https://sourceforge.net/projects/mingw-w64/

- Acesse o `C:` e crie uma pasta `mingw64`

- Extraia os arquivos baixados dentro de `C:\mingw64`

- Adicione o caminho `C:\mingw64\bin` ao __Path__ do Sistema:

- Adicione o caminho `C:\mingw64\x86_64-w64-mingw32\bin` ao __Path__ do Sistema.

- Abra o terminal e teste os comandos:

```cmd
    g++ --version
```

A saída deve ser semelhante a esta:

```cmd
g++.exe (x86_64-posix-seh-rev0, Built by MinGW-W64 project) 8.1.0
Copyright (C) 2018 Free Software Foundation, Inc.
This is free software; see the source for copying conditions.  There is NO
warranty; not even for MERCHANTABILITY or F
```

```cmd
    gcc --version
```

A saída deve ser semelhante a esta:

```cmd
gcc.exe (x86_64-posix-seh-rev0, Built by MinGW-W64 project) 8.1.0
Copyright (C) 2018 Free Software Foundation, Inc.
This is free software; see the source for copying conditions.  There is NO
warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
```

```cmd
    ld --version
```

A saída deve ser semelhante a esta:

```cmd
GNU ld (GNU Binutils) 2.30
Copyright (C) 2018 Free Software Foundation, Inc.
This program is free software; you may redistribute it under the terms of
the GNU General Public License version 3 or (at your option) a later version.
This program has absolutely no warranty.
```
