package armautomanage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automanage/armautomanage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/2dcad6d6e9a96882eb6d317e7500a94be007a9c6/specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/deleteConfigurationProfileHCIAssignment.json
func ExampleConfigurationProfileHCIAssignmentsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomanage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewConfigurationProfileHCIAssignmentsClient().Delete(ctx, "myResourceGroupName", "myClusterName", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
