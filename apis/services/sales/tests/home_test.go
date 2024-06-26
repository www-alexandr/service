package tests

import (
	"runtime/debug"
	"testing"
)

func Test_Home(t *testing.T) {
	t.Parallel()

	// -------------------------------------------------------------------------

	dbTest, appTest := startTest(t, "Test_Home")
	defer func() {
		if r := recover(); r != nil {
			t.Log(r)
			t.Error(string(debug.Stack()))
		}
		dbTest.Teardown()
	}()

	// -------------------------------------------------------------------------

	sd, err := insertHomeSeed(dbTest)
	if err != nil {
		t.Fatalf("Seeding error: %s", err)
	}

	// -------------------------------------------------------------------------

	appTest.Run(t, homeQuery200(sd), "home-query-200")
	appTest.Run(t, homeQueryByID200(sd), "home-querybyid-200")

	appTest.Run(t, homeCreate200(sd), "home-create-200")
	appTest.Run(t, homeCreate401(sd), "home-create-401")
	appTest.Run(t, homeCreate400(sd), "home-create-400")

	appTest.Run(t, homeUpdate200(sd), "home-update-200")
	appTest.Run(t, homeUpdate401(sd), "home-update-401")
	appTest.Run(t, homeUpdate400(sd), "home-update-400")

	appTest.Run(t, homeDelete200(sd), "home-delete-200")
	appTest.Run(t, homeDelete401(sd), "home-delete-401")
}
