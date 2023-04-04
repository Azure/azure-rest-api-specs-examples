package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getDscNodeConfiguration.json
func ExampleDscNodeConfigurationClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDscNodeConfigurationClient().Get(ctx, "rg", "myAutomationAccount33", "SetupServer.localhost", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DscNodeConfiguration = armautomation.DscNodeConfiguration{
	// 	Name: to.Ptr("SetupServer.localhost"),
	// 	Type: to.Ptr("Microsoft.Automation/AutomationAccounts/nodeConfigurations"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodeConfigurations/SetupServer.localhost"),
	// 	Properties: &armautomation.DscNodeConfigurationProperties{
	// 		Configuration: &armautomation.DscConfigurationAssociationProperty{
	// 			Name: to.Ptr("SetupServer"),
	// 		},
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:17:06.8901968+00:00"); return t}()),
	// 		IncrementNodeConfigurationBuild: to.Ptr(false),
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:17:06.8901968+00:00"); return t}()),
	// 		NodeCount: to.Ptr[int64](0),
	// 	},
	// }
}
