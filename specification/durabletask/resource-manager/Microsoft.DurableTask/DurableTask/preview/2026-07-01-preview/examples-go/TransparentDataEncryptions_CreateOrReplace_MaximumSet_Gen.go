package armdurabletask_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/durabletask/armdurabletask"
)

// Generated from example definition: 2026-07-01-preview/TransparentDataEncryptions_CreateOrReplace_MaximumSet_Gen.json
func ExampleTransparentDataEncryptionsClient_BeginCreateOrReplace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdurabletask.NewClientFactory("EE9BD735-67CE-4A90-89C4-439D3F6A4C93", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTransparentDataEncryptionsClient().BeginCreateOrReplace(ctx, "rgdurabletask", "testscheduler", armdurabletask.TransparentDataEncryption{
		Properties: &armdurabletask.TransparentDataEncryptionProperties{
			KeySource:           to.Ptr(armdurabletask.TransparentDataEncryptionKeySourceCustomerManaged),
			KeyVaultKeyURI:      to.Ptr("https://microsoft.com/ap"),
			AutoRotationEnabled: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdurabletask.TransparentDataEncryptionsClientCreateOrReplaceResponse{
	// 	TransparentDataEncryption: armdurabletask.TransparentDataEncryption{
	// 		Properties: &armdurabletask.TransparentDataEncryptionProperties{
	// 			ProvisioningState: to.Ptr(armdurabletask.ProvisioningStateSucceeded),
	// 			KeySource: to.Ptr(armdurabletask.TransparentDataEncryptionKeySourceCustomerManaged),
	// 			KeyVaultKeyURI: to.Ptr("https://microsoft.com/ap"),
	// 			AutoRotationEnabled: to.Ptr(true),
	// 		},
	// 		ID: to.Ptr("/subscriptions/EE9BD735-67CE-4A90-89C4-439D3F6A4C93/resourceGroups/rgdurabletask/providers/Microsoft.DurableTask/schedulers/testscheduler/transparentDataEncryptions/default"),
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.DurableTask/schedulers/transparentDataEncryptions"),
	// 		SystemData: &armdurabletask.SystemData{
	// 			CreatedBy: to.Ptr("tenmbevaunjzikxowqexrsx"),
	// 			CreatedByType: to.Ptr(armdurabletask.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2024, time.April, 17, 15, 34, 17, 365000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("xfvdcegtj"),
	// 			LastModifiedByType: to.Ptr(armdurabletask.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2024, time.April, 17, 15, 34, 17, 366000000, time.UTC)),
	// 		},
	// 	},
	// }
}
