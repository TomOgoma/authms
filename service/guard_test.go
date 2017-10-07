package service_test

import (
	"testing"

	"github.com/tomogoma/authms/service"
	testingH "github.com/tomogoma/authms/testing"
	"github.com/tomogoma/go-commons/errors"
)

func TestNewGuard(t *testing.T) {
	tt := []struct {
		name      string
		db        service.APIKeyStore
		kg        service.KeyGenerator
		masterKey string
		expErr    bool
	}{
		{
			name:      "all deps provided",
			db:        &testingH.DBMock{},
			kg:        &testingH.GeneratorMock{},
			masterKey: "a-master-key",
			expErr:    false,
		},
		{
			name:      "implicit deps",
			db:        &testingH.DBMock{},
			kg:        nil,
			masterKey: "",
			expErr:    false,
		},
		{
			name:      "nil db",
			db:        nil,
			kg:        &testingH.GeneratorMock{},
			masterKey: "a-master-key",
			expErr:    true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			g, err := service.NewGuard(
				tc.db,
				service.WithKeyGenerator(tc.kg),
				service.WithMasterKey(tc.masterKey),
			)
			if tc.expErr {
				if err == nil {
					t.Fatalf("Expected error got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("service.NewGuard(): %v", err)
			}
			if g == nil {
				t.Fatalf("service.NewGuard returned nil *service.Guard")
			}
		})
	}
}

func TestGuard_NewAPIKey(t *testing.T) {
	validKey := "some-api-key"
	tt := []struct {
		name     string
		userID   string
		kg       *testingH.GeneratorMock
		db       *testingH.DBMock
		expErr   bool
		expClErr bool
	}{
		{
			name:   "valid",
			userID: "12345",
			kg:     &testingH.GeneratorMock{ExpSRBs: []byte(validKey)},
			db:     &testingH.DBMock{ExpInsAPIK: &service.APIKey{UserID: "12345", ID: "apiid", APIKey: validKey}},
			expErr: false,
		},
		{
			name:     "missing userID",
			userID:   "",
			kg:       &testingH.GeneratorMock{ExpSRBs: []byte(validKey)},
			db:       &testingH.DBMock{ExpInsAPIK: &service.APIKey{UserID: "12345", ID: "apiid", APIKey: validKey}},
			expErr:   true,
			expClErr: true,
		},
		{
			name:   "key gen report error",
			userID: "12345",
			kg:     &testingH.GeneratorMock{ExpSRBsErr: errors.Newf("an error")},
			db:     &testingH.DBMock{ExpInsAPIK: &service.APIKey{UserID: "12345", ID: "apiid", APIKey: validKey}},
			expErr: true,
		},
		{
			name:   "db report error",
			userID: "12345",
			kg:     &testingH.GeneratorMock{ExpSRBs: []byte(validKey)},
			db:     &testingH.DBMock{ExpInsAPIKErr: errors.Newf("whoops, an error")},
			expErr: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			g := newGuard(t, "", tc.kg, tc.db)
			ak, err := g.NewAPIKey(tc.userID)
			if tc.expErr {
				if err == nil {
					t.Fatalf("Error: %v", err)
				}
				if tc.expClErr != g.IsClientError(err) {
					t.Errorf("Expect service.Guard#IsClientError %t, got %t",
						tc.expClErr, g.IsClientError(err))
				}
				return
			}
			if ak == nil {
				t.Fatalf("yielded nil *service.APIKey")
			}
			if ak.APIKey != validKey {
				t.Errorf("API Key mismatch: expect '%s' got '%s'",
					validKey, ak.APIKey)
			}
		})
	}
}

