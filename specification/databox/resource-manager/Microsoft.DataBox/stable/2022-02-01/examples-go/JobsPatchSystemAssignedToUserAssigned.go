package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databox/resource-manager/Microsoft.DataBox/stable/2022-02-01/examples/JobsPatchSystemAssignedToUserAssigned.json
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
	poller, err := clientFactory.NewJobsClient().BeginUpdate(ctx, "SdkRg9765", "SdkJob2965", armdatabox.JobResourceUpdateParameter{
		Identity: &armdatabox.ResourceIdentity{
			Type: to.Ptr("SystemAssigned,UserAssigned"),
			UserAssignedIdentities: map[string]*armdatabox.UserAssignedIdentity{
				"/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/akvenkat/providers/Microsoft.ManagedIdentity/userAssignedIdentities/sdkIdentity": {},
			},
		},
		Properties: &armdatabox.UpdateJobProperties{
			Details: &armdatabox.UpdateJobDetails{
				KeyEncryptionKey: &armdatabox.KeyEncryptionKey{
					IdentityProperties: &armdatabox.IdentityProperties{
						Type: to.Ptr("UserAssigned"),
						UserAssigned: &armdatabox.UserAssignedProperties{
							ResourceID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/akvenkat/providers/Microsoft.ManagedIdentity/userAssignedIdentities/sdkIdentity"),
						},
					},
					KekType:            to.Ptr(armdatabox.KekTypeCustomerManaged),
					KekURL:             to.Ptr("https://sdkkeyvault.vault.azure.net/keys/SSDKEY/"),
					KekVaultResourceID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/akvenkat/providers/Microsoft.KeyVault/vaults/SDKKeyVault"),
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
	// 	Name: to.Ptr("SdkJob2965"),
	// 	Type: to.Ptr("Microsoft.DataBox/jobs"),
	// 	ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/SdkRg9765/providers/Microsoft.DataBox/jobs/SdkJob2965"),
	// 	Properties: &armdatabox.JobProperties{
	// 		IsCancellable: to.Ptr(true),
	// 		IsShippingAddressEditable: to.Ptr(true),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-13T16:28:38.9999793+05:30"); return t}()),
	// 		Status: to.Ptr(armdatabox.StageNameDeviceOrdered),
	// 		TransferType: to.Ptr(armdatabox.TransferTypeImportToAzure),
	// 	},
	// }
}
