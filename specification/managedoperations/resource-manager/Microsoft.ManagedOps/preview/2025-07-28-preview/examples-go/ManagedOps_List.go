package armmanagedops_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managedops/armmanagedops"
)

// Generated from example definition: 2025-07-28-preview/ManagedOps_List.json
func ExampleClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedops.NewClientFactory("11809CA1-E126-4017-945E-AA795CD5C5A9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListPager(nil)
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
		// page = armmanagedops.ClientListResponse{
		// 	ManagedOpListResult: armmanagedops.ManagedOpListResult{
		// 		Value: []*armmanagedops.ManagedOp{
		// 			{
		// 				ID: to.Ptr("/subscriptions/11809CA1-E126-4017-945E-AA795CD5C5A9/providers/Microsoft.ManagedOps/managedOps/default"),
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Microsoft.ManagedOps/managedOps"),
		// 				SystemData: &armmanagedops.SystemData{
		// 					CreatedBy: to.Ptr("myUser"),
		// 					CreatedByType: to.Ptr(armmanagedops.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-19T00:28:48.610Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("myUser"),
		// 					LastModifiedByType: to.Ptr(armmanagedops.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-19T00:28:48.610Z"); return t}()),
		// 				},
		// 				Properties: &armmanagedops.Properties{
		// 					SKU: &armmanagedops.SKU{
		// 						Name: to.Ptr("ManagedOps"),
		// 						Tier: to.Ptr("Essential"),
		// 					},
		// 					DesiredConfiguration: &armmanagedops.DesiredConfiguration{
		// 						ChangeTrackingAndInventory: &armmanagedops.ChangeTrackingConfiguration{
		// 							LogAnalyticsWorkspaceID: to.Ptr("/subscriptions/11809CA1-E126-4017-945E-AA795CD5C5A9/resourceGroups/myResourceGroup/providers/Microsoft.OperationalInsights/workspaces/00000000-0000-0000-0000-000000000000-Default"),
		// 						},
		// 						AzureMonitorInsights: &armmanagedops.AzureMonitorConfiguration{
		// 							AzureMonitorWorkspaceID: to.Ptr("/subscriptions/11809CA1-E126-4017-945E-AA795CD5C5A9/resourceGroups/myResourceGroup/providers/Microsoft.Monitor/accounts/example"),
		// 						},
		// 						UserAssignedManagedIdentityID: to.Ptr("/subscriptions/11809CA1-E126-4017-945E-AA795CD5C5A9/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myManagedIdentity"),
		// 						DefenderForServers: to.Ptr(armmanagedops.DesiredConfigurationDefenderForServersEnable),
		// 						DefenderCspm: to.Ptr(armmanagedops.DesiredConfigurationDefenderForServersEnable),
		// 					},
		// 					ProvisioningState: to.Ptr(armmanagedops.ProvisioningStateProvisioning),
		// 					PolicyAssignmentProperties: &armmanagedops.PolicyAssignmentProperties{
		// 						PolicyInitiativeAssignmentID: to.Ptr("/subscriptions/aa4c0cba-5c0a-4751-a5e3-ea28e50afeeb/providers/microsoft.authorization/policyassignments/test-edition-policy-2"),
		// 					},
		// 					Services: &armmanagedops.ServiceInformation{
		// 						ChangeTrackingAndInventory: &armmanagedops.ChangeTrackingInformation{
		// 							DcrID: to.Ptr("/subscriptions/11809CA1-E126-4017-945E-AA795CD5C5A9/resourceGroups/xingftestrg/providers/Microsoft.Insights/dataCollectionRules/ct-dcr131878407"),
		// 							EnablementStatus: to.Ptr(armmanagedops.ChangeTrackingInformationEnablementStatusEnabled),
		// 						},
		// 						AzureMonitorInsights: &armmanagedops.AzureMonitorInformation{
		// 							DcrID: to.Ptr("/subscriptions/11809CA1-E126-4017-945E-AA795CD5C5A9/resourceGroups/xingftestrg/providers/Microsoft.Insights/dataCollectionRules/ami-dcr131878407"),
		// 							EnablementStatus: to.Ptr(armmanagedops.ChangeTrackingInformationEnablementStatusEnabled),
		// 						},
		// 						AzureUpdateManager: &armmanagedops.UpdateManagerInformation{
		// 							EnablementStatus: to.Ptr(armmanagedops.ChangeTrackingInformationEnablementStatusEnabled),
		// 						},
		// 						AzurePolicyAndMachineConfiguration: &armmanagedops.GuestConfigurationInformation{
		// 							EnablementStatus: to.Ptr(armmanagedops.ChangeTrackingInformationEnablementStatusEnabled),
		// 						},
		// 						DefenderForServers: &armmanagedops.DefenderForServersInformation{
		// 							EnablementStatus: to.Ptr(armmanagedops.ChangeTrackingInformationEnablementStatusEnabled),
		// 						},
		// 						DefenderCspm: &armmanagedops.DefenderCspmInformation{
		// 							EnablementStatus: to.Ptr(armmanagedops.ChangeTrackingInformationEnablementStatusEnabled),
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
