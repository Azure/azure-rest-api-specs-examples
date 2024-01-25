package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/41e4538ed7bb3ceac3c1322c9455a0812ed110ac/specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/CloudServiceRole_Get.json
func ExampleCloudServiceRolesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCloudServiceRolesClient().Get(ctx, "{role-name}", "ConstosoRG", "{cs-name}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CloudServiceRole = armcompute.CloudServiceRole{
	// 	Name: to.Ptr("{role-name}"),
	// 	Type: to.Ptr("Microsoft.Compute/cloudServices/roles"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Compute/cloudServices/{cs-name}/roles/{role-name}"),
	// 	Location: to.Ptr("eastus2euap"),
	// 	Properties: &armcompute.CloudServiceRoleProperties{
	// 		UniqueID: to.Ptr("b03bc269-766b-4921-b91a-7dffaae4d03b:{role-name}"),
	// 	},
	// 	SKU: &armcompute.CloudServiceRoleSKU{
	// 		Name: to.Ptr("Standard_D1_v2"),
	// 		Capacity: to.Ptr[int64](2),
	// 		Tier: to.Ptr("Standard"),
	// 	},
	// }
}
