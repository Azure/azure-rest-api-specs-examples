package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: 2024-10-23/createOrUpdateDscConfiguration.json
func ExampleDscConfigurationClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDscConfigurationClient().CreateOrUpdate(ctx, "rg", "myAutomationAccount18", "SetupServer", armautomation.DscConfigurationCreateOrUpdateParameters{
		Name:     to.Ptr("SetupServer"),
		Location: to.Ptr("East US 2"),
		Properties: &armautomation.DscConfigurationCreateOrUpdateProperties{
			Description: to.Ptr("sample configuration"),
			Source: &armautomation.ContentSource{
				Type: to.Ptr(armautomation.ContentSourceTypeEmbeddedContent),
				Hash: &armautomation.ContentHash{
					Algorithm: to.Ptr("sha256"),
					Value:     to.Ptr("A9E5DB56BA21513F61E0B3868816FDC6D4DF5131F5617D7FF0D769674BD5072F"),
				},
				Value: to.Ptr("Configuration SetupServer {\r\n    Node localhost {\r\n                               WindowsFeature IIS {\r\n                               Name = \"Web-Server\";\r\n            Ensure = \"Present\"\r\n        }\r\n    }\r\n}"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armautomation.DscConfigurationClientCreateOrUpdateResponse{
	// 	DscConfiguration: armautomation.DscConfiguration{
	// 		Name: to.Ptr("SetupServer"),
	// 		Type: to.Ptr("Microsoft.Automation/AutomationAccounts/Configurations"),
	// 		ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/configurations/SetupServer"),
	// 		Location: to.Ptr("East US 2"),
	// 		Properties: &armautomation.DscConfigurationProperties{
	// 			Description: to.Ptr("sample configuration"),
	// 			CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T00:00:00+00:00"); return t}()),
	// 			JobCount: to.Ptr[int32](0),
	// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T00:00:00+00:00"); return t}()),
	// 			LogVerbose: to.Ptr(false),
	// 			Parameters: map[string]*armautomation.DscConfigurationParameter{
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
