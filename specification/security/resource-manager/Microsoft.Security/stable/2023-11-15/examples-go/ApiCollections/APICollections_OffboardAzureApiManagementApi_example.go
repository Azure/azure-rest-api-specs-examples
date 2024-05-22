package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ac34f238dd6b9071f486b57e9f9f1a0c43ec6f6/specification/security/resource-manager/Microsoft.Security/stable/2023-11-15/examples/ApiCollections/APICollections_OffboardAzureApiManagementApi_example.json
func ExampleAPICollectionsClient_OffboardAzureAPIManagementAPI() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAPICollectionsClient().OffboardAzureAPIManagementAPI(ctx, "rg1", "apimService1", "echo-api", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
