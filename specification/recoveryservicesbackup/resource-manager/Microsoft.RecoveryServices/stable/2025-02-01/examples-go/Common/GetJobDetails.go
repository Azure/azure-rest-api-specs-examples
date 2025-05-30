package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/148c3b0b44f7789ced94859992493fafd0072f83/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/Common/GetJobDetails.json
func ExampleJobDetailsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobDetailsClient().Get(ctx, "NetSDKTestRsVault", "SwaggerTestRg", "00000000-0000-0000-0000-000000000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.JobResource = armrecoveryservicesbackup.JobResource{
	// 	Name: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupJobs"),
	// 	ID: to.Ptr("/Subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SwaggerTestRg/providers/Microsoft.RecoveryServices/vaults/NetSDKTestRsVault/backupJobs/00000000-0000-0000-0000-000000000000"),
	// 	Properties: &armrecoveryservicesbackup.AzureIaaSVMJob{
	// 		ActivityID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		BackupManagementType: to.Ptr(armrecoveryservicesbackup.BackupManagementTypeAzureIaasVM),
	// 		EntityFriendlyName: to.Ptr("testvm"),
	// 		JobType: to.Ptr("AzureIaaSVMJob"),
	// 		Operation: to.Ptr("Backup"),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-03T05:31:07.014Z"); return t}()),
	// 		Status: to.Ptr("InProgress"),
	// 		Duration: to.Ptr("PT9.8782791S"),
	// 		ExtendedInfo: &armrecoveryservicesbackup.AzureIaaSVMJobExtendedInfo{
	// 			PropertyBag: map[string]*string{
	// 				"VM Name": to.Ptr("testvm"),
	// 			},
	// 			TasksList: []*armrecoveryservicesbackup.AzureIaaSVMJobTaskDetails{
	// 				{
	// 					Duration: to.Ptr("PT0S"),
	// 					Status: to.Ptr("InProgress"),
	// 					TaskID: to.Ptr("Take Snapshot"),
	// 				},
	// 				{
	// 					Duration: to.Ptr("PT0S"),
	// 					Status: to.Ptr("NotStarted"),
	// 					TaskID: to.Ptr("Transfer data to vault"),
	// 			}},
	// 		},
	// 		VirtualMachineVersion: to.Ptr("Compute"),
	// 	},
	// }
}
