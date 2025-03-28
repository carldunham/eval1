// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type OASISElement struct {
	M0069 *string `json:"m0069,omitempty"`
	M0102 *string `json:"m0102,omitempty"`
	M0110 *string `json:"m0110,omitempty"`
	M0140 *string `json:"m0140,omitempty"`
	M0150 *string `json:"m0150,omitempty"`
	M1030 *string `json:"m1030,omitempty"`
	M1033 *string `json:"m1033,omitempty"`
	M1034 *string `json:"m1034,omitempty"`
	M1036 *string `json:"m1036,omitempty"`
	M1040 *string `json:"m1040,omitempty"`
	M1046 *string `json:"m1046,omitempty"`
	M1051 *string `json:"m1051,omitempty"`
	M1056 *string `json:"m1056,omitempty"`
	M1058 *string `json:"m1058,omitempty"`
	M1060 *string `json:"m1060,omitempty"`
}

type Query struct {
}

type VisitSummary struct {
	VitalSigns    *VitalSigns   `json:"vitalSigns"`
	OasisElements *OASISElement `json:"oasisElements"`
	VisitDate     string        `json:"visitDate"`
	VisitType     string        `json:"visitType"`
	VisitDuration int32         `json:"visitDuration"`
	Notes         *string       `json:"notes,omitempty"`
	Summary       *string       `json:"summary,omitempty"`
}

type VitalSigns struct {
	BloodPressure    *string  `json:"bloodPressure,omitempty"`
	HeartRate        *int32   `json:"heartRate,omitempty"`
	Temperature      *float64 `json:"temperature,omitempty"`
	RespiratoryRate  *int32   `json:"respiratoryRate,omitempty"`
	OxygenSaturation *int32   `json:"oxygenSaturation,omitempty"`
	BloodSugar       *int32   `json:"bloodSugar,omitempty"`
}
