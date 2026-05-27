package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: 2024-10-23/package/getPackage.json
func ExamplePackageClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPackageClient().Get(ctx, "rg", "myAutomationAccount33", "runtimeEnvironmentName", "OmsCompositeResources", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armautomation.PackageClientGetResponse{
	// 	Package: armautomation.Package{
	// 		Name: to.Ptr("OmsCompositeResources"),
	// 		Type: to.Ptr("Microsoft.Automation/automationAccounts/runtimeEnvironments/packages"),
	// 		ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/runtimeEnvironments/runtimeEnvironmentName/packages/OmsCompositeResources"),
	// 		Location: to.Ptr("East US 2"),
	// 		Properties: &armautomation.PackageProperties{
	// 			Default: to.Ptr(false),
	// 			ProvisioningState: to.Ptr(armautomation.PackageProvisioningStateSucceeded),
	// 			SizeInBytes: to.Ptr[int64](1560),
	// 			Version: to.Ptr("1.0.0"),
	// 		},
	// 		SystemData: &armautomation.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-05T07:32:41.4389914+00:00"); return t}()),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-05T07:33:07.5597465+00:00"); return t}()),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
