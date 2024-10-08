// Code generated by "stringer -type=StyleColorID,StyleVarID -output=StyleIDs_string.go -linecomment"; DO NOT EDIT.

package giu

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StyleColorText-0]
	_ = x[StyleColorTextDisabled-1]
	_ = x[StyleColorWindowBg-2]
	_ = x[StyleColorChildBg-3]
	_ = x[StyleColorPopupBg-4]
	_ = x[StyleColorBorder-5]
	_ = x[StyleColorBorderShadow-6]
	_ = x[StyleColorFrameBg-7]
	_ = x[StyleColorFrameBgHovered-8]
	_ = x[StyleColorFrameBgActive-9]
	_ = x[StyleColorTitleBg-10]
	_ = x[StyleColorTitleBgActive-11]
	_ = x[StyleColorTitleBgCollapsed-12]
	_ = x[StyleColorMenuBarBg-13]
	_ = x[StyleColorScrollbarBg-14]
	_ = x[StyleColorScrollbarGrab-15]
	_ = x[StyleColorScrollbarGrabHovered-16]
	_ = x[StyleColorScrollbarGrabActive-17]
	_ = x[StyleColorCheckMark-18]
	_ = x[StyleColorSliderGrab-19]
	_ = x[StyleColorSliderGrabActive-20]
	_ = x[StyleColorButton-21]
	_ = x[StyleColorButtonHovered-22]
	_ = x[StyleColorButtonActive-23]
	_ = x[StyleColorHeader-24]
	_ = x[StyleColorHeaderHovered-25]
	_ = x[StyleColorHeaderActive-26]
	_ = x[StyleColorSeparator-27]
	_ = x[StyleColorSeparatorHovered-28]
	_ = x[StyleColorSeparatorActive-29]
	_ = x[StyleColorResizeGrip-30]
	_ = x[StyleColorResizeGripHovered-31]
	_ = x[StyleColorResizeGripActive-32]
	_ = x[StyleColorTab-34]
	_ = x[StyleColorTabHovered-33]
	_ = x[StyleColorTabActive-35]
	_ = x[StyleColorTabUnfocused-37]
	_ = x[StyleColorTabUnfocusedActive-38]
	_ = x[StyleColorPlotLines-42]
	_ = x[StyleColorPlotLinesHovered-43]
	_ = x[StyleColorPlotHistogram-44]
	_ = x[StyleColorPlotHistogramHovered-45]
	_ = x[StyleColorTableHeaderBg-46]
	_ = x[StyleColorTableBorderStrong-47]
	_ = x[StyleColorTableBorderLight-48]
	_ = x[StyleColorTableRowBg-49]
	_ = x[StyleColorTableRowBgAlt-50]
	_ = x[StyleColorTextSelectedBg-52]
	_ = x[StyleColorDragDropTarget-53]
	_ = x[StyleColorNavHighlight-54]
	_ = x[StyleColorNavWindowingHighlight-55]
	_ = x[StyleColorNavWindowingDimBg-56]
	_ = x[StyleColorModalWindowDimBg-57]
}

const (
	_StyleColorID_name_0 = "colordisabled-colorbackground-colorchild-background-colorpopup-background-colorborder-colorborder-shadow-colorframe-background-colorframe-background-hovered-colorframe-background-active-colortitle-background-colortitle-background-active-colortitle-background-collapsed-colormenu-bar-background-colorscrollbar-background-colorscrollbar-grab-colorscrollbar-grab-hovered-colorscrollbar-grab-active-colorcheckmark-colorslider-grab-colorslider-grab-active-colorbutton-colorbutton-hovered-colorbutton-active-colorheader-colorheader-hovered-colorheader-active-colorseparator-colorseparator-hovered-colorseparator-active-colorresize-grip-colorresize-grip-hovered-colorresize-grip-active-colortab-hovered-colortab-colortab-active-color"
	_StyleColorID_name_1 = "tab-unfocused-colortab-unfocused-active-color"
	_StyleColorID_name_2 = "plot-lines-colorplot-lines-hovered-colorplot-histogram-colorplot-histogram-hovered-colortable-header-background-colortable-border-strong-colortable-border-light-colortable-row-background-colortable-row-alternate-background-color"
	_StyleColorID_name_3 = "text-selected-background-colordrag-drop-target-colornavigation-highlight-colorwindowing-highlight-colorwindowing-dim-background-colormodal-window-dim-background-color"
)

