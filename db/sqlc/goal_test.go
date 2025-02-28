package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/mbaxamb33/pantopia/util"
	"github.com/stretchr/testify/require"
)

func createRandomGoal(t *testing.T) Goals {
	user := createRandomUser(t) // Ensure a user exists first

	arg := CreateGoalParams{
		UserID:      user.ID,
		Name:        util.RandomGoalName(),
		Description: sql.NullString{String: util.RandomDescription(), Valid: true},
		Type:        sql.NullString{String: "standard", Valid: true},
		TargetValue: sql.NullString{String: util.RandomTargetValue(), Valid: true},
	}

	goal, err := testQueries.CreateGoal(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, goal)

	require.Equal(t, arg.UserID, goal.UserID)
	require.Equal(t, arg.Name, goal.Name)
	require.Equal(t, arg.Description.String, goal.Description.String)
	require.Equal(t, arg.Type.String, goal.Type.String)
	require.Equal(t, arg.TargetValue.String, goal.TargetValue.String)

	require.NotZero(t, goal.ID)
	return goal
}

func TestCreateGoal(t *testing.T) {
	createRandomGoal(t)
}

func TestGetGoal(t *testing.T) {
	goal1 := createRandomGoal(t)
	goal2, err := testQueries.GetGoal(context.Background(), goal1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, goal2)

	require.Equal(t, goal1.ID, goal2.ID)
	require.Equal(t, goal1.UserID, goal2.UserID)
	require.Equal(t, goal1.Name, goal2.Name)
	require.Equal(t, goal1.Description.String, goal2.Description.String)
	require.Equal(t, goal1.Type.String, goal2.Type.String)
	require.Equal(t, goal1.TargetValue.String, goal2.TargetValue.String)
}

func TestUpdateGoal(t *testing.T) {
	goal1 := createRandomGoal(t)

	arg := UpdateGoalParams{
		ID:          goal1.ID,
		Name:        util.RandomGoalName(),
		Description: sql.NullString{String: util.RandomDescription(), Valid: true},
		Type:        sql.NullString{String: "custom", Valid: true},
		TargetValue: sql.NullString{String: util.RandomTargetValue(), Valid: true},
	}

	goal2, err := testQueries.UpdateGoal(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, goal2)

	require.Equal(t, goal1.ID, goal2.ID)
	require.Equal(t, arg.Name, goal2.Name)
	require.Equal(t, arg.Description.String, goal2.Description.String)
	require.Equal(t, arg.Type.String, goal2.Type.String)
	require.Equal(t, arg.TargetValue.String, goal2.TargetValue.String)
}

func TestDeleteGoal(t *testing.T) {
	goal := createRandomGoal(t)
	err := testQueries.DeleteGoal(context.Background(), goal.ID)
	require.NoError(t, err)

	_, err = testQueries.GetGoal(context.Background(), goal.ID)
	require.Error(t, err)
}

func TestListGoals(t *testing.T) {
	user := createRandomUser(t) // Ensure all goals belong to the same user

	for i := 0; i < 5; i++ {
		arg := CreateGoalParams{
			UserID:      user.ID,
			Name:        util.RandomGoalName(),
			Description: sql.NullString{String: util.RandomDescription(), Valid: true},
			Type:        sql.NullString{String: "standard", Valid: true},
			TargetValue: sql.NullString{String: util.RandomTargetValue(), Valid: true},
		}
		_, err := testQueries.CreateGoal(context.Background(), arg)
		require.NoError(t, err)
	}

	arg := ListGoalsParams{
		UserID: user.ID, // Query goals for the specific user
		Limit:  5,
		Offset: 0,
	}

	goals, err := testQueries.ListGoals(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, goals, 5) // Should return 5 goals

	for _, goal := range goals {
		require.NotEmpty(t, goal)
		require.Equal(t, user.ID, goal.UserID) // Ensure all goals belong to the same user
	}
}
