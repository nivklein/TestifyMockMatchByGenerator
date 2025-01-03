package TestifyMockMatchByGenerator

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockMyService struct {
	mock.Mock
}

type InputStruct struct {
	ANumber int
	AString string
}

func (m *MockMyService) Process(data interface{}) string {
	args := m.Called(data)
	return args.String(0)
}

func TestCreateMatcher(t *testing.T) {
	m := new(MockMyService)

	input := InputStruct{
		ANumber: 42,
		AString: "Hello, World!",
	}

	m.On("Process", mock.MatchedBy(CreateMatcher(input))).Return("Success")

	anotherInput := InputStruct{
		ANumber: 42,
		AString: "Hello, World!",
	}

	res := m.Process(anotherInput)
	require.Equal(t, "Success", res)

	// now let try assert called with the other one
	m.AssertCalled(t, "Process", mock.MatchedBy(CreateMatcher(input)))

}
