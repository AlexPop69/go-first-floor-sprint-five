package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

func TestFitnessSuite(t *testing.T) {
	suite.Run(t, new(FitnessSuite))
}

type FitnessSuite struct {
	suite.Suite
}

func (s *FitnessSuite) TestTrainingInfo() {
	s.Run("running", func() {
		running := Running{
			Training: Training{
				TrainingType: "Бег",
				Action:       5000,
				LenStep:      LenStep,
				Duration:     30 * time.Minute,
				Weight:       85,
			},
		}

		res := ReadData(running)

		distance := running.distance()
		speed := running.meanSpeed()
		calories := running.Calories()
		expected := fmt.Sprintf("Тип тренировки: %s\nДлительность: %v мин\nДистанция: %.2f км.\nСр. скорость: %.2f км/ч\nПотрачено ккал: %.2f\n", running.TrainingType, running.Duration.Minutes(), distance, speed, calories)

		s.Assert().Equal(expected, res, "Результат выполнения функции ShowTrainingInfo не совпадает с ожидаемым")
	})

	s.Run("walking", func() {
		walking := Walking{
			Training: Training{
				TrainingType: "Ходьба",
				Action:       20000,
				LenStep:      LenStep,
				Duration:     3*time.Hour + 45*time.Minute,
				Weight:       85,
			},
			Height: 185,
		}

		res := ReadData(walking)

		distance := walking.distance()
		speed := walking.meanSpeed()
		calories := walking.Calories()

		expected := fmt.Sprintf("Тип тренировки: %s\nДлительность: %v мин\nДистанция: %.2f км.\nСр. скорость: %.2f км/ч\nПотрачено ккал: %.2f\n", walking.TrainingType, walking.Duration.Minutes(), distance, speed, calories)

		s.Assert().Equal(expected, res, "Результат выполнения функции ShowTrainingInfo не совпадает с ожидаемым")
	})

	s.Run("swimming", func() {
		swimming := Swimming{
			Training: Training{
				TrainingType: "Плавание",
				Action:       2000,
				LenStep:      SwimmingLenStep,
				Duration:     90 * time.Minute,
				Weight:       85,
			},
			LengthPool: 50,
			CountPool:  5,
		}
		res := ReadData(swimming)

		distance := swimming.distance()
		speed := swimming.meanSpeed()
		calories := swimming.Calories()
		expected := fmt.Sprintf("Тип тренировки: %s\nДлительность: %v мин\nДистанция: %.2f км.\nСр. скорость: %.2f км/ч\nПотрачено ккал: %.2f\n", swimming.TrainingType, swimming.Duration.Minutes(), distance, speed, calories)

		s.Assert().Equal(expected, res, "Результат выполнения функции ShowTrainingInfo не совпадает с ожидаемым")
	})
}
