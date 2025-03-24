package seed

import (
	"awesomeProject/api/workout-api/types"
	"time"
)

//TODO: this file should actually connect to the database in the future, but using seeded data for now

func SeedWorkoutData() []types.WorkoutComponent {
	// Sample titles for the workouts
	title1 := "Morning Strength Routine"
	title2 := "Evening Cardio Blast"
	title3 := "Full Body HIIT"
	title4 := "Leg Day Routine"
	title5 := "Upper Body Strength"

	// Create metadata
	metadata1 := types.NewItemMetadata("001", time.Now(), &title1)
	metadata2 := types.NewItemMetadata("002", time.Now(), &title2)
	metadata3 := types.NewItemMetadata("003", time.Now(), &title3)
	metadata4 := types.NewItemMetadata("004", time.Now(), &title4)
	metadata5 := types.NewItemMetadata("005", time.Now(), &title5)

	// Define example RepComponents
	repComponent1 := types.RepComponent{
		ItemMetadata: *metadata1,
		Reps:         10,
		Sets:         3,
		RestDuration: "1 minute",
	}

	repComponent2 := types.RepComponent{
		ItemMetadata: *metadata2,
		Reps:         15,
		Sets:         5,
		RestDuration: "1 minute",
	}

	// Define example NonRepComponents
	nonRepComponent1 := types.NonRepComponent{Duration: 30} // 30-minute cardio

	nonRepComponent2 := types.NonRepComponent{Duration: 45} // 45-minute cardio

	// Create a slice of WorkoutComponents
	workoutComponents := []types.WorkoutComponent{
		{
			ItemMetadata:  *metadata1,
			Part:          "Full Body",
			Tags:          []string{"strength", "morning", "routine"},
			Category:      "Strength Training",
			Type:          repComponent1,
			Name:          &title1,
			CaloriesBurnt: new(uint32),
			Notes:         []string{"Focus on form and controlled movements."},
		},
		{
			ItemMetadata:  *metadata2,
			Part:          "Cardio",
			Tags:          []string{"cardio", "evening", "routine"},
			Category:      "Cardio Training",
			Type:          repComponent2,
			Name:          &title2,
			CaloriesBurnt: new(uint32),
			Notes:         []string{"Increase intensity in the second half."},
		},
		{
			ItemMetadata:  *metadata3,
			Part:          "Full Body",
			Tags:          []string{"HIIT", "full body", "intense"},
			Category:      "HIIT",
			Type:          nonRepComponent1, // Non-rep component
			Name:          &title3,
			CaloriesBurnt: new(uint32),
			Notes:         []string{"Short intervals with high intensity."},
		},
		{
			ItemMetadata:  *metadata4,
			Part:          "Legs",
			Tags:          []string{"strength", "legs", "power"},
			Category:      "Strength Training",
			Type:          repComponent1,
			Name:          &title4,
			CaloriesBurnt: new(uint32),
			Notes:         []string{"Heavy squats and lunges."},
		},
		{
			ItemMetadata:  *metadata5,
			Part:          "Upper Body",
			Tags:          []string{"strength", "upper body", "routine"},
			Category:      "Strength Training",
			Type:          nonRepComponent2, // Non-rep component
			Name:          &title5,
			CaloriesBurnt: new(uint32),
			Notes:         []string{"Push-ups, pull-ups, and dips."},
		},
	}

	return workoutComponents
}
