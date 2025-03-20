package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/219b2e3ef270f18149774eb2793b48baacde982f/specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/deletePolicyDefinitionVersionAtManagementGroup.json
func ExampleDefinitionVersionsClient_DeleteAtManagementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDefinitionVersionsClient().DeleteAtManagementGroup(ctx, "MyManagementGroup", "ResourceNaming", "1.2.1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
