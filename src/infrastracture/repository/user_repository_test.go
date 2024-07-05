package repository

import (
	"database/sql"
	"testing"

	"github.com/KakinokiKanta/Mybrary-backend/domain"
	"github.com/google/go-cmp/cmp"
)

func TestCreateUserRepository(t *testing.T) {
	userDomain, _ := domain.NewUser("Test user")

	tests := []struct {
		testName string
		db *sql.DB
		args domain.User
		expected domain.User
		expectedErr bool
	}{
		{
			testName: "Successfully create user",
			db: testDB,
			args: *userDomain,
			expected: *userDomain,
			expectedErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			repo := NewUserRepository(testDB)

			result, err := repo.Create(tt.args)
			if (err != nil) != tt.expectedErr {
				t.Errorf("[TestCase '%s'] Result: '%v' | ExpectedError: '%v'", tt.testName, err, tt.expectedErr)
				return
			}
			diff := cmp.Diff(
				result, tt.expected,
				cmp.AllowUnexported(domain.User{}),
			)
			if diff != "" {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'", tt.testName, result, tt.expected)
			}
		})
	}
}