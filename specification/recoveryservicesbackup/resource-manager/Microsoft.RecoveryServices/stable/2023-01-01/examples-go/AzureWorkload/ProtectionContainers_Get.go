package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3751704f5318f1175875c94b66af769db917f2d3/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-01-01/examples/AzureWorkload/ProtectionContainers_Get.json
func ExampleProtectionContainersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicesbackup.NewProtectionContainersClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "testVault", "testRg", "Azure", "VMAppContainer;Compute;testRG;testSQL", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProtectionContainerResource = armrecoveryservicesbackup.ProtectionContainerResource{
	// 	Name: to.Ptr("VMAppContainer;Compute;testRG;testSQL"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupFabrics/protectionContainers"),
	// 	ID: to.Ptr("/Subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRg/providers/Microsoft.RecoveryServices/vaults/testVault/backupFabrics/Azure/protectionContainers/VMAppContainer;Compute;testRG;testSQL"),
	// 	Properties: &armrecoveryservicesbackup.AzureVMAppContainerProtectionContainer{
	// 		BackupManagementType: to.Ptr(armrecoveryservicesbackup.BackupManagementTypeAzureWorkload),
	// 		ContainerType: to.Ptr(armrecoveryservicesbackup.ProtectableContainerTypeVMAppContainer),
	// 		FriendlyName: to.Ptr("testSQL"),
	// 		ExtendedInfo: &armrecoveryservicesbackup.AzureWorkloadContainerExtendedInfo{
	// 			HostServerName: to.Ptr("testsql"),
	// 			InquiryInfo: &armrecoveryservicesbackup.InquiryInfo{
	// 				ErrorDetail: &armrecoveryservicesbackup.ErrorDetail{
	// 					Code: to.Ptr("Success"),
	// 					Message: to.Ptr("Not Available"),
	// 					Recommendations: []*string{
	// 						to.Ptr("Not Available")},
	// 					},
	// 					InquiryDetails: []*armrecoveryservicesbackup.WorkloadInquiryDetails{
	// 						{
	// 							Type: to.Ptr("sql"),
	// 							InquiryValidation: &armrecoveryservicesbackup.InquiryValidation{
	// 								ErrorDetail: &armrecoveryservicesbackup.ErrorDetail{
	// 									Code: to.Ptr("Success"),
	// 									Message: to.Ptr("Not Available"),
	// 									Recommendations: []*string{
	// 										to.Ptr("Not Available")},
	// 									},
	// 									Status: to.Ptr("Success"),
	// 								},
	// 								ItemCount: to.Ptr[int64](14),
	// 						}},
	// 						Status: to.Ptr("Success"),
	// 					},
	// 					NodesList: []*armrecoveryservicesbackup.DistributedNodesInfo{
	// 					},
	// 				},
	// 				SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.Compute/virtualMachines/testSQL"),
	// 			},
	// 		}
}
