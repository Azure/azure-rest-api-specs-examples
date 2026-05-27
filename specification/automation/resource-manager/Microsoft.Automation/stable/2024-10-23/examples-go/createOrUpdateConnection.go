package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: 2024-10-23/createOrUpdateConnection.json
func ExampleConnectionClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectionClient().CreateOrUpdate(ctx, "rg", "myAutomationAccount28", "mysConnection", armautomation.ConnectionCreateOrUpdateParameters{
		Name: to.Ptr("mysConnection"),
		Properties: &armautomation.ConnectionCreateOrUpdateProperties{
			Description: to.Ptr("my description goes here"),
			ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
				Name: to.Ptr("Azure"),
			},
			FieldDefinitionValues: map[string]*string{
				"AutomationCertificateName": to.Ptr("mysCertificateName"),
				"SubscriptionID":            to.Ptr("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armautomation.ConnectionClientCreateOrUpdateResponse{
	// 	Connection: armautomation.Connection{
	// 		Name: to.Ptr("mysConnection"),
	// 		ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount28/connections/mysConnection"),
	// 		Properties: &armautomation.ConnectionProperties{
	// 			Description: to.Ptr("my description goes here"),
	// 			ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
	// 				Name: to.Ptr("Azure"),
	// 			},
	// 			CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:29.29+00:00"); return t}()),
	// 			FieldDefinitionValues: map[string]*string{
	// 				"AutomationCertificateName": to.Ptr("mysCertificateName"),
	// 				"SubscriptionID": to.Ptr("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
	// 			},
	// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:29.29+00:00"); return t}()),
	// 		},
	// 	},
	// }
}
