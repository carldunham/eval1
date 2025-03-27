package graph

// THIS CODE WILL BE UPDATED WITH SCHEMA CHANGES. PREVIOUS IMPLEMENTATION FOR SCHEMA CHANGES WILL BE KEPT IN THE COMMENT SECTION. IMPLEMENTATION FOR UNCHANGED SCHEMA WILL BE KEPT.

import (
	"context"
	"fmt"
	"nestmed/eval1/backend/graph/model"
	"regexp"
	"strings"
	"time"
)

type Resolver struct{}

// ProcessTranscript is the resolver for the processTranscript field.
func (r *mutationResolver) ProcessTranscript(ctx context.Context, transcript string) (*model.VisitSummary, error) {
	return r.Resolver.Query().AnalyzeTranscript(ctx, transcript)
}

// AnalyzeTranscript is the resolver for the analyzeTranscript field.
func (r *queryResolver) AnalyzeTranscript(ctx context.Context, transcript string) (*model.VisitSummary, error) {
	// Extract vital signs using regex patterns
	vitalSigns := &model.VitalSigns{}

	// Blood pressure pattern (e.g., "120/80")
	bpPattern := regexp.MustCompile(`(\d{2,3}/\d{2,3})`)
	if matches := bpPattern.FindStringSubmatch(transcript); len(matches) > 1 {
		vitalSigns.BloodPressure = &matches[1]
	}

	// Heart rate pattern (e.g., "HR: 72")
	hrPattern := regexp.MustCompile(`(?i)HR:?\s*(\d+)`)
	if matches := hrPattern.FindStringSubmatch(transcript); len(matches) > 1 {
		if hr, err := parseInt32(matches[1]); err == nil {
			vitalSigns.HeartRate = hr
		}
	}

	// Temperature pattern (e.g., "Temp: 98.6")
	tempPattern := regexp.MustCompile(`(?i)temp:?\s*(\d+\.?\d*)`)
	if matches := tempPattern.FindStringSubmatch(transcript); len(matches) > 1 {
		if temp, err := parseFloat64(matches[1]); err == nil {
			vitalSigns.Temperature = temp
		}
	}

	// Respiratory rate pattern (e.g., "RR: 16")
	rrPattern := regexp.MustCompile(`(?i)RR:?\s*(\d+)`)
	if matches := rrPattern.FindStringSubmatch(transcript); len(matches) > 1 {
		if rr, err := parseInt32(matches[1]); err == nil {
			vitalSigns.RespiratoryRate = rr
		}
	}

	// O2 saturation pattern (e.g., "O2 Sat: 98")
	o2Pattern := regexp.MustCompile(`(?i)O2\s*sat:?\s*(\d+)`)
	if matches := o2Pattern.FindStringSubmatch(transcript); len(matches) > 1 {
		if o2, err := parseInt32(matches[1]); err == nil {
			vitalSigns.OxygenSaturation = o2
		}
	}

	// Extract OASIS elements
	oasisElements := &model.OASISElement{}

	// Extract living situation (M0069)
	if strings.Contains(strings.ToLower(transcript), "lives alone") {
		alone := "Alone"
		oasisElements.M0069 = &alone
	} else if strings.Contains(strings.ToLower(transcript), "lives with") {
		withOthers := "With others"
		oasisElements.M0069 = &withOthers
	}

	// Extract visit type and duration
	visitType := "SOC" // Start of Care
	if strings.Contains(strings.ToLower(transcript), "recertification") {
		visitType = "Recertification"
	} else if strings.Contains(strings.ToLower(transcript), "follow-up") {
		visitType = "Follow-up"
	}

	// Create visit summary
	summary := &model.VisitSummary{
		VitalSigns:    vitalSigns,
		OasisElements: oasisElements,
		VisitDate:     time.Now().Format("2006-01-02"),
		VisitType:     visitType,
		VisitDuration: 60, // Default duration in minutes
		Notes:         &transcript,
	}

	return summary, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// Helper functions
func parseInt32(s string) (*int32, error) {
	var i int32
	n, err := fmt.Sscanf(s, "%d", &i)
	if err != nil || n != 1 {
		return nil, fmt.Errorf("failed to parse integer: %v", err)
	}
	return &i, nil
}

func parseFloat64(s string) (*float64, error) {
	var f float64
	n, err := fmt.Sscanf(s, "%f", &f)
	if err != nil || n != 1 {
		return nil, fmt.Errorf("failed to parse float: %v", err)
	}
	return &f, nil
}
