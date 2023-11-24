package armwindowsesu_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/windowsesu/armwindowsesu"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/windowsesu/resource-manager/Microsoft.WindowsESU/preview/2019-09-16-preview/examples/GetMultipleActivationKey.json
func ExampleMultipleActivationKeysClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armwindowsesu.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMultipleActivationKeysClient().Get(ctx, "testgr1", "server08-key-2019", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MultipleActivationKey = armwindowsesu.MultipleActivationKey{
	// 	Name: to.Ptr("server08-key-2019"),
	// 	Type: to.Ptr("Microsoft.WindowsESU/multipleActivationKeys"),
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testgr1/providers/Microsoft.WindowsESU/multipleActivationKeys/server08-key-2019"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armwindowsesu.MultipleActivationKeyProperties{
	// 		AgreementNumber: to.Ptr("1a2b45ag"),
	// 		ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-21T21:12:26.000Z"); return t}()),
	// 		InstalledServerNumber: to.Ptr[int32](100),
	// 		IsEligible: to.Ptr(true),
	// 		MultipleActivationKey: to.Ptr("aaaaa-bbbbb-ccccc-ddddd-eeeee"),
	// 		OSType: to.Ptr(armwindowsesu.OsTypeWindowsServer2008),
	// 		ProvisioningState: to.Ptr(armwindowsesu.ProvisioningStateSucceeded),
	// 		SupportType: to.Ptr(armwindowsesu.SupportTypeSupplementalServicing),
	// 	},
	// }
}
