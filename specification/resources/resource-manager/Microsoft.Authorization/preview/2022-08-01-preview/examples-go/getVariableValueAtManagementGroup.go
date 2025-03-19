package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/219b2e3ef270f18149774eb2793b48baacde982f/specification/resources/resource-manager/Microsoft.Authorization/preview/2022-08-01-preview/examples/getVariableValueAtManagementGroup.json
func ExampleVariableValuesClient_GetAtManagementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVariableValuesClient().GetAtManagementGroup(ctx, "DevOrg", "DemoTestVariable", "TestValue", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VariableValue = armpolicy.VariableValue{
	// 	Name: to.Ptr("TestValue"),
	// 	Type: to.Ptr("Microsoft.Authorization/variables/values"),
	// 	ID: to.Ptr("/providers/Microsoft.Management/managementGroups/DevOrg/providers/Microsoft.Authorization/variables/DemoTestVariable/values/TestValue"),
	// 	Properties: &armpolicy.VariableValueProperties{
	// 		Values: []*armpolicy.VariableValueColumnValue{
	// 			{
	// 				ColumnName: to.Ptr("StringColumn"),
	// 				ColumnValue: "SampleValue",
	// 			},
	// 			{
	// 				ColumnName: to.Ptr("IntegerColumn"),
	// 				ColumnValue: float64(10),
	// 		}},
	// 	},
	// 	SystemData: &armpolicy.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-01T01:01:01.107Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armpolicy.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-01T01:01:01.107Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armpolicy.CreatedByTypeUser),
	// 	},
	// }
}
