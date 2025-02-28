package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/mbaxamb33/pantopia/util"
	"github.com/stretchr/testify/require"
)

func createRandomContact(t *testing.T) Contacts {
	user := createRandomUser(t) // Ensure user exists first

	arg := CreateContactParams{
		UserID:      user.ID,
		FirstName:   sql.NullString{String: util.RandomName(), Valid: true},
		LastName:    sql.NullString{String: util.RandomName(), Valid: true},
		Email:       sql.NullString{String: util.RandomEmail(), Valid: true},
		Phone:       sql.NullString{String: util.RandomPhoneNumber(), Valid: true},
		CompanyName: sql.NullString{String: util.RandomCompanyName(), Valid: true},
		Address:     sql.NullString{String: util.RandomAddress(), Valid: true},
	}

	contact, err := testQueries.CreateContact(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, contact)

	require.Equal(t, arg.UserID, contact.UserID)
	require.Equal(t, arg.FirstName.String, contact.FirstName.String)
	require.Equal(t, arg.LastName.String, contact.LastName.String)
	require.Equal(t, arg.Email.String, contact.Email.String)
	require.Equal(t, arg.Phone.String, contact.Phone.String)
	require.Equal(t, arg.CompanyName.String, contact.CompanyName.String)
	require.Equal(t, arg.Address.String, contact.Address.String)

	require.NotZero(t, contact.ID)
	return contact
}

func TestCreateContact(t *testing.T) {
	createRandomContact(t)
}

func TestGetContact(t *testing.T) {
	contact1 := createRandomContact(t)
	contact2, err := testQueries.GetContact(context.Background(), contact1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, contact2)

	require.Equal(t, contact1.ID, contact2.ID)
	require.Equal(t, contact1.UserID, contact2.UserID)
	require.Equal(t, contact1.FirstName.String, contact2.FirstName.String)
	require.Equal(t, contact1.LastName.String, contact2.LastName.String)
	require.Equal(t, contact1.Email.String, contact2.Email.String)
	require.Equal(t, contact1.Phone.String, contact2.Phone.String)
	require.Equal(t, contact1.CompanyName.String, contact2.CompanyName.String)
	require.Equal(t, contact1.Address.String, contact2.Address.String)
}

func TestUpdateContact(t *testing.T) {
	contact1 := createRandomContact(t)

	arg := UpdateContactParams{
		ID:          contact1.ID,
		FirstName:   sql.NullString{String: util.RandomName(), Valid: true},
		LastName:    sql.NullString{String: util.RandomName(), Valid: true},
		Email:       sql.NullString{String: util.RandomEmail(), Valid: true},
		Phone:       sql.NullString{String: util.RandomPhoneNumber(), Valid: true},
		CompanyName: sql.NullString{String: util.RandomCompanyName(), Valid: true},
		Address:     sql.NullString{String: util.RandomAddress(), Valid: true},
	}

	contact2, err := testQueries.UpdateContact(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, contact2)

	require.Equal(t, contact1.ID, contact2.ID)
	require.Equal(t, arg.FirstName.String, contact2.FirstName.String)
	require.Equal(t, arg.LastName.String, contact2.LastName.String)
	require.Equal(t, arg.Email.String, contact2.Email.String)
	require.Equal(t, arg.Phone.String, contact2.Phone.String)
	require.Equal(t, arg.CompanyName.String, contact2.CompanyName.String)
	require.Equal(t, arg.Address.String, contact2.Address.String)
}

func TestDeleteContact(t *testing.T) {
	contact := createRandomContact(t)
	err := testQueries.DeleteContact(context.Background(), contact.ID)
	require.NoError(t, err)

	_, err = testQueries.GetContact(context.Background(), contact.ID)
	require.Error(t, err)
}

func TestListContacts(t *testing.T) {
	user := createRandomUser(t) // Create a user to associate all contacts with

	for i := 0; i < 5; i++ {
		arg := CreateContactParams{
			UserID:      user.ID,
			FirstName:   sql.NullString{String: util.RandomName(), Valid: true},
			LastName:    sql.NullString{String: util.RandomName(), Valid: true},
			Email:       sql.NullString{String: util.RandomEmail(), Valid: true},
			Phone:       sql.NullString{String: util.RandomPhoneNumber(), Valid: true},
			CompanyName: sql.NullString{String: util.RandomCompanyName(), Valid: true},
			Address:     sql.NullString{String: util.RandomAddress(), Valid: true},
		}
		_, err := testQueries.CreateContact(context.Background(), arg)
		require.NoError(t, err)
	}

	arg := ListContactsParams{
		UserID: user.ID, // Query contacts for the specific user
		Limit:  5,
		Offset: 0,
	}

	contacts, err := testQueries.ListContacts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, contacts, 5) // Should return 5 contacts

	for _, contact := range contacts {
		require.NotEmpty(t, contact)
		require.Equal(t, user.ID, contact.UserID) // Ensure they belong to the same user
	}
}
