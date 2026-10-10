package armresiliencemanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resiliencemanagement/armresiliencemanagement"
)

// Generated from example definition: 2026-10-01/GoalAssignments_List_MaximumSet_Gen.json
func ExampleGoalAssignmentsClient_NewListPager_goalAssignmentsListMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresiliencemanagement.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGoalAssignmentsClient().NewListPager("production-sg", &armresiliencemanagement.GoalAssignmentsClientListOptions{
		SkipToken: to.Ptr("xntbyoswztnmvitj"),
		Top:       to.Ptr[int32](69)})
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
		// page = armresiliencemanagement.GoalAssignmentsClientListResponse{
		// 	GoalAssignmentListResult: armresiliencemanagement.GoalAssignmentListResult{
		// 		Value: []*armresiliencemanagement.GoalAssignment{
		// 			{
		// 				Properties: &armresiliencemanagement.GoalAssignmentProperties{
		// 					ServiceLevelResources: []*armresiliencemanagement.ServiceLevelResource{
		// 						{
		// 							ServiceLevelIndicatorResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/MyResourceGroup/providers/Microsoft.Compute/virtualMachines/MyVirtualMachine"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armresiliencemanagement.ProvisioningStateSucceeded),
		// 					RequireZonalResiliency: to.Ptr(true),
		// 				},
		// 				ID: to.Ptr("/providers/Microsoft.Management/serviceGroups/production-sg/providers/Microsoft.AzureResilienceManagement/goalAssignments/zonal-resiliency-goal"),
		// 				Name: to.Ptr("zonal-resiliency-goal"),
		// 				Type: to.Ptr("Microsoft.AzureResilienceManagement/goalAssignments"),
		// 				SystemData: &armresiliencemanagement.SystemData{
		// 					CreatedBy: to.Ptr("admin@contoso.com"),
		// 					CreatedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2025, time.February, 6, 15, 3, 42, 796000000, time.UTC)),
		// 					LastModifiedBy: to.Ptr("admin@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2025, time.February, 6, 15, 3, 42, 797000000, time.UTC)),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/providers/Microsoft.Management/serviceGroups/production-sg/providers/Microsoft.AzureResilienceManagement/goalAssignments?api-version=2026-10-01&$skipToken=eyJuZXh0UGFnZSI6Mn0%3D"),
		// 	},
		// }
	}
}
