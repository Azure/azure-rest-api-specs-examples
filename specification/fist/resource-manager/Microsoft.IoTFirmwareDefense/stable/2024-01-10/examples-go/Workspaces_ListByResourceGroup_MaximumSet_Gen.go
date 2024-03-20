package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/Workspaces_ListByResourceGroup_MaximumSet_Gen.json
func ExampleWorkspacesClient_NewListByResourceGroupPager_workspacesListByResourceGroupMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspacesClient().NewListByResourceGroupPager("rgworkspaces", nil)
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
		// page.WorkspaceList = armiotfirmwaredefense.WorkspaceList{
		// 	Value: []*armiotfirmwaredefense.Workspace{
		// 		{
		// 			Name: to.Ptr("tbrqhnzpsatbrnhtj"),
		// 			Type: to.Ptr("angrpzpxmuppzqpzpljasjirao"),
		// 			ID: to.Ptr("/subscriptions/blah/resourceGroups/blah/providers/blah/workspaces/blah"),
		// 			SystemData: &armiotfirmwaredefense.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-19T05:35:25.545Z"); return t}()),
		// 				CreatedBy: to.Ptr("kmgstbzxtl"),
		// 				CreatedByType: to.Ptr(armiotfirmwaredefense.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-19T05:35:25.545Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("rfnhyhbyqzxnumjmjhwkhduztidk"),
		// 				LastModifiedByType: to.Ptr(armiotfirmwaredefense.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("jjwbseilitjgdrhbvvkwviqj"),
		// 			Tags: map[string]*string{
		// 				"key450": to.Ptr("rzqqumbpfsbibnpirsm"),
		// 			},
		// 			Properties: &armiotfirmwaredefense.WorkspaceProperties{
		// 				ProvisioningState: to.Ptr(armiotfirmwaredefense.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
