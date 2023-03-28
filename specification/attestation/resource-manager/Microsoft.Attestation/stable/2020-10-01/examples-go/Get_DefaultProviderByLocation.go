package armattestation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/attestation/armattestation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/Get_DefaultProviderByLocation.json
func ExampleProvidersClient_GetDefaultByLocation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armattestation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProvidersClient().GetDefaultByLocation(ctx, "Central US", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Provider = armattestation.Provider{
	// 	Name: to.Ptr("sharedcus"),
	// 	Type: to.Ptr("Microsoft.Attestation/attestationProviders"),
	// 	ID: to.Ptr("providers/Microsoft.Attestation/attestationProviders/sharedcus"),
	// 	Location: to.Ptr("Central US"),
	// 	Properties: &armattestation.StatusResult{
	// 		AttestURI: to.Ptr("https://sharedcus.cus.attest.azure.net"),
	// 		Status: to.Ptr(armattestation.AttestationServiceStatusReady),
	// 		TrustModel: to.Ptr("AAD"),
	// 	},
	// }
}
