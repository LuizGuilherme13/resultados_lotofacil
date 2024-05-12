# Gerador de apostas Lotofácil

### 1. Clone o projeto

```bash
git clone github.com/LuizGuilherme13/resultados_lotofacil
```

```bash
cd resultados_lotofacil
```

```bash
go build -o loto
```

### 2. Como usar

Busca os resultados dos sorteios

```bash
./loto get
```

Busca os resultados dos sorteios de acorda com o período passado

```bash
./loto get i- dd/mm/yyyy -f  dd/mm/yyyy
```

Exibe os resultados no terminal

```bash
./loto see
```

Gera uma aposta de acordo com os números mais sorteados de acordo com os resultados buscados.

```bash
./loto generate
```

Gera a quantidade informada em -q

```bash
./loto generate -q 3
```
