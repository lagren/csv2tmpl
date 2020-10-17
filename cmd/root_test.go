package cmd

func ExampleRun() {
	rootCmd.SetArgs([]string{
		"-i", "testdata/basic.csv",
		"-t", "{{ .col0 }} -> {{ .col1 }}",
	})

	rootCmd.Execute()

	// Output:
	// foo -> bar
	// baz -> boz
}
