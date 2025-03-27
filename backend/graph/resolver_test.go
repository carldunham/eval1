package graph

import (
	"context"
	"testing"

	"github.com/carldunham/nestmed/eval1/backend/openai"
)

func TestAnalyzeTranscript(t *testing.T) {
	testClient := NewTestOpenAIClient()
	service := &Service{openaiClient: testClient}
	resolver := &Resolver{service: service}
	ctx := context.Background()

	tests := []struct {
		name       string
		transcript string
		wantBP     string
		wantHR     int32
		wantTemp   float64
		wantRR     int32
		wantO2     int32
		wantBS     int32
		wantType   string
		wantM0069  string
		wantM0102  string
		wantM0110  string
		wantM0140  string
		wantM0150  string
		wantM1030  string
		wantM1033  string
		wantM1034  string
		wantM1036  string
		wantM1040  string
		wantM1046  string
		wantM1051  string
		wantM1056  string
		wantM1058  string
		wantM1060  string
	}{
		{
			name:       "Complete vital signs and OASIS elements",
			transcript: "Patient presented with BP 120/80, HR: 72, Temp: 98.6, RR: 16, O2 Sat: 98, Blood Sugar: 95. Follow-up visit. Patient lives alone in a single-family home. Primary language is English. Patient is Caucasian. Current health status is stable. High risk for hospitalization due to recent surgery. Moderate risk for falls.",
			wantBP:     "120/80",
			wantHR:     72,
			wantTemp:   98.6,
			wantRR:     16,
			wantO2:     98,
			wantBS:     95,
			wantType:   "Follow-up",
			wantM0069:  "Alone",
			wantM0102:  "English",
			wantM0110:  "Not Hispanic or Latino",
			wantM0140:  "Caucasian",
			wantM0150:  "Stable",
			wantM1030:  "High",
			wantM1033:  "Low",
			wantM1034:  "Low",
			wantM1036:  "Moderate",
			wantM1040:  "Low",
			wantM1046:  "Low",
			wantM1051:  "Low",
			wantM1056:  "Moderate",
			wantM1058:  "Low",
			wantM1060:  "Low",
		},
		{
			name:       "Partial vital signs and OASIS elements",
			transcript: "Patient presented with BP 140/90, HR: 85, Blood Sugar: 180. Start of care visit. Patient lives with family. Primary language is Spanish. Patient is Hispanic. Current health status is declining. High risk for pressure ulcers.",
			wantBP:     "140/90",
			wantHR:     85,
			wantBS:     180,
			wantType:   "SOC",
			wantM0069:  "With family",
			wantM0102:  "Spanish",
			wantM0110:  "Hispanic or Latino",
			wantM0140:  "Hispanic",
			wantM0150:  "Declining",
			wantM1030:  "Low",
			wantM1033:  "Low",
			wantM1034:  "High",
			wantM1036:  "Low",
			wantM1040:  "Low",
			wantM1046:  "Low",
			wantM1051:  "High",
			wantM1056:  "Low",
			wantM1058:  "Low",
			wantM1060:  "Low",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up test response
			testClient.SetExtractDataFunc(func(ctx context.Context, transcript string) (*openai.ExtractedData, error) {
				return &openai.ExtractedData{
					VitalSigns: struct {
						BloodPressure    *string  `json:"bloodPressure"`
						HeartRate        *int     `json:"heartRate"`
						Temperature      *float64 `json:"temperature"`
						RespiratoryRate  *int     `json:"respiratoryRate"`
						OxygenSaturation *int     `json:"oxygenSaturation"`
						BloodSugar       *int     `json:"bloodSugar"`
					}{
						BloodPressure:    &tt.wantBP,
						HeartRate:        intPtr(int(tt.wantHR)),
						Temperature:      floatPtr(tt.wantTemp),
						RespiratoryRate:  intPtr(int(tt.wantRR)),
						OxygenSaturation: intPtr(int(tt.wantO2)),
						BloodSugar:       intPtr(int(tt.wantBS)),
					},
					OasisElements: struct {
						M0069 *string `json:"m0069"`
						M0102 *string `json:"m0102"`
						M0110 *string `json:"m0110"`
						M0140 *string `json:"m0140"`
						M0150 *string `json:"m0150"`
						M1030 *string `json:"m1030"`
						M1033 *string `json:"m1033"`
						M1034 *string `json:"m1034"`
						M1036 *string `json:"m1036"`
						M1040 *string `json:"m1040"`
						M1046 *string `json:"m1046"`
						M1051 *string `json:"m1051"`
						M1056 *string `json:"m1056"`
						M1058 *string `json:"m1058"`
						M1060 *string `json:"m1060"`
					}{
						M0069: &tt.wantM0069,
						M0102: &tt.wantM0102,
						M0110: &tt.wantM0110,
						M0140: &tt.wantM0140,
						M0150: &tt.wantM0150,
						M1030: &tt.wantM1030,
						M1033: &tt.wantM1033,
						M1034: &tt.wantM1034,
						M1036: &tt.wantM1036,
						M1040: &tt.wantM1040,
						M1046: &tt.wantM1046,
						M1051: &tt.wantM1051,
						M1056: &tt.wantM1056,
						M1058: &tt.wantM1058,
						M1060: &tt.wantM1060,
					},
					VisitType: &tt.wantType,
				}, nil
			})

			got, err := resolver.Query().AnalyzeTranscript(ctx, tt.transcript)
			if err != nil {
				t.Errorf("AnalyzeTranscript() error = %v", err)
				return
			}

			if got.VitalSigns.BloodPressure != nil && *got.VitalSigns.BloodPressure != tt.wantBP {
				t.Errorf("BloodPressure = %v, want %v", *got.VitalSigns.BloodPressure, tt.wantBP)
			}
			if got.VitalSigns.HeartRate != nil && *got.VitalSigns.HeartRate != tt.wantHR {
				t.Errorf("HeartRate = %v, want %v", *got.VitalSigns.HeartRate, tt.wantHR)
			}
			if got.VitalSigns.Temperature != nil && *got.VitalSigns.Temperature != tt.wantTemp {
				t.Errorf("Temperature = %v, want %v", *got.VitalSigns.Temperature, tt.wantTemp)
			}
			if got.VitalSigns.RespiratoryRate != nil && *got.VitalSigns.RespiratoryRate != tt.wantRR {
				t.Errorf("RespiratoryRate = %v, want %v", *got.VitalSigns.RespiratoryRate, tt.wantRR)
			}
			if got.VitalSigns.OxygenSaturation != nil && *got.VitalSigns.OxygenSaturation != tt.wantO2 {
				t.Errorf("OxygenSaturation = %v, want %v", *got.VitalSigns.OxygenSaturation, tt.wantO2)
			}
			if got.VitalSigns.BloodSugar != nil && *got.VitalSigns.BloodSugar != tt.wantBS {
				t.Errorf("BloodSugar = %v, want %v", *got.VitalSigns.BloodSugar, tt.wantBS)
			}
			if got.VisitType != tt.wantType {
				t.Errorf("VisitType = %v, want %v", got.VisitType, tt.wantType)
			}
			if got.OasisElements.M0069 != nil && *got.OasisElements.M0069 != tt.wantM0069 {
				t.Errorf("M0069 = %v, want %v", *got.OasisElements.M0069, tt.wantM0069)
			}
			if got.OasisElements.M0102 != nil && *got.OasisElements.M0102 != tt.wantM0102 {
				t.Errorf("M0102 = %v, want %v", *got.OasisElements.M0102, tt.wantM0102)
			}
			if got.OasisElements.M0110 != nil && *got.OasisElements.M0110 != tt.wantM0110 {
				t.Errorf("M0110 = %v, want %v", *got.OasisElements.M0110, tt.wantM0110)
			}
			if got.OasisElements.M0140 != nil && *got.OasisElements.M0140 != tt.wantM0140 {
				t.Errorf("M0140 = %v, want %v", *got.OasisElements.M0140, tt.wantM0140)
			}
			if got.OasisElements.M0150 != nil && *got.OasisElements.M0150 != tt.wantM0150 {
				t.Errorf("M0150 = %v, want %v", *got.OasisElements.M0150, tt.wantM0150)
			}
			if got.OasisElements.M1030 != nil && *got.OasisElements.M1030 != tt.wantM1030 {
				t.Errorf("M1030 = %v, want %v", *got.OasisElements.M1030, tt.wantM1030)
			}
			if got.OasisElements.M1033 != nil && *got.OasisElements.M1033 != tt.wantM1033 {
				t.Errorf("M1033 = %v, want %v", *got.OasisElements.M1033, tt.wantM1033)
			}
			if got.OasisElements.M1034 != nil && *got.OasisElements.M1034 != tt.wantM1034 {
				t.Errorf("M1034 = %v, want %v", *got.OasisElements.M1034, tt.wantM1034)
			}
			if got.OasisElements.M1036 != nil && *got.OasisElements.M1036 != tt.wantM1036 {
				t.Errorf("M1036 = %v, want %v", *got.OasisElements.M1036, tt.wantM1036)
			}
			if got.OasisElements.M1040 != nil && *got.OasisElements.M1040 != tt.wantM1040 {
				t.Errorf("M1040 = %v, want %v", *got.OasisElements.M1040, tt.wantM1040)
			}
			if got.OasisElements.M1046 != nil && *got.OasisElements.M1046 != tt.wantM1046 {
				t.Errorf("M1046 = %v, want %v", *got.OasisElements.M1046, tt.wantM1046)
			}
			if got.OasisElements.M1051 != nil && *got.OasisElements.M1051 != tt.wantM1051 {
				t.Errorf("M1051 = %v, want %v", *got.OasisElements.M1051, tt.wantM1051)
			}
			if got.OasisElements.M1056 != nil && *got.OasisElements.M1056 != tt.wantM1056 {
				t.Errorf("M1056 = %v, want %v", *got.OasisElements.M1056, tt.wantM1056)
			}
			if got.OasisElements.M1058 != nil && *got.OasisElements.M1058 != tt.wantM1058 {
				t.Errorf("M1058 = %v, want %v", *got.OasisElements.M1058, tt.wantM1058)
			}
			if got.OasisElements.M1060 != nil && *got.OasisElements.M1060 != tt.wantM1060 {
				t.Errorf("M1060 = %v, want %v", *got.OasisElements.M1060, tt.wantM1060)
			}
		})
	}
}

func intPtr(i int) *int {
	return &i
}

func floatPtr(f float64) *float64 {
	return &f
}
