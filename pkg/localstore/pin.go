package localstore

import (
	"errors"

	"github.com/wssll789/bee/pkg/shed"
	"github.com/wssll789/bee/pkg/storage"
	"github.com/wssll789/bee/pkg/swarm"
	"github.com/syndtr/goleveldb/leveldb"
)

// pinCounter returns the pin counter for a given swarm address, provided that the
// address has been pinned.
func (db *DB) pinCounter(address swarm.Address) (uint64, error) {
	out, err := db.pinIndex.Get(shed.Item{
		Address: address.Bytes(),
	})

	if err != nil {
		if errors.Is(err, leveldb.ErrNotFound) {
			return 0, storage.ErrNotFound
		}
		return 0, err
	}
	return out.PinCounter, nil
}
