package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-03-01-preview/UpdateContents_Get_MaximumSet_Gen.json
func ExampleUpdateContentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("2886575D-173A-44A0-80E2-7DBA57F18B46", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewUpdateContentsClient().Get(ctx, "westus2", "12.2510.0.1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armazurestackhci.UpdateContentsClientGetResponse{
	// 	UpdateContent: &armazurestackhci.UpdateContent{
	// 		Properties: &armazurestackhci.UpdateContentProperties{
	// 			UpdatePayloads: []*armazurestackhci.ContentPayload{
	// 				{
	// 					URL: to.Ptr("https://microsoft.com/a"),
	// 					Hash: to.Ptr("F51FFEEB50A5C51B930F107132DDFA389DC8983313438764204A6F70F4BEFC60"),
	// 					HashAlgorithm: to.Ptr("SHA256"),
	// 					Identifier: to.Ptr("SolutionUpdate_zip"),
	// 					PackageSizeInBytes: to.Ptr("9367"),
	// 					Group: to.Ptr("Update"),
	// 					FileName: to.Ptr("SolutionUpdate.12.2510.1002.84.zip"),
	// 				},
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/b8d594e5-51f3-4c11-9c54-a7771b81c712/providers/Microsoft.AzureStackHCI/updateContents/12.2510.0.1"),
	// 		Name: to.Ptr("12.2510.0.1"),
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/UpdateContents"),
	// 	},
	// }
}
