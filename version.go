package main

import (
	"bytes"
	"fmt"
)

const (
	Name string = "fixla"

	// Version indicates which version of the binary is running.
	Version string = "0.1.0"

	// A pre-release marker for the version. If this is "" (empty string) the it means that it is a final release.
	// Otherwise, this is a pre-release such as "dev" (in development), "rc", etc.
	VersionPrerelease = "dev"
)

func FormattedVersion() string {
	var formattedVersion bytes.Buffer
	fmt.Fprintf(&formattedVersion, "%s version %s", Name, Version)
	if VersionPrerelease != "" {
		fmt.Fprintf(&formattedVersion, "-%s", VersionPrerelease)
	}
	return formattedVersion.String()
}
