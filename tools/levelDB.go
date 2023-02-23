package tools

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
	"math/big"
)

func ContractState(ldbPath string, addr string) {
	ldb := getLDB(ldbPath)
	stateRootNode := getStateTrees(ldb, 1)[0].stateRoot
	getStateForContract(ldb, stateRootNode, addr)
}

/*
==================================================================================================================================
*/

type stateFound struct {
	blockNumber *big.Int
	stateRoot   common.Hash
}

func getStateTrees(ldb ethdb.Database, stopAt int) []stateFound {
	var res []stateFound
	headerHash, _ := ldb.Get(headHeaderKey)
	for headerHash != nil {
		// print the header hash
		var blockHeader types.Header
		blockNb, _ := ldb.Get(append(headerNumberPrefix, headerHash...))
		if blockNb == nil {
			break
		}
		blockHeaderRaw, _ := ldb.Get(append(headerPrefix[:], append(blockNb, headerHash...)...))
		rlp.DecodeBytes(blockHeaderRaw, &blockHeader)

		stateRootNode, _ := ldb.Get(blockHeader.Root.Bytes())

		if len(stateRootNode) > 0 {
			res = append(res, stateFound{blockHeader.Number, blockHeader.Root})
			if stopAt > 0 && len(res) == stopAt {
				return res
			}
		}

		headerHash = blockHeader.ParentHash.Bytes()
	}

	return res
}

func getStateForContract(ldb ethdb.Database, stateRootNode common.Hash, addr string) {

	trieDB := trie.NewDatabase(ldb)
	treeState, _ := trie.New(stateRootNode, trieDB)

	addrHash := crypto.Keccak256Hash(common.Hex2Bytes(addr))

	addrState := treeState.Get(addrHash.Bytes())
	var values [][]byte
	if err := rlp.DecodeBytes(addrState, &values); err != nil {
		panic(err)
	}

	// decoded value must be length 4
	// 0: nonce
	// 1: balance
	// 2: storage trie
	// 3: code hash

	// get the storage trie
	storageTrie, _ := trie.New(common.BytesToHash(values[2]), trieDB)

	it := trie.NewIterator(storageTrie.NodeIterator(nil))
	for it.Next() {
		var value []byte
		if err := rlp.DecodeBytes(it.Value, &value); err != nil {
			panic(err)
		}
		// print out he xencoded key and value
		fmt.Printf("0x%x: 0x%x\n", it.Key, value)
	}
}

func getLDB(ldbPath string) ethdb.Database {
	ldb, err := rawdb.NewLevelDBDatabase(ldbPath, 0, 0, "", true)
	if err != nil {
		fmt.Println("Did not find leveldb at path:", ldbPath)
		fmt.Println("Are you sure you are pointing to the 'chaindata' folder?")
		panic(err)
	}
	return ldb
}
