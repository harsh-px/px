/*
Copyright © 2019 Portworx

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

package util

import (
	"fmt"
	"os"
)

var (
	// Stdout points to the output buffer to send screen output
	Stdout = os.Stdout
	// Stderr points to the output buffer to send errors to the screen
	Stderr = os.Stderr
)

// Printf is just like fmt.Printf except that it send the output to Stdout. It
// is equal to fmt.Fprintf(util.Stdout, format, args)
func Printf(format string, args ...string) {
	fmt.Fprintf(Stdout, format, args)
}

// Eprintf prints the errors to the output buffer Stderr. It is equal to
// fmt.Fprintf(util.Stderr, format, args)
func Eprintf(format string, args ...string) {
	fmt.Fprintf(Stderr, format, args)
}
