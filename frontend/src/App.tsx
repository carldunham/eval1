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
