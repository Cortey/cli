package registry

// write tests for the Import function
// the tests should cover the happy path and the error path
// the tests should not actually push an image to a registry
// instead, they should mock the calls to the registry

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/stretchr/testify/require"
)

type mockAuthenticator struct {
	authn.Authenticator
}

func (ma *mockAuthenticator) Authorization() (*authn.AuthConfig, error) {
	return nil, errors.New("authorization error")
}

type mockPortforwardTransport struct {
}

func (mpt *mockPortforwardTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return nil, errors.New("portforward error")
}

func TestImport1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		imageName   string
		opts        ImportOptions
		expectedErr error
	}{
		{
			name:      "Happy path",
			imageName: "test:tag",
			opts: ImportOptions{
				ClusterAPIRestConfig: nil,
				RegistryAuth:         &mockAuthenticator{},
				RegistryPullHost:     "registry",
				RegistryPodName:      "pod",
				RegistryPodNamespace: "namespace",
				RegistryPodPort:      "port",
			},
			expectedErr: errors.New("failed to push image to the in-cluster registry: authorization error"),
		},
		{
			name:      "Error path",
			imageName: "test:tag",
			opts: ImportOptions{
				ClusterAPIRestConfig: nil,
				RegistryAuth:         &mockAuthenticator{},
				RegistryPullHost:     "registry",
				RegistryPodName:      "pod",
				RegistryPodNamespace: "namespace",
				RegistryPodPort:      "port",
			},
			expectedErr: errors.New("failed to push image to the in-cluster registry: authorization error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Import(context.Background(), tt.imageName, tt.opts)
			require.Equal(t, tt.expectedErr, err)
		})
	}
}

// Write a test for the function Import the test should cover the happy path and the error path the test should not actually push an image to a registry instead, it should mock the calls to the registry the test should also cover the case when the image name is not in the expected format 'image:tag'

func TestImport2(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		imageName   string
		opts        ImportOptions
		expectedErr error
	}{
		{
			name:      "Happy path",
			imageName: "test:tag",
			opts: ImportOptions{
				ClusterAPIRestConfig: nil,
				RegistryAuth:         &mockAuthenticator{},
				RegistryPullHost:     "registry",
				RegistryPodName:      "pod",
				RegistryPodNamespace: "namespace",
				RegistryPodPort:      "port",
			},
			expectedErr: errors.New("failed to push image to the in-cluster registry: authorization error"),
		},
		{
			name:      "Error path",
			imageName: "test:tag",
			opts: ImportOptions{
				ClusterAPIRestConfig: nil,
				RegistryAuth:         &mockAuthenticator{},
				RegistryPullHost:     "registry",
				RegistryPodName:      "pod",
				RegistryPodNamespace: "namespace",
				RegistryPodPort:      "port",
			},
			expectedErr: errors.New("failed to push image to the in-cluster registry: authorization error"),
		},
		{
			name:      "Invalid image name",
			imageName: "test",
			opts: ImportOptions{
				ClusterAPIRestConfig: nil,
				RegistryAuth:         &mockAuthenticator{},
				RegistryPullHost:     "registry",
				RegistryPodName:      "pod",
				RegistryPodNamespace: "namespace",
				RegistryPodPort:      "port",
			},
			expectedErr: errors.New("image 'test' not in expected format 'image:tag'"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := Import(context.Background(), test.imageName, test.opts)
			require.Equal(t, test.expectedErr, err)
		})
	}
}