var (
	_StyleColorID_index_0 = [...]uint16{0, 5, 19, 35, 57, 79, 91, 110, 132, 162, 191, 213, 242, 274, 299, 325, 345, 373, 400, 415, 432, 456, 468, 488, 507, 519, 539, 558, 573, 596, 618, 635, 660, 684, 701, 710, 726}
	_StyleColorID_index_1 = [...]uint8{0, 19, 45}
	_StyleColorID_index_2 = [...]uint8{0, 16, 40, 60, 88, 117, 142, 166, 192, 228}
	_StyleColorID_index_3 = [...]uint8{0, 30, 52, 78, 103, 133, 166}
)

func (i StyleColorID) String() string {
	switch {
	case 0 <= i && i <= 35:
		return _StyleColorID_name_0[_StyleColorID_index_0[i]:_StyleColorID_index_0[i+1]]
	case 37 <= i && i <= 38:
		i -= 37
		return _StyleColorID_name_1[_StyleColorID_index_1[i]:_StyleColorID_index_1[i+1]]
	case 42 <= i && i <= 50:
		i -= 42
		return _StyleColorID_name_2[_StyleColorID_index_2[i]:_StyleColorID_index_2[i+1]]
	case 52 <= i && i <= 57:
		i -= 52
		return _StyleColorID_name_3[_StyleColorID_index_3[i]:_StyleColorID_index_3[i+1]]
	default:
		return "StyleColorID(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StyleVarAlpha-0]
	_ = x[StyleVarDisabledAlpha-1]
	_ = x[StyleVarWindowPadding-2]
	_ = x[StyleVarWindowRounding-3]
	_ = x[StyleVarWindowBorderSize-4]
	_ = x[StyleVarWindowMinSize-5]
	_ = x[StyleVarWindowTitleAlign-6]
	_ = x[StyleVarChildRounding-7]
	_ = x[StyleVarChildBorderSize-8]
	_ = x[StyleVarPopupRounding-9]
	_ = x[StyleVarPopupBorderSize-10]
	_ = x[StyleVarFramePadding-11]
	_ = x[StyleVarFrameRounding-12]
	_ = x[StyleVarFrameBorderSize-13]
	_ = x[StyleVarItemSpacing-14]
	_ = x[StyleVarItemInnerSpacing-15]
	_ = x[StyleVarIndentSpacing-16]
	_ = x[StyleVarCellPadding-17]
	_ = x[StyleVarScrollbarSize-18]
	_ = x[StyleVarScrollbarRounding-19]
	_ = x[StyleVarGrabMinSize-20]
	_ = x[StyleVarGrabRounding-21]
	_ = x[StyleVarTabRounding-22]
	_ = x[StyleVarTabBarBorderSize-24]
	_ = x[StyleVarButtonTextAlign-28]
	_ = x[StyleVarSelectableTextAlign-29]
	_ = x[StyleVarSeparatorTextBorderSize-30]
	_ = x[StyleVarSeparatorTextAlign-31]
	_ = x[StyleVarSeparatorTextPadding-32]
	_ = x[StyleVarDockingSeparatorSize-33]
}

const (
	_StyleVarID_name_0 = "alphadisabled-alphawindow-paddingwindow-roundingwindow-border-sizewindow-minValue-sizewindow-title-alignchild-roundingchild-border-sizepopup-roundingpopup-border-sizeframe-paddingframe-roundingframe-border-sizeitem-spacingitem-inner-spacingindent-spacingStyleVarCellPaddingscrollbar-sizescrollbar-roundinggrab-minValue-sizegrab-roundingtab-rounding"
	_StyleVarID_name_1 = "StyleVarTabBarBorderSize"
	_StyleVarID_name_2 = "button-text-alignselectable-text-alignStyleVarSeparatorTextBorderSizeStyleVarSeparatorTextAlignStyleVarSeparatorTextPaddingStyleVarDockingSeparatorSize"
)

var (
	_StyleVarID_index_0 = [...]uint16{0, 5, 19, 33, 48, 66, 86, 104, 118, 135, 149, 166, 179, 193, 210, 222, 240, 254, 273, 287, 305, 323, 336, 348}
	_StyleVarID_index_2 = [...]uint8{0, 17, 38, 69, 95, 123, 151}
)

func (i StyleVarID) String() string {
	switch {
	case 0 <= i && i <= 22:
		return _StyleVarID_name_0[_StyleVarID_index_0[i]:_StyleVarID_index_0[i+1]]
	case i == 24:
		return _StyleVarID_name_1
	case 28 <= i && i <= 33:
		i -= 28
		return _StyleVarID_name_2[_StyleVarID_index_2[i]:_StyleVarID_index_2[i+1]]
	default:
		return "StyleVarID(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
