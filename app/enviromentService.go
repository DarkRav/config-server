package main

type Environment struct {
	Application string
	Profile     string
}

type EnvironmentService struct {
}

func (environment EnvironmentService) getConfigs(application string, profile string) (string, error) {
	return application + " " + profile, nil
}
