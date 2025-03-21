package giu

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mazznoer/csscolorparser"
	"github.com/napsy/go-css"
)

// Code in this file allows to apply style using a CSS stylesheet
// The intention of the code looks as follows:
// - The CSS stylesheet is parsed and the resulting rules are stored in giu context
//   NOTE: more than one CSS stylesheets can be parsed, however the app could panic if the same tag is present twice
// - CSSTagWidget allows to apply style to a specified layout
// - main tag allows to apply style to the whole application
//
// tools:
// css parser - for now github.com/napsy/go-css. it is a bit poor, but I don't think we need anything more
// css colors - github.com/mazznoer/csscolorparser
//
// docs: docs/css.md

// ErrCSSParse represents a CSS parsing error and includes details about what is failing.
type ErrCSSParse struct {
	What   string // description of what we are parsing
	Value  string // the value which failed
	Detail error  // (optional) error to add extra detail (i.e. result of calling another function like strconv.ParseFloat)
}

func (e ErrCSSParse) Error() string {
	errStr := fmt.Sprintf("unable to parse %s: %q", e.What, e.Value)

	if e.Detail != nil {
		errStr += fmt.Sprintf(" - %s", e.Detail.Error())
	}

	return errStr
}

// ParseCSSStyleSheet parses CSS stylesheet and stores the rules in giu context.
//
//nolint:gocognit // no
func ParseCSSStyleSheet(data []byte) error {
	// css does not support windows formatting
	// https://github.com/AllenDang/giu/issues/842
	data = []byte(strings.ReplaceAll(string(data), "\r\n", "\n"))

	stylesheet, err := css.Unmarshal(data)
	if err != nil {
		return fmt.Errorf("error marshaling CSS file: %w", err)
	}

	for rule, style := range stylesheet {
		setter := Style()

		for styleVarName, styleVarValue := range style {
			// convert style variable name to giu style variable name
			styleVarID, err := StyleVarIDString(styleVarName)

			if err == nil {
				if err := parseStyleVar(styleVarValue, func(v float32) {
					setter.SetStyleFloat(styleVarID, v)
				}, func(x, y float32) {
					setter.SetStyle(styleVarID, x, y)
				}); err != nil {
					return err
				}

				continue
			}

			styleColorID, err := StyleColorIDString(styleVarName)
			if err == nil {
				col, err := csscolorparser.Parse(styleVarValue)
				if err != nil {
					return ErrCSSParse{What: "color", Value: styleVarValue, Detail: err}
				}

				setter.SetColor(styleColorID, col)

				continue
			}

			stylePlotVarID, err := StylePlotVarIDString(styleVarName)
			if err == nil {
				if err := parseStyleVar(styleVarValue, func(v float32) {
					setter.SetPlotStyleFloat(stylePlotVarID, v)
				}, func(x, y float32) {
					setter.SetPlotStyle(stylePlotVarID, x, y)
				}); err != nil {
					return err
				}

				continue
			}

			stylePlotColorID, err := StylePlotColorIDString(styleVarName)
			if err == nil {
				col, err := csscolorparser.Parse(styleVarValue)
				if err != nil {
					return ErrCSSParse{What: "color", Value: styleVarValue, Detail: err}
				}

				setter.SetPlotColor(stylePlotColorID, col)

				continue
			}

			return ErrCSSParse{What: "style variable name", Value: styleVarName}
		}

		Context.cssStylesheet[string(rule)] = setter
	}

	return nil
}

func parseStyleVar(styleVarValue string, setFloat func(v float32), setVec2 func(x, y float32)) error {
	// the style is StyleVarID - set it
	f, err2 := strconv.ParseFloat(styleVarValue, 32)
	if err2 == nil {
		setFloat(float32(f))
		return nil
	}

	// so maybe it is a vec2 value:
	// var-name: x, y;
	styleVarValue = strings.ReplaceAll(styleVarValue, " ", "")
	vec2 := strings.Split(styleVarValue, ",")

	if len(vec2) != 2 {
		return ErrCSSParse{What: "value (not float or vec2)", Value: styleVarValue}
	}

	x, err2 := strconv.ParseFloat(vec2[0], 32)
	if err2 != nil {
		return ErrCSSParse{What: "value (not float)", Value: vec2[0], Detail: err2}
	}

	y, err2 := strconv.ParseFloat(vec2[1], 32)
	if err2 != nil {
		return ErrCSSParse{What: "value (not float)", Value: vec2[1], Detail: err2}
	}

	setVec2(float32(x), float32(y))

	return nil
}

// cssStylesheet is a map tag:StyleSetter.
type cssStylesheet map[string]*StyleSetter

var _ Widget = &CSSTagWidget{}

// CSSTagWidget is a widget that allows to apply CSS style to a specified layout.
type CSSTagWidget struct {
	tag    string
	layout Layout
}

// CSSTag creates CSSTagWidget.
func CSSTag(tag string) *CSSTagWidget {
	return &CSSTagWidget{tag: tag}
}

// To specifies a layout to which the style will be applied.
func (c *CSSTagWidget) To(layout ...Widget) *CSSTagWidget {
	c.layout = layout
	return c
}

// Build implements Widget interface.
func (c *CSSTagWidget) Build() {
	// get style from context.
	// if it doesn't exist Assert.
	style, exists := Context.cssStylesheet[c.tag]
	Assert(exists, "CSSTagWidget", "Build", "CSS stylesheet doesn't contain tag: %s", c.tag)

	style.To(c.layout).Build()
}
