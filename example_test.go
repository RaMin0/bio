package bio

var b Bio

func ExampleBio_SetID() {
	if err := b.SetID("13-8994"); err != nil {
		// handle err...
	}
}

func ExampleBio_SetName() {
	if err := b.SetID("Ramy Aboul Naga"); err != nil {
		// handle err...
	}
}
func ExampleBio_SetDateOfBirth() {
	if err := b.SetDateOfBirth("21/03/1989"); err != nil {
		// handle err...
	}
}
func ExampleBio_SetMajor() {
	if err := b.SetMajor("DMET"); err != nil {
		// handle err...
	}
}
