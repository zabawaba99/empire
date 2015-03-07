package empire

type mockReleasesService struct {
	ReleasesService // Just to satisfy the interface.

	ReleasesCreateFunc func(*App, *Config, *Slug, string) (*Release, error)
}

func (s *mockReleasesService) ReleasesCreate(app *App, config *Config, slug *Slug, desc string) (*Release, error) {
	if s.ReleasesCreateFunc != nil {
		return s.ReleasesCreateFunc(app, config, slug, desc)
	}

	return nil, nil
}