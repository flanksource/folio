// Copyright 2026 Carlos Munoz and the Folio Authors
// SPDX-License-Identifier: Apache-2.0

package html

import (
	"github.com/carlos7ags/folio/font"
	"github.com/carlos7ags/folio/layout"
)

// ComputedStyle is the exported alias for computedStyle.
type ComputedStyle = computedStyle

// DefaultStyle returns browser-like defaults (exported wrapper).
func DefaultStyle() ComputedStyle {
	return defaultStyle()
}

// Inherit creates a child style that inherits text properties from the parent.
func (s *ComputedStyle) Inherit() ComputedStyle {
	return s.inherit()
}

// HasPadding returns true if any padding is set.
func HasPadding(s ComputedStyle) bool {
	return s.hasPadding()
}

// HasBorder returns true if any border is set.
func HasBorder(s ComputedStyle) bool {
	return s.hasBorder()
}

// HasMargin returns true if any margin is set.
func HasMargin(s ComputedStyle) bool {
	return s.hasMargin()
}

// ResolveFont selects the standard PDF font matching the style's font family, weight, and style.
func ResolveFont(style ComputedStyle) *font.Standard {
	return resolveFont(style)
}

// ApplyDivStyles applies the computed style's padding, borders, margins,
// background, and dimensions to a Div element.
func ApplyDivStyles(div *layout.Div, style ComputedStyle, containerWidth float64) {
	applyDivStyles(div, style, containerWidth)
}

// BuildCellBorders creates layout.CellBorders from a computed style.
func BuildCellBorders(style ComputedStyle) layout.CellBorders {
	return buildCellBorders(style)
}

// BuildBorder creates a single layout.Border from width, style, and color.
func BuildBorder(width float64, style string, color layout.Color) layout.Border {
	return buildBorder(width, style, color)
}

// ParseColor parses a CSS color value into a layout.Color.
// Supports: named colors, #RGB, #RRGGBB, rgb(r,g,b).
func ParseColor(value string) (layout.Color, bool) {
	return parseColor(value)
}
