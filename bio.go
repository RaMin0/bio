package bio

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

var (
	majorRegexp = regexp.MustCompile(`(?i)^CSEN|DMET$`)
)

// Bio holds the data for a bio
type Bio struct {
	id          string
	name        string
	dateOfBirth time.Time
	major       string
}

// SetID sets the ID and returns an error if it is not valid
func (bio *Bio) SetID(id string) error {
	if !strings.Contains(id, "-") {
		return fmt.Errorf("ID should look like 13-8994: %v", id)
	}

	bio.id = id

	return nil
}

// ID returns the ID
func (bio Bio) ID() string {
	return bio.id
}

// SetName sets the name and returns an error if it is empty
func (bio *Bio) SetName(name string) error {
	if name == "" {
		return fmt.Errorf("name is empty")
	}

	bio.name = name

	return nil
}

// Name returns the name
func (bio Bio) Name() string {
	return bio.name
}

// SetDateOfBirth sets the date of birth and returns an error if it is not valid
func (bio *Bio) SetDateOfBirth(dateOfBirth string) (err error) {
	if bio.dateOfBirth, err = time.Parse("02/01/2006", dateOfBirth); err != nil {
		err = fmt.Errorf("%s: %v", strings.SplitN(err.Error(), ": ", 2)[1], dateOfBirth)
	}

	return
}

// DateOfBirth returns the date of birth
func (bio Bio) DateOfBirth() time.Time {
	return bio.dateOfBirth
}

// SetMajor sets the major and returns an error if it is not valid
func (bio *Bio) SetMajor(major string) error {
	if !majorRegexp.MatchString(major) {
		return fmt.Errorf("major should be either CSEN or DMET: %v", major)
	}

	bio.major = strings.ToUpper(major)

	return nil
}

// Major returns the major
func (bio Bio) Major() string {
	return bio.major
}

// Age returns the age
func (bio Bio) Age() int {
	return time.Now().Year() - bio.dateOfBirth.Year()
}

// FirstName returns the first name
func (bio Bio) FirstName() string {
	return strings.Fields(bio.name)[0]
}

// FormattedDateOfBith returns the formatted date of birth
func (bio Bio) FormattedDateOfBith() string {
	return bio.dateOfBirth.Format("Monday, January 02, 2006")
}
