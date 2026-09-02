package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy/v2"
)

// Generated from example definition: 2026-01-01-preview/listVariableValuesForManagementGroup.json
func ExampleVariableValuesClient_NewListForManagementGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVariableValuesClient().NewListForManagementGroupPager("DevOrg", "DemoTestVariable", nil)
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
		// page = armpolicy.VariableValuesClientListForManagementGroupResponse{
		// 	VariableValueListResult: armpolicy.VariableValueListResult{
		// 		Value: []*armpolicy.VariableValue{
		// 			{
		// 				Properties: &armpolicy.VariableValueProperties{
		// 					Values: []*armpolicy.VariableValueColumnValue{
		// 						{
		// 							ColumnName: to.Ptr("StringColumn"),
		// 							ColumnValue: "SampleValue",
		// 						},
		// 						{
		// 							ColumnName: to.Ptr("IntegerColumn"),
		// 							ColumnValue: 10,
		// 						},
		// 					},
		// 				},
		// 				SystemData: &armpolicy.SystemData{
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2022, time.July, 1, 1, 1, 1, 107505600, time.UTC)),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2022, time.August, 1, 1, 1, 1, 107505600, time.UTC)),
		// 				},
		// 				ID: to.Ptr("/providers/Microsoft.Management/managementGroups/DevOrg/providers/Microsoft.Authorization/variables/DemoTestVariable/values/TestValue"),
		// 				Type: to.Ptr("Microsoft.Authorization/variables/values"),
		// 				Name: to.Ptr("TestValue"),
		// 			},
		// 			{
		// 				Properties: &armpolicy.VariableValueProperties{
		// 					Values: []*armpolicy.VariableValueColumnValue{
		// 						{
		// 							ColumnName: to.Ptr("NullColumnName"),
		// 							ColumnValue: nil,
		// 						},
		// 					},
		// 				},
		// 				SystemData: &armpolicy.SystemData{
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2020, time.July, 1, 1, 1, 1, 107505600, time.UTC)),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2020, time.July, 1, 2, 1, 1, 107505600, time.UTC)),
		// 				},
		// 				ID: to.Ptr("/providers/Microsoft.Management/managementGroups/DevOrg/providers/Microsoft.Authorization/variables/DemoTestVariable/values/NullableTestValue"),
		// 				Type: to.Ptr("Microsoft.Authorization/variables/values"),
		// 				Name: to.Ptr("NullableTestValue"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
