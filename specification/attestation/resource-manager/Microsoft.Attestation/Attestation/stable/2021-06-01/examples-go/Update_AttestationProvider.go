package armattestation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/attestation/armattestation/v2"
)

// Generated from example definition: 2021-06-01/Update_AttestationProvider.json
func ExampleProvidersClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armattestation.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProvidersClient().Update(ctx, "MyResourceGroup", "myattestationprovider", armattestation.ServicePatchParams{
		Properties: &armattestation.ServicePatchSpecificParams{
			PublicNetworkAccess:          to.Ptr(armattestation.PublicNetworkAccessTypeDisabled),
			TpmAttestationAuthentication: to.Ptr(armattestation.TpmAttestationAuthenticationTypeDisabled),
		},
		Tags: map[string]*string{
			"Property1": to.Ptr("Value1"),
			"Property2": to.Ptr("Value2"),
			"Property3": to.Ptr("Value3"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armattestation.ProvidersClientUpdateResponse{
	// 	Provider: &armattestation.Provider{
	// 		Name: to.Ptr("myattestationprovider"),
	// 		Type: to.Ptr("Microsoft.Attestation/attestationProviders"),
	// 		ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MyResourceGroup/providers/Microsoft.Attestation/attestationProviders/myattestationprovider"),
	// 		Location: to.Ptr("East US"),
	// 		Properties: &armattestation.StatusResult{
	// 			AttestURI: to.Ptr("https://superservice.attestation.azure.net"),
	// 			PublicNetworkAccess: to.Ptr(armattestation.PublicNetworkAccessTypeDisabled),
	// 			Status: to.Ptr(armattestation.AttestationServiceStatusReady),
	// 			TpmAttestationAuthentication: to.Ptr(armattestation.TpmAttestationAuthenticationTypeDisabled),
	// 			TrustModel: to.Ptr("Isolated"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"Property1": to.Ptr("Value1"),
	// 			"Property2": to.Ptr("Value2"),
	// 			"Property3": to.Ptr("Value3"),
	// 		},
	// 	},
	// }
}
