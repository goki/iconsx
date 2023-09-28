// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package icons provides Go constant names for Material Design Symbols as SVG files.
package icons

import (
	"embed"
	"path/filepath"
	"strings"

	_ "github.com/iancoleman/strcase" // needed so that it gets included in the mod (the generator uses it)
	"goki.dev/glop/dirs"
)

//go:generate go run gen.go

// Icons contains all of the default embedded svg icons
//
//go:embed svg/*.svg
var Icons embed.FS

// Icon contains the name of an icon
type Icon string

// Fill returns the icon as a filled icon.
// It returns the icon unchanged if it is already filled.
func (i Icon) Fill() Icon {
	if i.IsFilled() {
		return i
	}
	return i + "-fill"
}

// Unfill returns the icon as an unfilled icon.
// It returns the icon unchanged if it is already unfilled.
// Icons are unfilled by default, so you only
// need to call this to reverse a prior [Icon.Fill] call
func (i Icon) Unfill() Icon {
	return Icon(strings.TrimSuffix(string(i), "-fill"))
}

// IsFilled returns whether the icon
// is a filled icon.
func (i Icon) IsFilled() bool {
	return strings.HasSuffix(string(i), "-fill")
}

// IsNil returns whether the icon name is empty,
// [None], or "nil"; those indicate not to use an icon.
func (i Icon) IsNil() bool {
	return i == "" || i == None || i == "nil"
}

func (i Icon) Filename() string {
	return filepath.Join("svg", string(i)) + ".svg"
}

// IsValid returns whether the icon name corresponds to
// a valid existing icon.
func (i Icon) IsValid() bool {
	if i.IsNil() {
		return false
	}
	ex, _ := dirs.FileExistsFS(Icons, i.Filename())
	return ex
}

// None is an icon that indicates to not use an icon.
// It completely prevents the rendering of an icon,
// whereas [Blank] renders a blank icon.
const None Icon = "none"

// Blank is a blank icon that can be used as a
// placeholder when no other icon is appropriate.
// It still renders an icon, just a blank one,
// whereas [None] indicates to not render one at all.
const Blank Icon = "blank"
