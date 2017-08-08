/*
 * MIT License
 *
 * Copyright (c) 2017 SmartestEE Co., Ltd..
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2017/08/08        Yang Chenglong
 */

package copy

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func PrepareCollection(session *mgo.Session, database string, collection string) (*mgo.Collection, chan struct{}) {
	shutdown := make(chan struct{})
	s := session.Copy()
	RefCollection := s.DB(database).C(collection)

	go func() {
		select {
		case <-shutdown:
			s.Close()
		}
	}()

	return RefCollection, shutdown
}

// IsValidObjectID check whether the id is a qualified ObjectId
func IsValidObjectID(id bson.ObjectId) bool {
	return IsValidObjectHex(id.Hex())
}

// IsValidObjectHex check whether the id can be converted to be a qualified ObjectId
func IsValidObjectHex(id string) bool {
	return bson.IsObjectIdHex(id)
}

// GetByID get a single record by ID
func GetByID(session *mgo.Session, database string, collection string, id string, i interface{}) {
	col, shutdown := PrepareCollection(session, database, collection)
	defer func() {
		close(shutdown)
	}()

	col.FindId(bson.ObjectIdHex(id)).One(i)
}

// GetUniqueOne get a single record by query
func GetUniqueOne(session *mgo.Session, database string, collection string, q interface{}, doc interface{}) error {
	col, shutdown := PrepareCollection(session, database, collection)
	defer func() {
		close(shutdown)
	}()

	return col.Find(q).One(doc)
}

// GetMany get multiple records based on a condition
func GetMany(session *mgo.Session, database string, collection string, q interface{}, doc interface{}) error {
	col, shutdown := PrepareCollection(session, database, collection)
	defer func() {
		close(shutdown)
	}()

	return col.Find(q).All(doc)
}

// Insert add new documents to a collection.
func Insert(session *mgo.Session, database string, collection string, doc interface{}) error {
	col, shutdown := PrepareCollection(session, database, collection)
	defer func() {
		close(shutdown)
	}()

	return col.Insert(doc)
}

// UpdateByQueryField modify all eligible documents.
func UpdateByQueryField(session *mgo.Session, database string, collection string, q interface{}, field string, value interface{}) (*mgo.ChangeInfo, error) {
	col, shutdown := PrepareCollection(session, database, collection)
	defer func() {
		close(shutdown)
	}()

	info, err := col.UpdateAll(q, bson.M{"$set": bson.M{field: value}})

	return info, err
}

// Update modify existing documents in a collection.
func Update(session *mgo.Session, database string, collection string, query interface{}, i interface{}) error {
	col, shutdown := PrepareCollection(session, database, collection)
	defer func() {
		close(shutdown)
	}()

	return col.Update(query, i)
}

// Upsert creates a new document and inserts it if no documents match the specified filter.
// If there are matching documents, then the operation modifies or replaces the matching document or documents.
func Upsert(session *mgo.Session, database string, collection string, query interface{}, i interface{}) (*mgo.ChangeInfo, error) {
	col, shutdown := PrepareCollection(session, database, collection)
	defer func() {
		close(shutdown)
	}()

	info, err := col.Upsert(query, i)

	return info, err
}

// Delete remove documents from a collection.
func Delete(session *mgo.Session, database string, collection string, query interface{}) error {
	col, shutdown := PrepareCollection(session, database, collection)
	defer func() {
		close(shutdown)
	}()

	return col.Remove(query)
}
