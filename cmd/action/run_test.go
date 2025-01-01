package action

import (
	"fmt"
	"k8scli/cmd/pod"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		flags      map[string]string
		wantErr    bool
		errMessage string
	}{
		{
			name:       "pod name missing",
			args:       []string{""},
			flags:      nil,
			wantErr:    true,
			errMessage: "name or generateName is required",
		},
		{
			name:       "image flag missing",
			args:       []string{"test-pod"},
			flags:      nil,
			wantErr:    true,
			errMessage: "spec.containers[0].image: Required value",
		},
		{
			name: "valid pod creation",
			args: []string{"test-pod"},
			flags: map[string]string{
				"image": "nginx",
			},
			wantErr: false,
		},
		// invalid image provided
		// invalid port provided
		// valid port - valid pod
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// setup
			cmd := &cobra.Command{}
			cmd.Flags().String("image", "", "container image")
			cmd.Flags().String("port", "", "container port")

			// set flags
			if tc.flags != nil {
				for flag, value := range tc.flags {
					cmd.Flags().Set(flag, value)
				}
			}

			err := run(cmd, tc.args)

			if tc.wantErr {
				assert.Error(t, err)
				if tc.errMessage != "" {
					assert.Contains(t, err.Error(), tc.errMessage)
				}
			} else {
				assert.NoError(t, err)
			}
		})

		// cleanup
		cleanup()
	}
}

func cleanup() {
	cmd := &cobra.Command{}
	err := pod.Delete(cmd, []string{"test-pod"})
	if err != nil {
		fmt.Printf("Warn: failed to cleanup: %v\n", err)
	}
}
