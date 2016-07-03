package identity

import (
	"testing"
)

func setup() {
	MigrateSchema()
	ResetData()
}

func TestCreateUniqueUser(t *testing.T) {
	setup()

	t.Log("User creation succeeds with unique email")

	user := CreateUser("Test User", "test@example.com", "password")

	if user == nil {
		t.Errorf("Expected user to be created.")
	}
}

func TestCreateNonUniqueUser(t *testing.T) {
	setup()

	t.Log("User creation fails with non-unique email")

	user0 := CreateUser("Test User", "test@example.com", "password")
	if user0 == nil {
		t.Errorf("Expected user to be created.")
	}

	user1 := CreateUser("Test User", "test@example.com", "password")
	if user1 != nil {
		t.Errorf("Expected user not to be created.")
	}
}

func TestAuthenticateUserSuccess(t *testing.T) {
	setup()

	t.Log("User authentication succeeds with valid credentials")

	if CreateUser("Test User", "test@example.com", "password") == nil {
		t.Errorf("Expected user to be created.")
	}

	user := AuthenticateUser("test@example.com", "password")
	if user == nil {
		t.Errorf("Expected user to be authenticated.")
	}
}

func TestAuthenticateUserFailure(t *testing.T) {
	setup()

	t.Log("User authentication fails with invalid credentials")

	if CreateUser("Test User", "test@example.com", "password") == nil {
		t.Errorf("Expected user to be created.")
	}

	user := AuthenticateUser("test@example.com", "wrongpassword")
	if user != nil {
		t.Errorf("Expected user to not be authenticated.")
	}
}

func TestCreateToken(t *testing.T) {
	setup()

	t.Log("Token creation succeeds for valid user")

	user := CreateUser("Test User", "test@example.com", "password")
	if user == nil {
		t.Errorf("Expected user to be created.")
	}

	token := CreateToken(user)
	if token == nil {
		t.Errorf("Expected token to be created.")
	}
}

func TestDeleteTokenSuccess(t *testing.T) {
	setup()

	t.Log("Token deletion succeeds with valid token id")

	user := CreateUser("Test User", "test@example.com", "password")
	if user == nil {
		t.Errorf("Expected user to be created.")
	}

	token := CreateToken(user)
	if token == nil {
		t.Errorf("Expected token to be created.")
	}

	deletedToken := DeleteToken(token.ID)
	if !deletedToken {
		t.Errorf("Expected token to be deleted.")
	}
}

func TestDeleteTokenFailure(t *testing.T) {
	setup()

	t.Log("Token deletion fails with invalid token id")

	user := CreateUser("Test User", "test@example.com", "password")
	if user == nil {
		t.Errorf("Expected user to be created.")
	}

	deletedToken := DeleteToken(1)
	if deletedToken {
		t.Errorf("Expected no token to be deleted.")
	}
}

func TestUpdateUserSuccess(t *testing.T) {
	setup()

	t.Log("User update succeeds")

	user := CreateUser("Test User", "test@example.com", "password")
	if user == nil {
		t.Errorf("Expected user to be created.")
	}

	user.Name = "New Name"

	if !user.Save() {
		t.Errorf("Expected user to be saved")
	}
}
