package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-03-01-preview/PlatformUpdates_Get_MaximumSet_Gen.json
func ExamplePlatformUpdatesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("b8d594e5-51f3-4c11-9c54-a7771b81c712", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPlatformUpdatesClient().Get(ctx, "westus2", "10.2408.0.1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armazurestackhci.PlatformUpdatesClientGetResponse{
	// 	PlatformUpdate: &armazurestackhci.PlatformUpdate{
	// 		Properties: &armazurestackhci.PlatformUpdateProperties{
	// 			PlatformUpdateDetails: []*armazurestackhci.PlatformUpdateDetails{
	// 				{
	// 					ValidatedSolutionRecipeVersion: to.Ptr("10.2408.0.1"),
	// 					PlatformVersion: to.Ptr("10.2408.0.1"),
	// 					PlatformPayloads: []*armazurestackhci.PlatformPayload{
	// 						{
	// 							PayloadHash: to.Ptr("sdfdsfdsfsadfdsafdsafawfdsafdsfa"),
	// 							PayloadURL: to.Ptr("https://microsoft.com/a"),
	// 						},
	// 					},
	// 				},
	// 				{
	// 					ValidatedSolutionRecipeVersion: to.Ptr("10.2408.0.2"),
	// 					PlatformVersion: to.Ptr("10.2408.0.2"),
	// 					PlatformPayloads: []*armazurestackhci.PlatformPayload{
	// 						{
	// 							PayloadHash: to.Ptr("sdfdsfdsfsadfdsafdsafawfdsafdsfa"),
	// 							PayloadURL: to.Ptr("https://microsoft.com/a"),
	// 							PayloadPackageSizeInBytes: to.Ptr("2048"),
	// 							PayloadIdentifier: to.Ptr("payload-identifier-1234"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/b8d594e5-51f3-4c11-9c54-a7771b81c712/providers/Microsoft.AzureStackHCI/platformUpdates/10.2408.0.1"),
	// 		Name: to.Ptr("10.2408.0.1"),
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/PlatformUpdates"),
	// 	},
	// }
}
