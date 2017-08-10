package bitmovintypes

type ConditionAttribute string

const (
	ConditionAttributeHeight      ConditionAttribute = "HEIGHT"
	ConditionAttributeWidth       ConditionAttribute = "WIDTH"
	ConditionAttributeFPS         ConditionAttribute = "FPS"
	ConditionAttributeBitrate     ConditionAttribute = "BITRATE"
	ConditionAttributeAspectRatio ConditionAttribute = "ASPECTRATIO"
)

type ConditionType string

const (
	ConditionTypeAnd       ConditionType = "AND"
	ConditionTypeOr        ConditionType = "OR"
	ConditionTypeCondition ConditionType = "CONDITION"
)
