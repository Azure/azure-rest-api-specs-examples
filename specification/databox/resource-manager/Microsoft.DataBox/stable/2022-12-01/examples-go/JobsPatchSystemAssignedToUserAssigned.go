package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/JobsPatchSystemAssignedToUserAssigned.json
func ExampleJobsClient_BeginUpdate_jobsPatchSystemAssignedToUserAssigned() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabox.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewJobsClient().BeginUpdate(ctx, "YourResourceGroupName", "TestJobName1", armdatabox.JobResourceUpdateParameter{
		Identity: &armdatabox.ResourceIdentity{
			Type: to.Ptr("SystemAssigned,UserAssigned"),
			UserAssignedIdentities: map[string]*armdatabox.UserAssignedIdentity{
				"/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testIdentity": {},
			},
		},
		Properties: &armdatabox.UpdateJobProperties{
			Details: &armdatabox.UpdateJobDetails{
				KeyEncryptionKey: &armdatabox.KeyEncryptionKey{
					IdentityProperties: &armdatabox.IdentityProperties{
						Type: to.Ptr("UserAssigned"),
						UserAssigned: &armdatabox.UserAssignedProperties{
							ResourceID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testIdentity"),
						},
					},
					KekType:            to.Ptr(armdatabox.KekTypeCustomerManaged),
					KekURL:             to.Ptr("https://xxx.xxx.xx"),
					KekVaultResourceID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.KeyVault/vaults/YourKeyVaultName"),
				},
			},
		},
	}, &armdatabox.JobsClientBeginUpdateOptions{IfMatch: nil})
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
	// res.JobResource = armdatabox.JobResource{
	// 	Location: to.Ptr("westus"),
	// 	SKU: &armdatabox.SKU{
	// 		Name: to.Ptr(armdatabox.SKUNameDataBox),
	// 	},
	// 	Tags: map[string]*string{
	// 	},
	// 	Name: to.Ptr("TestJobName1"),
	// 	Type: to.Ptr("Microsoft.DataBox/jobs"),
	// 	ID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.DataBox/jobs/TestJobName1"),
	// 	Properties: &armdatabox.JobProperties{
	// 		IsCancellable: to.Ptr(true),
	// 		IsShippingAddressEditable: to.Ptr(true),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-13T10:58:38.999Z"); return t}()),
	// 		Status: to.Ptr(armdatabox.StageNameDeviceOrdered),
	// 		TransferType: to.Ptr(armdatabox.TransferTypeImportToAzure),
	// 	},
	// }
}
