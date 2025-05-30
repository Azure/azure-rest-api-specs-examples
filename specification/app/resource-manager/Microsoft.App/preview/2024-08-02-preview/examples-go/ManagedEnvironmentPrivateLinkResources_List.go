package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/ManagedEnvironmentPrivateLinkResources_List.json
func ExampleManagedEnvironmentPrivateLinkResourcesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedEnvironmentPrivateLinkResourcesClient().NewListPager("examplerg", "managedEnv", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.PrivateLinkResourceListResult = armappcontainers.PrivateLinkResourceListResult{
		// 	Value: []*armappcontainers.PrivateLinkResource{
		// 		{
		// 			Name: to.Ptr("managedEnvironments"),
		// 			Type: to.Ptr("Microsoft.App/managedEnvironments/privateLinkResources"),
		// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/managedEnv/privateLinkResources/managedEnvironments"),
		// 			Properties: &armappcontainers.PrivateLinkResourceProperties{
		// 				GroupID: to.Ptr("managedEnvironments"),
		// 				RequiredMembers: []*string{
		// 					to.Ptr("managedEnvironments")},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("privatelink.northcentralusstage.azurecontainerapps.io")},
		// 					},
		// 			}},
		// 		}
	}
}
