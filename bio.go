package bio

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

var (
	MajorRegexp = regexp.MustCompile(`(?i)^CSEN|DMET$`)
)

type Bio struct {
	id          string
	name        string
	dateOfBirth time.Time
	major       string
}

func (bio *Bio) SetID(id string) error {
	if !strings.Contains(id, "-") {
		return fmt.Errorf("Failed in parsing ID \"%s\" as \"13-8994\". ", id)
	}

	bio.id = id

	return nil
}

func (bio *Bio) ID() string {
	return bio.id
}

func (bio *Bio) SetName(name string) {
	bio.name = name
}

func (bio *Bio) Name() string {
	return bio.name
}

func (bio *Bio) SetDateOfBirth(dateOfBirth string) (err error) {
	if bio.dateOfBirth, err = time.Parse("02/01/2006", dateOfBirth); err != nil {
		return fmt.Errorf("Failed in %s. ", strings.Split(err.Error(), ":")[0])
	}

	return
}

func (bio *Bio) DateOfBirth() time.Time {
	return bio.dateOfBirth
}

func (bio *Bio) SetMajor(major string) error {
	if !MajorRegexp.MatchString(major) {
		return fmt.Errorf("Failed in parsing major \"%s\" as \"CSEN\" or \"DMET\". ", major)
	}

	bio.major = strings.ToUpper(major)

	return nil
}

func (bio *Bio) Major() string {
	return bio.major
}

func (bio *Bio) Age() int {
	return time.Now().Year() - bio.dateOfBirth.Year()
}

func (bio *Bio) FirstName() string {
	return strings.Fields(bio.name)[0]
}

func (bio *Bio) FormattedDateOfBith() string {
	return bio.dateOfBirth.Format("Monday, January 02, 2006")
}
