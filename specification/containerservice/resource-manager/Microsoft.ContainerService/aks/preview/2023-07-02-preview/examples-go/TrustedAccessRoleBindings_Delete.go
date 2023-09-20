package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-07-02-preview/examples/TrustedAccessRoleBindings_Delete.json
func ExampleTrustedAccessRoleBindingsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTrustedAccessRoleBindingsClient().Delete(ctx, "rg1", "clustername1", "binding1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
