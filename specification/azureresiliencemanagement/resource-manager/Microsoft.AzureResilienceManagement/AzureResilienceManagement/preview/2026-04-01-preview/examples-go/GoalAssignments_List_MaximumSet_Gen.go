package armresiliencemanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resiliencemanagement/armresiliencemanagement"
)

// Generated from example definition: 2026-04-01-preview/GoalAssignments_List_MaximumSet_Gen.json
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
	pager := clientFactory.NewGoalAssignmentsClient().NewListPager("zldmpkvqzifygkqau", &armresiliencemanagement.GoalAssignmentsClientListOptions{
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
		// 					GoalTemplateID: to.Ptr("/providers/Microsoft.AzureResilienceManagement/goaltemplates/gt1"),
		// 					GoalAssignmentType: to.Ptr(armresiliencemanagement.GoalAssignmentTypeResiliency),
		// 					ServiceLevelResources: []*armresiliencemanagement.ServiceLevelResource{
		// 						{
		// 							ServiceLevelIndicatorResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/MyResourceGroup/providers/Microsoft.Compute/virtualMachines/MyVirtualMachine"),
		// 							ServiceLevelObjectiveResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/MyResourceGroup/providers/Microsoft.Compute/virtualMachines/MyVirtualMachine"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armresiliencemanagement.ProvisioningStateSucceeded),
		// 				},
		// 				ID: to.Ptr("/providers/Microsoft.AzureResilienceManagement/goalAssignments/ga1"),
		// 				Name: to.Ptr("ga1"),
		// 				Type: to.Ptr("Microsoft.AzureResilienceManagement/goalAssignments"),
		// 				SystemData: &armresiliencemanagement.SystemData{
		// 					CreatedBy: to.Ptr("dvnfxbuyqhvivfjddjccdtlwajfht"),
		// 					CreatedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-06T15:03:42.796Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("lndhhaimomorael"),
		// 					LastModifiedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-06T15:03:42.797Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/aoswipdy"),
		// 	},
		// }
	}
}
