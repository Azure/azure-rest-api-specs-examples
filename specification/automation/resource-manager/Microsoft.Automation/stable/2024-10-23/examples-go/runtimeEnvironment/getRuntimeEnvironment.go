package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: 2024-10-23/runtimeEnvironment/getRuntimeEnvironment.json
func ExampleRuntimeEnvironmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRuntimeEnvironmentsClient().Get(ctx, "rg", "myAutomationAccount9", "myRuntimeEnvironmentName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armautomation.RuntimeEnvironmentsClientGetResponse{
	// 	RuntimeEnvironment: armautomation.RuntimeEnvironment{
	// 		Name: to.Ptr("myRuntimeEnvironmentName"),
	// 		Type: to.Ptr("Microsoft.Automation/automationAccounts/runtimeEnvironments"),
	// 		ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount9/runtimeEnvironments/myRuntimeEnvironmentName"),
	// 		Location: to.Ptr("East US 2"),
	// 		Properties: &armautomation.RuntimeEnvironmentProperties{
	// 			Description: to.Ptr("Description of the Runtime Environment"),
	// 			DefaultPackages: map[string]*string{
	// 				"Az": to.Ptr("12.3.0"),
	// 			},
	// 			Runtime: &armautomation.RuntimeProperties{
	// 				Version: to.Ptr("7.4"),
	// 				Language: to.Ptr("PowerShell"),
	// 			},
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
