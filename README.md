# Gerador de apostas Lotofácil

### 1. Instalando

```bash
go install github.com/LuizGuilherme13/resultados_lotofacil@latest
```

### 2. Como usar

```bash
cd ~/go/bin
```

Busca os resultados dos sorteios

```bash
./resultados_lotofacil get
```

Busca os resultados dos sorteios de acorda com o período passado

```bash
./resultados_lotofacil get i- dd/mm/yyyy -f  dd/mm/yyyy
```

Exibe os resultados no terminal

```bash
./resultados_lotofacil see
```

Gera uma aposta de acordo com os números mais sorteados de acordo com os resultados buscados.

```bash
./resultados_lotofacil generate
```

Gera a quantidade informada em -q

```bash
./resultados_lotofacil generate -q 3
```
