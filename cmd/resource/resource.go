/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
package resource

import (
	"github.com/spf13/cobra"
)

var (
	filename string
)

// CreateResourceCmd is for inputting a new resource
var CreateResourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "Create a new resource",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// UpdateResourceCmd is for updating an existing resource
var UpdateResourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "Update an existing resource",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// DeleteResourceCmd is for delete an existing resource
var DeleteResourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "Delete an existing resource",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {

	// CreateResourceCmd
	CreateResourceCmd.Flags().StringVarP(&filename, "filename", "f", "", "Add a resource file")
	// DeleteResourceCmd
	UpdateResourceCmd.Flags().StringVarP(&filename, "filename", "f", "", "Update a resource file")
	// UpdateResourceCmd
	DeleteResourceCmd.Flags().StringVarP(&filename, "filename", "f", "", "Delete a resource file")
}
