/*
Copyright hechain. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package pvtstatepurgemgmt

func (pvtdataKeys *PvtdataKeys) addAll(toAdd *PvtdataKeys) {
	for ns, colls := range toAdd.Map {
		for coll, keysAndHashes := range colls.Map {
			for _, k := range keysAndHashes.List {
				pvtdataKeys.add(ns, coll, k.Key, k.Hash)
			}
		}
	}
}

func (pvtdataKeys *PvtdataKeys) add(ns string, coll string, key string, keyhash []byte) {
	colls := pvtdataKeys.getOrCreateCollections(ns)
	keysAndHashes := colls.getOrCreateKeysAndHashes(coll)
	keysAndHashes.List = append(keysAndHashes.List, &KeyAndHash{Key: key, Hash: keyhash})
}

func (pvtdataKeys *PvtdataKeys) getOrCreateCollections(ns string) *Collections {
	colls, ok := pvtdataKeys.Map[ns]
	if !ok {
		colls = newCollections()
		pvtdataKeys.Map[ns] = colls
	}
	return colls
}

func (colls *Collections) getOrCreateKeysAndHashes(coll string) *KeysAndHashes {
	keysAndHashes, ok := colls.Map[coll]
	if !ok {
		keysAndHashes = &KeysAndHashes{}
		colls.Map[coll] = keysAndHashes
	}
	return keysAndHashes
}

func newPvtdataKeys() *PvtdataKeys {
	return &PvtdataKeys{Map: make(map[string]*Collections)}
}

func newCollections() *Collections {
	return &Collections{Map: make(map[string]*KeysAndHashes)}
}
