import React, { useState } from 'react';
import { ApolloClient, InMemoryCache, ApolloProvider, gql, useMutation } from '@apollo/client';
import {
  Container,
  TextField,
  Button,
  Paper,
  Typography,
  Box,
  CircularProgress,
  Grid,
} from '@mui/material';

const client = new ApolloClient({
  uri: 'http://localhost:8085/query',
  cache: new InMemoryCache(),
});

const ANALYZE_TRANSCRIPT = gql`
  mutation AnalyzeTranscript($transcript: String!) {
    processTranscript(transcript: $transcript) {
      vitalSigns {
        bloodPressure
        heartRate
        temperature
        respiratoryRate
        oxygenSaturation
        bloodSugar
      }
      oasisElements {
        m0069
        m0102
        m0110
        m0140
        m0150
        m1030
        m1033
        m1034
        m1036
        m1040
        m1046
        m1051
        m1056
        m1058
        m1060
      }
      visitDate
      visitType
      visitDuration
      notes
    }
  }
`;

function TranscriptAnalyzer() {
  const [transcript, setTranscript] = useState('');
  const [analyzeTranscript, { loading, error, data }] = useMutation(ANALYZE_TRANSCRIPT);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await analyzeTranscript({ variables: { transcript } });
    } catch (err) {
      console.error('Error analyzing transcript:', err);
    }
  };

  return (
    <Container maxWidth="lg">
      <Box sx={{ my: 4 }}>
        <Typography variant="h4" component="h1" gutterBottom>
          Home Health Visit Transcript Analyzer
        </Typography>

        <form onSubmit={handleSubmit}>
          <TextField
            fullWidth
            multiline
            rows={6}
            variant="outlined"
            label="Enter Visit Transcript"
            value={transcript}
            onChange={(e) => setTranscript(e.target.value)}
            margin="normal"
            disabled={loading}
          />
          <Button
            type="submit"
            variant="contained"
            color="primary"
            disabled={loading || !transcript.trim()}
            sx={{ mt: 2 }}
          >
            {loading ? <CircularProgress size={24} /> : 'Analyze Transcript'}
          </Button>
        </form>

        {error && (
          <Paper sx={{ p: 2, mt: 2, bgcolor: 'error.light' }}>
            <Typography color="error">Error: {error.message}</Typography>
          </Paper>
        )}

        {data?.processTranscript && (
          <Paper sx={{ p: 3, mt: 3 }}>
            <Typography variant="h5" gutterBottom>
              Analysis Results
            </Typography>

            <Grid container spacing={3}>
              <Grid item xs={12} md={6}>
                <Typography variant="h6" gutterBottom>
                  Vital Signs
                </Typography>
                <Box component="dl">
                  <Box sx={{ display: 'flex', mb: 1 }}>
                    <Typography component="dt" sx={{ fontWeight: 'bold', mr: 1 }}>
                      Blood Pressure:
                    </Typography>
                    <Typography component="dd">
                      {data.processTranscript.vitalSigns.bloodPressure || 'N/A'}
                    </Typography>
                  </Box>
                  <Box sx={{ display: 'flex', mb: 1 }}>
                    <Typography component="dt" sx={{ fontWeight: 'bold', mr: 1 }}>
                      Heart Rate:
                    </Typography>
                    <Typography component="dd">
                      {data.processTranscript.vitalSigns.heartRate || 'N/A'} bpm
                    </Typography>
                  </Box>
                  <Box sx={{ display: 'flex', mb: 1 }}>
                    <Typography component="dt" sx={{ fontWeight: 'bold', mr: 1 }}>
                      Temperature:
                    </Typography>
                    <Typography component="dd">
                      {data.processTranscript.vitalSigns.temperature || 'N/A'} °F
                    </Typography>
                  </Box>
                  <Box sx={{ display: 'flex', mb: 1 }}>
                    <Typography component="dt" sx={{ fontWeight: 'bold', mr: 1 }}>
                      Respiratory Rate:
                    </Typography>
                    <Typography component="dd">
                      {data.processTranscript.vitalSigns.respiratoryRate || 'N/A'} breaths/min
                    </Typography>
                  </Box>
                  <Box sx={{ display: 'flex', mb: 1 }}>
                    <Typography component="dt" sx={{ fontWeight: 'bold', mr: 1 }}>
                      O2 Saturation:
                    </Typography>
                    <Typography component="dd">
                      {data.processTranscript.vitalSigns.oxygenSaturation || 'N/A'}%
                    </Typography>
                  </Box>
                  <Box sx={{ display: 'flex', mb: 1 }}>
                    <Typography component="dt" sx={{ fontWeight: 'bold', mr: 1 }}>
                      Blood Sugar:
                    </Typography>
                    <Typography component="dd">
                      {data.processTranscript.vitalSigns.bloodSugar || 'N/A'} mg/dL
                    </Typography>
                  </Box>
                </Box>
              </Grid>

              <Grid item xs={12} md={6}>
                <Typography variant="h6" gutterBottom>
                  Visit Information
                </Typography>
                <Box component="dl">
                  <Box sx={{ display: 'flex', mb: 1 }}>
                    <Typography component="dt" sx={{ fontWeight: 'bold', mr: 1 }}>
                      Visit Type:
                    </Typography>
                    <Typography component="dd">
                      {data.processTranscript.visitType}
                    </Typography>
                  </Box>
                  <Box sx={{ display: 'flex', mb: 1 }}>
                    <Typography component="dt" sx={{ fontWeight: 'bold', mr: 1 }}>
                      Visit Date:
                    </Typography>
                    <Typography component="dd">
                      {data.processTranscript.visitDate}
                    </Typography>
                  </Box>
                  <Box sx={{ display: 'flex', mb: 1 }}>
                    <Typography component="dt" sx={{ fontWeight: 'bold', mr: 1 }}>
                      Duration:
                    </Typography>
                    <Typography component="dd">
                      {data.processTranscript.visitDuration} minutes
                    </Typography>
                  </Box>
                </Box>
              </Grid>

              <Grid item xs={12}>
                <Typography variant="h6" gutterBottom>
                  OASIS Elements
                </Typography>
                <Grid container spacing={2}>
                  <Grid item xs={12} md={6} lg={4}>
                    <Box component="dl">
                      <Box sx={{ display: 'flex', mb: 1 }}>
                        <Typography component="dt" sx={{ fontWeight: 'bold', mr: 1 }}>
                          Living Situation (M0069):
                        </Typography>
                        <Typography component="dd">
                          {data.processTranscript.oasisElements.m0069 || 'N/A'}
                        </Typography>
                      </Box>
                      <Box sx={{ display: 'flex', mb: 1 }}>
                        <Typography component="dt" sx={{ fontWeight: 'bold', mr: 1 }}>
                          Primary Language (M0102):
                        </Typography>
                        <Typography component="dd">
                          {data.processTranscript.oasisElements.m0102 || 'N/A'}
                        </Typography>
                      </Box>
                      <Box sx={{ display: 'flex', mb: 1 }}>
                        <Typography component="dt" sx={{ fontWeight: 'bold', mr: 1 }}>
                          Ethnicity (M0110):
                        </Typography>
                        <Typography component="dd">
                          {data.processTranscript.oasisElements.m0110 || 'N/A'}
                        </Typography>
                      </Box>
                      <Box sx={{ display: 'flex', mb: 1 }}>
                        <Typography component="dt" sx={{ fontWeight: 'bold', mr: 1 }}>
                          Race (M0140):
                        </Typography>
                        <Typography component="dd">
                          {data.processTranscript.oasisElements.m0140 || 'N/A'}
                        </Typography>
                      </Box>
                      <Box sx={{ display: 'flex', mb: 1 }}>
                        <Typography component="dt" sx={{ fontWeight: 'bold', mr: 1 }}>
                          Health Status (M0150):
                        </Typography>
                        <Typography component="dd">
                          {data.processTranscript.oasisElements.m0150 || 'N/A'}
                        </Typography>
                      </Box>
                    </Box>
                  </Grid>
                  <Grid item xs={12} md={6} lg={4}>
                    <Box component="dl">
                      <Box sx={{ display: 'flex', mb: 1 }}>
                        <Typography component="dt" sx={{ fontWeight: 'bold', mr: 1 }}>
                          Hospitalization Risk (M1030):
                        </Typography>
                        <Typography component="dd">
                          {data.processTranscript.oasisElements.m1030 || 'N/A'}
                        </Typography>
                      </Box>
                      <Box sx={{ display: 'flex', mb: 1 }}>
                        <Typography component="dt" sx={{ fontWeight: 'bold', mr: 1 }}>
                          Death Risk (M1033):
                        </Typography>
                        <Typography component="dd">
                          {data.processTranscript.oasisElements.m1033 || 'N/A'}
                        </Typography>
                      </Box>
                      <Box sx={{ display: 'flex', mb: 1 }}>
                        <Typography component="dt" sx={{ fontWeight: 'bold', mr: 1 }}>
                          Pressure Ulcer Risk (M1034):
                        </Typography>
                        <Typography component="dd">
                          {data.processTranscript.oasisElements.m1034 || 'N/A'}
                        </Typography>
                      </Box>
                      <Box sx={{ display: 'flex', mb: 1 }}>
                        <Typography component="dt" sx={{ fontWeight: 'bold', mr: 1 }}>
                          Falls Risk (M1036):
                        </Typography>
                        <Typography component="dd">
                          {data.processTranscript.oasisElements.m1036 || 'N/A'}
                        </Typography>
                      </Box>
                      <Box sx={{ display: 'flex', mb: 1 }}>
                        <Typography component="dt" sx={{ fontWeight: 'bold', mr: 1 }}>
                          Depression Risk (M1040):
                        </Typography>
                        <Typography component="dd">
                          {data.processTranscript.oasisElements.m1040 || 'N/A'}
                        </Typography>
                      </Box>
                    </Box>
                  </Grid>
                  <Grid item xs={12} md={6} lg={4}>
                    <Box component="dl">
                      <Box sx={{ display: 'flex', mb: 1 }}>
                        <Typography component="dt" sx={{ fontWeight: 'bold', mr: 1 }}>
                          Weight Loss Risk (M1046):
                        </Typography>
                        <Typography component="dd">
                          {data.processTranscript.oasisElements.m1046 || 'N/A'}
                        </Typography>
                      </Box>
                      <Box sx={{ display: 'flex', mb: 1 }}>
                        <Typography component="dt" sx={{ fontWeight: 'bold', mr: 1 }}>
                          Pressure Ulcer Risk (M1051):
                        </Typography>
                        <Typography component="dd">
                          {data.processTranscript.oasisElements.m1051 || 'N/A'}
                        </Typography>
                      </Box>
                      <Box sx={{ display: 'flex', mb: 1 }}>
                        <Typography component="dt" sx={{ fontWeight: 'bold', mr: 1 }}>
                          Falls Risk (M1056):
                        </Typography>
                        <Typography component="dd">
                          {data.processTranscript.oasisElements.m1056 || 'N/A'}
                        </Typography>
                      </Box>
                      <Box sx={{ display: 'flex', mb: 1 }}>
                        <Typography component="dt" sx={{ fontWeight: 'bold', mr: 1 }}>
                          Depression Risk (M1058):
                        </Typography>
                        <Typography component="dd">
                          {data.processTranscript.oasisElements.m1058 || 'N/A'}
                        </Typography>
                      </Box>
                      <Box sx={{ display: 'flex', mb: 1 }}>
                        <Typography component="dt" sx={{ fontWeight: 'bold', mr: 1 }}>
                          Weight Loss Risk (M1060):
                        </Typography>
                        <Typography component="dd">
                          {data.processTranscript.oasisElements.m1060 || 'N/A'}
                        </Typography>
                      </Box>
                    </Box>
                  </Grid>
                </Grid>
              </Grid>
            </Grid>
          </Paper>
        )}
      </Box>
    </Container>
  );
}

function App() {
  return (
    <ApolloProvider client={client}>
      <TranscriptAnalyzer />
    </ApolloProvider>
  );
}

export default App;
