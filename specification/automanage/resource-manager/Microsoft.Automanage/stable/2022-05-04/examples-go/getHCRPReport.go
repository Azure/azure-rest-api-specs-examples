package armautomanage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automanage/armautomanage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/2dcad6d6e9a96882eb6d317e7500a94be007a9c6/specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/getHCRPReport.json
func ExampleHCRPReportsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomanage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewHCRPReportsClient().Get(ctx, "myResourceGroupName", "myMachineName", "default", "b4e9ee6b-1717-4ff0-a8d2-e6d72c33d5f4", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Report = armautomanage.Report{
	// 	Name: to.Ptr("b4e9ee6b-1717-4ff0-a8d2-e6d72c33d5f4"),
	// 	Type: to.Ptr("Microsoft.Automanage/configurationProfileAssignments/reports"),
	// 	ID: to.Ptr("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.HybridCompute/machines/myClustereName/providers/Microsoft.Automanage/configurationProfileAssignments/default/reports/b4e9ee6b-1717-4ff0-a8d2-e6d72c33d5f4"),
	// 	Properties: &armautomanage.AssignmentReportProperties{
	// 		Type: to.Ptr("Consistency"),
	// 		ConfigurationProfile: to.Ptr("anyConfigurationProfile"),
	// 		Duration: to.Ptr("PT15M32S"),
	// 		EndTime: to.Ptr("2021-03-31T22:17:42Z"),
	// 		LastModifiedTime: to.Ptr("2021-03-31T22:32:42Z"),
	// 		ReportFormatVersion: to.Ptr("1.0"),
	// 		Resources: []*armautomanage.ReportResource{
	// 			{
	// 				Name: to.Ptr("myResourceGroupName"),
	// 				Type: to.Ptr("Microsoft.Resources/resourceGroups"),
	// 				ID: to.Ptr("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName"),
	// 				Status: to.Ptr("Conformant"),
	// 		}},
	// 		StartTime: to.Ptr("2021-03-31T22:13:06Z"),
	// 		Status: to.Ptr("Conformant"),
	// 	},
	// 	SystemData: &armautomanage.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.107Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1@outlook.com"),
	// 		CreatedByType: to.Ptr(armautomanage.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-04T02:03:01.197Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2@outlook.com"),
	// 		LastModifiedByType: to.Ptr(armautomanage.CreatedByTypeUser),
	// 	},
	// }
}
