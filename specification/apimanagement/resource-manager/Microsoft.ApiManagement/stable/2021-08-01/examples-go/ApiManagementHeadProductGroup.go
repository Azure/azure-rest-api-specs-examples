package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadProductGroup.json
func ExampleProductGroupClient_CheckEntityExists() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapimanagement.NewProductGroupClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.CheckEntityExists(ctx,
		"rg1",
		"apimService1",
		"5931a75ae4bbd512a88c680b",
		"59306a29e4bbd510dc24e5f9",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
