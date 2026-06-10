/*
Copyright 2026.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fbc

import (
	"github.com/konflux-ci/operator-foundry/pkg/lifecycle"
	"github.com/spf13/cobra"
)

func newInjectLifecycleCmd() *cobra.Command {
	var dockerfilePath string
	var buildContextPath string
	var lifecycleDir string
	var packages string

	cmd := &cobra.Command{
		Use:   "inject-lifecycle",
		Short: "Inject lifecycle.json into FBC catalog source directories",
		Long: `Injects pre-generated lifecycle.json files into the catalog source
directories for the given OLM packages.

Lifecycle injection is only performed for OCP versions >= 5.0.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return lifecycle.InjectLifecycle(dockerfilePath, buildContextPath, lifecycleDir, packages)
		},
	}

	cmd.Flags().StringVar(&dockerfilePath, "dockerfile", "", "Path to the FBC Dockerfile (required)")
	cmd.Flags().StringVar(&buildContextPath, "build-context", "", "Path to the build context directory (required)")
	cmd.Flags().StringVar(&lifecycleDir, "lifecycle-dir", "", "Directory containing per-package lifecycle.json files, structured as <dir>/<package>/lifecycle.json (required)")
	cmd.Flags().StringVar(&packages, "packages", "", "Comma-separated list of package names (required)")

	for _, flag := range []string{"dockerfile", "build-context", "lifecycle-dir", "packages"} {
		if err := cmd.MarkFlagRequired(flag); err != nil {
			panic(err)
		}
	}

	return cmd
}
