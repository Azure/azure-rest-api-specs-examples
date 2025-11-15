package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6bbb6daca7175b2cab69b20b2dd01094e3c6a534/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/ProjectCapabilityHost/get.json
func ExampleProjectCapabilityHostsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProjectCapabilityHostsClient().Get(ctx, "test-rg", "account-1", "project-1", "capabilityHostName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProjectCapabilityHost = armcognitiveservices.ProjectCapabilityHost{
	// 	Name: to.Ptr("capabilityHostName"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/accounts/projects/capabilityHosts"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/test-rg/providers/Microsoft.CognitiveServices/accounts/account-1/projects/project-1/capabilityHosts/capabilityHostName"),
	// 	Properties: &armcognitiveservices.ProjectCapabilityHostProperties{
	// 		AiServicesConnections: []*string{
	// 			to.Ptr("aoai_connection")},
	// 			ProvisioningState: to.Ptr(armcognitiveservices.CapabilityHostProvisioningStateSucceeded),
	// 			StorageConnections: []*string{
	// 				to.Ptr("blob_connection")},
	// 				ThreadStorageConnections: []*string{
	// 					to.Ptr("aca_connection")},
	// 					VectorStoreConnections: []*string{
	// 						to.Ptr("acs_connection")},
	// 					},
	// 				}
}
