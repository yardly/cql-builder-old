package cql_test

import (
	"fmt"
	"testing"

	"github.com/samuelngs/cql-builder"
)

func TestNew(t *testing.T) {
	// query, args := cql.
	// 	Connect(nil).
	// 	Prepare(
	// 		cql.
	// 			Table("profile_by_user_id").
	// 			Select(
	// 				cql.Field("profile_id"),
	// 			).
	// 			Where(
	// 				cql.Eq("profile_id", "1234563c-6282-11e7-907b-a6006ad3dba0"),
	// 				cql.If(true, cql.Eq("created_at", time.Now())),
	// 			).
	// 			Limit(1),
	// 	).
	// 	Compile()
	// fmt.Println(query, args)
	query, args := cql.
		Connect(nil).
		Prepare(
			cql.
				Table("profile_by_user_id").
				Update(
					cql.Field("counter1", cql.Increment(1)),
					cql.Field("counter2", cql.Decrement(1)),
				).
				Where(
					cql.Eq("id", "test"),
				),
		).
		Compile()
	fmt.Println(query, args)
	// query, args := cql.
	// 	Connect(nil).
	// 	Prepare(
	// 		cql.
	// 			Table("profile_by_user_id").
	// 			Delete(
	// 				cql.Field("profile_id1"),
	// 				cql.Field("profile_id2"),
	// 				cql.Field("profile_id3"),
	// 				cql.Field("profile_id4"),
	// 			).
	// 			Where(
	// 				cql.Lucene(
	// 					lucene.Column("haha"),
	// 					lucene.Filter(
	// 						lucene.BooleanShould(
	// 							lucene.Match("field_name", "field_value"),
	// 							lucene.GeoShapePoint("field_name", 10, 20),
	// 						),
	// 					),
	// 					lucene.Query(),
	// 					lucene.Sort(),
	// 					lucene.Refresh(),
	// 				),
	// 			),
	// 	).
	// 	Compile()
	// fmt.Println(query, args)
	// query, _ := cql.
	// 	Connect(nil).
	// 	Batch(
	// 		cql.
	// 			Table("profile_by_user_id").
	// 			Insert(
	// 				cql.Field("profile_id", "0467303c-6282-11e7-907b-a6006ad3dba0"),
	// 				cql.Field("profile_channel", "invitation"),
	// 			),
	// 		cql.
	// 			Table("profile_by_profile_id").
	// 			Insert(
	// 				cql.Field("profile_id", "0467303c-6282-11e7-907b-a6006ad3dba0"),
	// 				cql.Field("profile_channel", "invitation"),
	// 			).
	// 			IfNotExist(),
	// 	).
	// 	Compile()
	// fmt.Println(query)
	// query, args := cql.
	// 	Connect(nil).
	// 	Batch(
	// 		cql.
	// 			Table("profile_by_user_id").
	// 			Update(
	// 				cql.Field("profile_channel", "invitation"),
	// 				cql.Field("profile_user_id", "0467303c-6282-11e7-907b-a6006ad3dba0"),
	// 			).
	// 			Where(
	// 				cql.Eq("profile_id", "1234563c-6282-11e7-907b-a6006ad3dba0"),
	// 				cql.If(true, cql.Eq("created_at", time.Now())),
	// 			).
	// 			IfNotExist(),
	// 	).
	// 	Compile()
	// fmt.Println(query, args)
}
