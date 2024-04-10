package utils_test

type MockSentryService struct {
	CapturedErrors []error
}

func (m *MockSentryService) CaptureException(err error) {
	m.CapturedErrors = append(m.CapturedErrors, err)
}
