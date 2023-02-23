package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListWorkspaceManagedSqlServerSecurityAlertPolicies.json
func ExampleWorkspaceManagedSQLServerSecurityAlertPolicyClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewWorkspaceManagedSQLServerSecurityAlertPolicyClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("wsg-7398", "testWorkspace", nil)
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
		// page.ServerSecurityAlertPolicyListResult = armsynapse.ServerSecurityAlertPolicyListResult{
		// 	Value: []*armsynapse.ServerSecurityAlertPolicy{
		// 		{
		// 			Name: to.Ptr("Default"),
		// 			Type: to.Ptr("Microsoft.Synapse/workspaces/securityAlertPolicies"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/wsg-7398/providers/Microsoft.Synapse/workspaces/testWorkspace/securityAlertPolicies"),
		// 			Properties: &armsynapse.ServerSecurityAlertPolicyProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-23T04:52:49.52Z"); return t}()),
		// 				DisabledAlerts: []*string{
		// 					to.Ptr("Access_Anomaly")},
		// 					EmailAccountAdmins: to.Ptr(true),
		// 					EmailAddresses: []*string{
		// 						to.Ptr("test@microsoft.com;user@microsoft.com")},
		// 						RetentionDays: to.Ptr[int32](0),
		// 						State: to.Ptr(armsynapse.SecurityAlertPolicyStateDisabled),
		// 						StorageAccountAccessKey: to.Ptr(""),
		// 						StorageEndpoint: to.Ptr("https://mystorage.blob.core.windows.net"),
		// 					},
		// 			}},
		// 		}
	}
}
