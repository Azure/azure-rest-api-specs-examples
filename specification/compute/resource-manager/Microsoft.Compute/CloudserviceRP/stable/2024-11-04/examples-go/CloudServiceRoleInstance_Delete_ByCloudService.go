package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/250861bb6a886b75255edfa0aa5ee2dd0d6e7a11/specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2024-11-04/examples/CloudServiceRoleInstance_Delete_ByCloudService.json
func ExampleCloudServicesClient_BeginDeleteInstances() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudServicesClient().BeginDeleteInstances(ctx, "ConstosoRG", "{cs-name}", &armcompute.CloudServicesClientBeginDeleteInstancesOptions{Parameters: &armcompute.RoleInstances{
		RoleInstances: []*string{
			to.Ptr("ContosoFrontend_IN_0"),
			to.Ptr("ContosoBackend_IN_1")},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
