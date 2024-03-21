package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-07-01/examples/BackupPolicies_Create.json
func ExampleBackupPoliciesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupPoliciesClient().BeginCreate(ctx, "myRG", "account1", "backupPolicyName", armnetapp.BackupPolicy{
		Location: to.Ptr("westus"),
		Properties: &armnetapp.BackupPolicyProperties{
			DailyBackupsToKeep:   to.Ptr[int32](10),
			Enabled:              to.Ptr(true),
			MonthlyBackupsToKeep: to.Ptr[int32](10),
			WeeklyBackupsToKeep:  to.Ptr[int32](10),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BackupPolicy = armnetapp.BackupPolicy{
	// 	Name: to.Ptr("account1/backupPolicyName"),
	// 	Type: to.Ptr("Microsoft.NetApp/netAppAccounts/backupPolicies"),
	// 	ID: to.Ptr("/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/backupPolocies/backupPolicyName"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armnetapp.BackupPolicyProperties{
	// 		DailyBackupsToKeep: to.Ptr[int32](10),
	// 		Enabled: to.Ptr(true),
	// 		MonthlyBackupsToKeep: to.Ptr[int32](10),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		WeeklyBackupsToKeep: to.Ptr[int32](10),
	// 	},
	// }
}
