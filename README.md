# geth-contract-state-export

A command line tool to parse geth chaindata and extract the key value store for specific smart contract.

Usage:

```
go run main.go <LevelDB path> <account address>
```

Example:

```
go run main.go /path/to/chaindata ff970a61a04b1ca14834a43f5de4533ebddb5cc8
```

Output:

```
0x00000a7d1b90cce0892b47dbe294dd3bd9d4b4ff3dc238d9836068e6b0f9504f: 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffff88559c86
0x00001c9389ec58770864cae103371fc26200caefad72deb05545ba206b431c72: 0x0a9ab3
0x0000227638f0db7fce130a9e3dccfc6c65f98e340f448eaa30e5ab7635084fe6: 0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffff1704196
0x00004864222ab7d987191ee1d480c45382882a3081bf7267234ade456908e158: 0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffffb59bb89
...
```

Note:

The code here is a modified version [this](https://github.com/MartiTM/geth-leveldb-explorer)
