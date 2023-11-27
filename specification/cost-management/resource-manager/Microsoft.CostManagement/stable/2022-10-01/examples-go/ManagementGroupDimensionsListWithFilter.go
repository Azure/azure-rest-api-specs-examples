package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ManagementGroupDimensionsListWithFilter.json
func ExampleDimensionsClient_NewListPager_managementGroupDimensionsListWithFilterLegacy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDimensionsClient().NewListPager("providers/Microsoft.Management/managementGroups/MyMgId", &armcostmanagement.DimensionsClientListOptions{Filter: to.Ptr("properties/category eq 'resourceId'"),
		Expand:    to.Ptr("properties/data"),
		Skiptoken: nil,
		Top:       to.Ptr[int32](5),
	})
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
		// page.DimensionsListResult = armcostmanagement.DimensionsListResult{
		// 	Value: []*armcostmanagement.Dimension{
		// 		{
		// 			Name: to.Ptr("dimensions_ResourceId_2018-05-01_2018-05-31_5"),
		// 			Type: to.Ptr("microsoft.CostManagement/dimensions"),
		// 			ID: to.Ptr("providers/Microsoft.Management/managementGroups/MyMgId/providers/microsoft.CostManagement/dimensions_ResourceId_2018-05-01_2018-05-31_5"),
		// 			Properties: &armcostmanagement.DimensionProperties{
		// 				Description: to.Ptr("Resource Id"),
		// 				Category: to.Ptr("ResourceId"),
		// 				Data: []*string{
		// 					to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/system.orlando/providers/microsoft.storage/storageaccounts/urphealthaccount"),
		// 					to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/system.orlando/providers/microsoft.storage/storageaccounts/srphytenaccount"),
		// 					to.Ptr("/subscriptions/67e24f6b-1ec2-4c90-993a-dc2d25b00b6c/resourcegroups/defaultresourcegroup-eus/providers/microsoft.operationalinsights/workspaces/defaultworkspace-67e24f6b-1ec2-4c90-993a-dc2d25b00b6c-eus"),
		// 					to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg-sql-ha/providers/microsoft.compute/virtualmachines/sql-4qqp1"),
		// 					to.Ptr("/subscriptions/a98d6dc5-eb8f-46cf-8938-f1fb08f03706/resourcegroups/databricks-rg-testwsp-xijmsdubneexm/providers/microsoft.compute/disks/488cdb42bf74474a98075415be3f806c-containerrootvolume")},
		// 					FilterEnabled: to.Ptr(true),
		// 					GroupingEnabled: to.Ptr(true),
		// 					Total: to.Ptr[int32](1409),
		// 					UsageEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-31T07:00:00.000Z"); return t}()),
		// 					UsageStart: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-01T07:00:00.000Z"); return t}()),
		// 				},
		// 		}},
		// 	}
	}
}
