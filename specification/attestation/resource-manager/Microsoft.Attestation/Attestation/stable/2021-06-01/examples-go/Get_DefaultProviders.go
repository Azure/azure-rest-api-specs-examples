package armattestation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/attestation/armattestation/v2"
)

// Generated from example definition: 2021-06-01/Get_DefaultProviders.json
func ExampleProvidersClient_ListDefault() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armattestation.NewClientFactory("6c96b33e-f5b8-40a6-9011-5cb1c58b0915", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProvidersClient().ListDefault(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armattestation.ProvidersClientListDefaultResponse{
	// 	ProviderListResult: &armattestation.ProviderListResult{
	// 		Value: []*armattestation.Provider{
	// 			{
	// 				Name: to.Ptr("sharedcus"),
	// 				Type: to.Ptr("Microsoft.Attestation/attestationProviders"),
	// 				ID: to.Ptr("providers/Microsoft.Attestation/attestationProviders/sharedcus"),
	// 				Location: to.Ptr("Central US"),
	// 				Properties: &armattestation.StatusResult{
	// 					AttestURI: to.Ptr("https://sharedcus.cus.attest.azure.net"),
	// 					PublicNetworkAccess: to.Ptr(armattestation.PublicNetworkAccessTypeEnabled),
	// 					Status: to.Ptr(armattestation.AttestationServiceStatusReady),
	// 					TpmAttestationAuthentication: to.Ptr(armattestation.TpmAttestationAuthenticationTypeEnabled),
	// 					TrustModel: to.Ptr("AAD"),
	// 				},
	// 			},
	// 			{
	// 				Name: to.Ptr("shareduks"),
	// 				Type: to.Ptr("Microsoft.Attestation/attestationProviders"),
	// 				ID: to.Ptr("providers/Microsoft.Attestation/attestationProviders/shareduks"),
	// 				Location: to.Ptr("UK South"),
	// 				Properties: &armattestation.StatusResult{
	// 					AttestURI: to.Ptr("https://shareduks.uks.attest.azure.net"),
	// 					PublicNetworkAccess: to.Ptr(armattestation.PublicNetworkAccessTypeEnabled),
	// 					Status: to.Ptr(armattestation.AttestationServiceStatusReady),
	// 					TpmAttestationAuthentication: to.Ptr(armattestation.TpmAttestationAuthenticationTypeEnabled),
	// 					TrustModel: to.Ptr("AAD"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
