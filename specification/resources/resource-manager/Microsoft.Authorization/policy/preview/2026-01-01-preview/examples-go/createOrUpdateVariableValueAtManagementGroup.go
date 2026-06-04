package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy/v2"
)

// Generated from example definition: 2026-01-01-preview/createOrUpdateVariableValueAtManagementGroup.json
func ExampleVariableValuesClient_CreateOrUpdateAtManagementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVariableValuesClient().CreateOrUpdateAtManagementGroup(ctx, "DevOrg", "DemoTestVariable", "TestValue", armpolicy.VariableValue{
		Properties: &armpolicy.VariableValueProperties{
			Values: []*armpolicy.VariableValueColumnValue{
				{
					ColumnName:  to.Ptr("StringColumn"),
					ColumnValue: "SampleValue",
				},
				{
					ColumnName:  to.Ptr("IntegerColumn"),
					ColumnValue: 10,
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpolicy.VariableValuesClientCreateOrUpdateAtManagementGroupResponse{
	// 	VariableValue: armpolicy.VariableValue{
	// 		Properties: &armpolicy.VariableValueProperties{
	// 			Values: []*armpolicy.VariableValueColumnValue{
	// 				{
	// 					ColumnName: to.Ptr("StringColumn"),
	// 					ColumnValue: "SampleValue",
	// 				},
	// 				{
	// 					ColumnName: to.Ptr("IntegerColumn"),
	// 					ColumnValue: 10,
	// 				},
	// 			},
	// 		},
	// 		SystemData: &armpolicy.SystemData{
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armpolicy.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-01T01:01:01.1075056Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armpolicy.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-01T02:01:01.1075056Z"); return t}()),
	// 		},
	// 		ID: to.Ptr("/providers/Microsoft.Management/managementGroups/DevOrg/providers/Microsoft.Authorization/variables/DemoTestVariable/values/TestValue"),
	// 		Type: to.Ptr("Microsoft.Authorization/variables/values"),
	// 		Name: to.Ptr("TestValue"),
	// 	},
	// }
}
