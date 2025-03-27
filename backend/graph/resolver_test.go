package graph

import (
	"context"
	"testing"
)

func TestAnalyzeTranscript(t *testing.T) {
	resolver := &Resolver{}
	ctx := context.Background()

	tests := []struct {
		name       string
		transcript string
		wantBP     string
		wantHR     int32
		wantTemp   float64
		wantRR     int32
		wantO2     int32
		wantType   string
	}{
		{
			name:       "Complete vital signs",
			transcript: "Patient presented with BP 120/80, HR: 72, Temp: 98.6, RR: 16, O2 Sat: 98. Follow-up visit.",
			wantBP:     "120/80",
			wantHR:     72,
			wantTemp:   98.6,
			wantRR:     16,
			wantO2:     98,
			wantType:   "Follow-up",
		},
		{
			name:       "Partial vital signs",
			transcript: "Patient presented with BP 140/90, HR: 85. Start of care visit.",
			wantBP:     "140/90",
			wantHR:     85,
			wantType:   "SOC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
			if got.VisitType != tt.wantType {
				t.Errorf("VisitType = %v, want %v", got.VisitType, tt.wantType)
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
