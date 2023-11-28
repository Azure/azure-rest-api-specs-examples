package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getPython2Package.json
func ExamplePython2PackageClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.Module = armautomation.Module{
	// 	Name: to.Ptr("OmsCompositeResources"),
	// 	Type: to.Ptr("Microsoft.Automation/AutomationAccounts/Python2Packages"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/python2Packages/OmsCompositeResources"),
	// 	Location: to.Ptr("East US 2"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armautomation.ModuleProperties{
	// 		ActivityCount: to.Ptr[int32](0),
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-29T15:41:47.003Z"); return t}()),
	// 		Error: &armautomation.ModuleErrorInfo{
	// 			Message: to.Ptr(""),
	// 		},
	// 		IsComposite: to.Ptr(true),
	// 		IsGlobal: to.Ptr(false),
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-29T15:42:10.567Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armautomation.ModuleProvisioningStateSucceeded),
	// 		SizeInBytes: to.Ptr[int64](0),
	// 	},
	// }
}
