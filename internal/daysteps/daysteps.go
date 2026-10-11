package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps                 int           // кол-во шагов
	Duration              time.Duration // Длительность прогулки
	personaldata.Personal               // встроенная структура Personal из пакета personaldata, у которой есть метод Print()
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("неверное количество параметров")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть положительным")
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("длительность тренировки должна быть положительной")
	}
	ds.Duration = duration
	ds.Steps = steps
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {

	if ds.Steps <= 0 {
		return "", errors.New("количество шагов должно быть положительным")
	}
	if ds.Duration <= 0 {
		return "", errors.New("длительность тренировки должна быть положительной")
	}
	if ds.Weight <= 0 || ds.Height <= 0 {
		return "", errors.New("данные должны быть положительными")
	}
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(
		ds.Steps, ds.Weight, ds.Height, ds.Duration,
	)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories), nil

}
