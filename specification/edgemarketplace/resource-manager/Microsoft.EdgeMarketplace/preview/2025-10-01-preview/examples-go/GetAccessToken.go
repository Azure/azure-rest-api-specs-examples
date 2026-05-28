package armedgemarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgemarketplace/armedgemarketplace"
)

// Generated from example definition: 2025-10-01-preview/GetAccessToken.json
func ExampleOffersClient_GetAccessToken() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armedgemarketplace.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOffersClient().GetAccessToken(ctx, "subscriptions/4bed37fd-19a1-4d31-8b44-40267555bec5/resourceGroups/edgemarketplace-rg/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/edgemarketplace-demo", "0001-com-ubuntu-pro-jammy", armedgemarketplace.AccessTokenReadRequest{
		RequestID: to.Ptr("1.0.18062.1"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armedgemarketplace.OffersClientGetAccessTokenResponse{
	// 	DiskAccessToken: &armedgemarketplace.DiskAccessToken{
	// 		DiskID: to.Ptr("abcd"),
	// 		Status: to.Ptr("created"),
	// 		AccessToken: to.Ptr("hidden"),
	// 	},
	// }
}
