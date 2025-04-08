package cloudattestation

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/edgelesssys/go-azguestattestation/maa"
)

const maaURL = "https://sharedeus.eus.attest.azure.net"

func generateNonce() ([]byte, error) {
	nonce := make([]byte, 16)
	_, err := rand.Read(nonce)
	if err != nil {
		return nil, err
	}
	return nonce, nil
}

func main() {
	fmt.Println("Hello, fetching attestation report.")

	nonce, err := generateNonce()
	if err != nil {
		fmt.Printf("generate nonce: %s\n", err.Error())
		return
	}

	fmt.Printf("Nonce: %s\n", hex.EncodeToString(nonce))
	fmt.Printf("\nTesting Attest form go-azure\n\n")

	maa.OSBuild = "UVC"
	maa.OSType = "Linux"
	maa.OSDistro = "UVC"
	token, err := maa.Attest(context.Background(), nonce, maaURL, http.DefaultClient)
	if err != nil {
		fmt.Printf("error for fetching azure token: %s\n", err.Error())
		return
	}

	fmt.Printf("Token: %s\n", token)
}

// package cloudattestation

// import (
// 	"bytes"
// 	"fmt"
// 	"os"
// 	"os/exec"
// )

// // ConvertAttestationReportFileToToken takes a file path to an attestation report (JSON)
// // and runs `gotpm token` to generate an attestation token.
// func ConvertAttestationReportFileToToken(reportFilePath string) (string, error) {
// 	// Check if file exists
// 	if _, err := os.Stat(reportFilePath); os.IsNotExist(err) {
// 		return "", fmt.Errorf("attestation report file not found: %s", reportFilePath)
// 	}

// 	// Run `gotpm token --input <file>`
// 	cmd := exec.Command("gotpm", "token", "--input", reportFilePath)

// 	// Capture the output (attestation token)
// 	var out bytes.Buffer
// 	cmd.Stdout = &out
// 	cmd.Stderr = os.Stderr // Print errors to terminal

// 	err := cmd.Run()
// 	if err != nil {
// 		return "", fmt.Errorf("error running gotpm token: %v", err)
// 	}

// 	// Return the generated attestation token
// 	return out.String(), nil
// }

// func main() {
// 	// Example usage: Provide the path to your attestation report file
// 	reportFilePath := "attestation_snp_vtpm_gcp.json" // Replace with the actual file path

// 	token, err := ConvertAttestationReportFileToToken(reportFilePath)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	fmt.Println("Attestation Token:", token)
// }

// package attestation

// import (
//     "context"
//     "encoding/json"
//     "errors"
//     "fmt"
//     "net/http"
//     "strings"
// )

// type GCPAttestation struct {
//     ProjectID       string
//     Location        string
//     AttestationURL  string
// }

// type GCPAttestationRequest struct {
//     Evidence string `json:"evidence"`
// }

// type GCPAttestationResponse struct {
//     Result string `json:"result"`
// }

// func NewGCPAttestation(projectID, location string) *GCPAttestation {
//     return &GCPAttestation{
//         ProjectID:      projectID,
//         Location:       location,
//         AttestationURL: fmt.Sprintf("https://confidentialcomputing.googleapis.com/v1/projects/%s/locations/%s/verifyAttestation", projectID, location),
//     }
// }

// // VerifyEvidence verifies the evidence using GCP's attestation API.
// func (g *GCPAttestation) VerifyEvidence(ctx context.Context, evidence string) (bool, error) {
//     requestBody, err := json.Marshal(GCPAttestationRequest{Evidence: evidence})
//     if err != nil {
//         return false, err
//     }

//     req, err := http.NewRequestWithContext(ctx, "POST", g.AttestationURL, strings.NewReader(string(requestBody)))
//     if err != nil {
//         return false, err
//     }
//     req.Header.Set("Content-Type", "application/json")

//     resp, err := http.DefaultClient.Do(req)
//     if err != nil {
//         return false, err
//     }
//     defer resp.Body.Close()

//     if resp.StatusCode != http.StatusOK {
//         return false, fmt.Errorf("attestation service returned status %d", resp.StatusCode)
//     }

//     var attResp GCPAttestationResponse
//     if err := json.NewDecoder(resp.Body).Decode(&attResp); err != nil {
//         return false, err
//     }

//     return attResp.Result == "VERIFIED", nil
// }
