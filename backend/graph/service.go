package graph

import (
	"context"
	"time"

	"github.com/carldunham/nestmed/eval1/backend/graph/model"
	"github.com/carldunham/nestmed/eval1/backend/openai"
)

type Service struct {
	openaiClient openai.ClientInterface
}

func NewService() (*Service, error) {
	client, err := openai.NewClient()
	if err != nil {
		return nil, err
	}
	return &Service{openaiClient: client}, nil
}

func (s *Service) AnalyzeTranscript(ctx context.Context, transcript string) (*model.VisitSummary, error) {
	extractedData, err := s.openaiClient.ExtractData(ctx, transcript)
	if err != nil {
		return nil, err
	}

	// Convert int to int32 for the model
	var heartRate, respiratoryRate, oxygenSaturation *int32
	if extractedData.VitalSigns.HeartRate != nil {
		hr := int32(*extractedData.VitalSigns.HeartRate)
		heartRate = &hr
	}
	if extractedData.VitalSigns.RespiratoryRate != nil {
		rr := int32(*extractedData.VitalSigns.RespiratoryRate)
		respiratoryRate = &rr
	}
	if extractedData.VitalSigns.OxygenSaturation != nil {
		o2 := int32(*extractedData.VitalSigns.OxygenSaturation)
		oxygenSaturation = &o2
	}

	// Create visit summary from extracted data
	summary := &model.VisitSummary{
		VitalSigns: &model.VitalSigns{
			BloodPressure:    extractedData.VitalSigns.BloodPressure,
			HeartRate:        heartRate,
			Temperature:      extractedData.VitalSigns.Temperature,
			RespiratoryRate:  respiratoryRate,
			OxygenSaturation: oxygenSaturation,
		},
		OasisElements: &model.OASISElement{
			M0069: extractedData.OasisElements.M0069,
			M0102: extractedData.OasisElements.M0102,
			M0110: extractedData.OasisElements.M0110,
			M0140: extractedData.OasisElements.M0140,
			M0150: extractedData.OasisElements.M0150,
			M1030: extractedData.OasisElements.M1030,
			M1033: extractedData.OasisElements.M1033,
			M1034: extractedData.OasisElements.M1034,
			M1036: extractedData.OasisElements.M1036,
			M1040: extractedData.OasisElements.M1040,
			M1046: extractedData.OasisElements.M1046,
			M1051: extractedData.OasisElements.M1051,
			M1056: extractedData.OasisElements.M1056,
			M1058: extractedData.OasisElements.M1058,
			M1060: extractedData.OasisElements.M1060,
		},
		VisitDate:     time.Now().Format("2006-01-02"),
		VisitType:     getVisitType(extractedData.VisitType),
		VisitDuration: 60, // Default duration in minutes
		Notes:         &transcript,
	}

	return summary, nil
}

func getVisitType(visitType *string) string {
	if visitType == nil {
		return "SOC" // Default to Start of Care if not specified
	}
	return *visitType
}
