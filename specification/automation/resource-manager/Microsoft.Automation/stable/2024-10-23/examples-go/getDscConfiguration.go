package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: 2024-10-23/getDscConfiguration.json
func ExampleDscConfigurationClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
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
	// res = armautomation.DscConfigurationClientGetResponse{
	// 	DscConfiguration: armautomation.DscConfiguration{
	// 		Name: to.Ptr("TemplateBasic"),
	// 		Type: to.Ptr("Microsoft.Automation/AutomationAccounts/Configurations"),
	// 		Etag: to.Ptr("\"636263396635600000\""),
	// 		ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/configurations/TemplateBasic"),
	// 		Location: to.Ptr("East US 2"),
	// 		Properties: &armautomation.DscConfigurationProperties{
	// 			Description: to.Ptr("sample configuration"),
	// 			CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:23.56+00:00"); return t}()),
	// 			JobCount: to.Ptr[int32](0),
	// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:23.56+00:00"); return t}()),
	// 			LogVerbose: to.Ptr(false),
	// 			Parameters: map[string]*armautomation.DscConfigurationParameter{
	// 			},
	// 			State: to.Ptr(armautomation.DscConfigurationStatePublished),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
