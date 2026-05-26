package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: 2024-10-23/getPython2Package.json
func ExamplePython2PackageClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPython2PackageClient().Get(ctx, "rg", "myAutomationAccount33", "OmsCompositeResources", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armautomation.Python2PackageClientGetResponse{
	// 	Module: armautomation.Module{
	// 		Name: to.Ptr("OmsCompositeResources"),
	// 		Type: to.Ptr("Microsoft.Automation/AutomationAccounts/Python2Packages"),
	// 		ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/python2Packages/OmsCompositeResources"),
	// 		Location: to.Ptr("East US 2"),
	// 		Properties: &armautomation.ModuleProperties{
	// 			ActivityCount: to.Ptr[int32](0),
	// 			CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-29T15:41:47.003+00:00"); return t}()),
	// 			Error: &armautomation.ModuleErrorInfo{
	// 				Message: to.Ptr(""),
	// 			},
	// 			IsComposite: to.Ptr(true),
	// 			IsGlobal: to.Ptr(false),
	// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-29T15:42:10.567+00:00"); return t}()),
	// 			ProvisioningState: to.Ptr(armautomation.ModuleProvisioningStateCreating),
	// 			SizeInBytes: to.Ptr[int64](0),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
