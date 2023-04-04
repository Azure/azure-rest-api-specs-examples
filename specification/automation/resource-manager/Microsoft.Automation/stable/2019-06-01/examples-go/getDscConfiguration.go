package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/getDscConfiguration.json
func ExampleDscConfigurationClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDscConfigurationClient().Get(ctx, "rg", "myAutomationAccount33", "TemplateBasic", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DscConfiguration = armautomation.DscConfiguration{
	// 	Name: to.Ptr("TemplateBasic"),
	// 	Type: to.Ptr("Microsoft.Automation/AutomationAccounts/Configurations"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/configurations/TemplateBasic"),
	// 	Location: to.Ptr("East US 2"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Etag: to.Ptr("\"636263396635600000\""),
	// 	Properties: &armautomation.DscConfigurationProperties{
	// 		Description: to.Ptr("sample configuration"),
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:23.56+00:00"); return t}()),
	// 		JobCount: to.Ptr[int32](0),
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:23.56+00:00"); return t}()),
	// 		LogVerbose: to.Ptr(false),
	// 		Parameters: map[string]*armautomation.DscConfigurationParameter{
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		State: to.Ptr(armautomation.DscConfigurationStatePublished),
	// 	},
	// }
}