func TestGuard_APIKeyValid(t *testing.T) {
	validKey := "some-api-key"
	tt := []struct {
		name            string
		key             string
		expUsrID        string
		db              *testingH.DBMock
		masterKey       string
		expErr          bool
		expForbidden    bool
		expUnauthorized bool
	}{
		{
			name:     "valid (db)",
			expUsrID: "12345",
			key:      "12345." + validKey,
			db: &testingH.DBMock{ExpAPIKsBUsrID: []service.APIKey{
				{APIKey: "first-api-key"},
				{APIKey: validKey},
				{APIKey: "last-api-key"},
			}},
			expErr: false,
		},
		{
			name:      "valid (master)",
			key:       "the-master-key",
			masterKey: "the-master-key",
			expUsrID:  "master",
			db: &testingH.DBMock{ExpAPIKsBUsrID: []service.APIKey{
				{APIKey: "first-api-key"},
				{APIKey: validKey},
				{APIKey: "last-api-key"},
			}},
			expErr: false,
		},
		{
			name:     "empty",
			key:      "",
			expUsrID: "",
			db: &testingH.DBMock{ExpAPIKsBUsrID: []service.APIKey{
				{APIKey: "first-api-key"},
				{APIKey: validKey},
				{APIKey: "last-api-key"},
			}},
			expErr:          true,
			expForbidden:    false,
			expUnauthorized: true,
		},
		{
			name:     "separator only",
			key:      ".",
			expUsrID: "",
			db: &testingH.DBMock{ExpAPIKsBUsrID: []service.APIKey{
				{APIKey: "first-api-key"},
				{APIKey: validKey},
				{APIKey: "last-api-key"},
			}},
			expErr:          true,
			expForbidden:    false,
			expUnauthorized: true,
		},
		{
			name:     "missing separator + userID",
			key:      "" + validKey,
			expUsrID: "",
			db: &testingH.DBMock{ExpAPIKsBUsrID: []service.APIKey{
				{APIKey: "first-api-key"},
				{APIKey: validKey},
				{APIKey: "last-api-key"},
			}},
			expErr:          true,
			expForbidden:    false,
			expUnauthorized: true,
		},
		{
			name:     "missing separator + key",
			key:      "12345" + "",
			expUsrID: "",
			db: &testingH.DBMock{ExpAPIKsBUsrID: []service.APIKey{
				{APIKey: "first-api-key"},
				{APIKey: validKey},
				{APIKey: "last-api-key"},
			}},
			expErr:          true,
			expForbidden:    false,
			expUnauthorized: true,
		},
		{
			name:     "invalid key",
			key:      "12345." + "some-invalid-key",
			expUsrID: "12345",
			db: &testingH.DBMock{ExpAPIKsBUsrID: []service.APIKey{
				{APIKey: "first-api-key"},
				{APIKey: validKey},
				{APIKey: "last-api-key"},
			}},
			expErr:          true,
			expForbidden:    true,
			expUnauthorized: false,
		},
		{
			name:            "none found",
			key:             "12345." + validKey,
			expUsrID:        "12345",
			db:              &testingH.DBMock{ExpAPIKsBUsrIDErr: errors.NewNotFound("no keys for 12345")},
			expErr:          true,
			expForbidden:    true,
			expUnauthorized: false,
		},
		{
			name:     "db report error",
			key:      "12345." + "some-invalid-key",
			expUsrID: "12345",
			db:       &testingH.DBMock{ExpAPIKsBUsrIDErr: errors.New("some errors")},
			expErr:   true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			g := newGuard(t, tc.masterKey, &testingH.GeneratorMock{}, tc.db)
			usrID, err := g.APIKeyValid(tc.key)
			if usrID != tc.expUsrID {
				t.Errorf("Expected userID '%s', got '%s'", tc.expUsrID, usrID)
			}
			if tc.expErr {
				if err == nil {
					t.Fatalf("Expected an error, got nil")
				}
				if tc.expUnauthorized != g.IsUnauthorizedError(err) {
					t.Errorf("Expect service.Guard#IsUnauthorizedError %t, got %t",
						tc.expUnauthorized, g.IsUnauthorizedError(err))
				}
				if tc.expForbidden != g.IsForbiddenError(err) {
					t.Errorf("Expect service.Guard#IsForbiddenError %t, got %t",
						tc.expForbidden, g.IsForbiddenError(err))
				}
				return
			}
			if err != nil {
				t.Fatalf("Expected nil error, got %v", err)
			}
		})
	}
}

func newGuard(t *testing.T, master string, kg service.KeyGenerator, db service.APIKeyStore) *service.Guard {
	g, err := service.NewGuard(
		db,
		service.WithKeyGenerator(kg),
		service.WithMasterKey(master),
	)
	if err != nil {
		t.Fatalf("service.NewGuard(): %v", err)
	}
	return g
}
