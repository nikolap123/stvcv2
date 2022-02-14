package commandFormatter

func FormatQuery() *Command {

	c3 := Command{
		Command: "test123",
		Args:    []string{"nikola", "car"},
	}

	c2 := Command{
		Command: "test222",
		Args:    []string{"nikola2131", "car123"},
		Next:    &c3,
	}

	c1 := Command{
		Command: "test888",
		Args:    []string{"nikola754", "cargsfa"},
		Next:    &c2,
	}

	return &c1
}
