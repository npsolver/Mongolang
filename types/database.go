package types

type Database struct {
	collection Collection
}

func (db *Database) getCollection(name Name) *Database {
	db.collection = Collection{name}
	return db
}

func init() {
	db := Database{}

	name := Name{"newCollection"}
	
	db.getCollection(name)
}

