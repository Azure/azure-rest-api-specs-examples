package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/updateConnection.json
func ExampleConnectionClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectionClient().Update(ctx, "rg", "myAutomationAccount28", "myConnection", armautomation.ConnectionUpdateParameters{
		Name: to.Ptr("myConnection"),
		Properties: &armautomation.ConnectionUpdateProperties{
			Description: to.Ptr("my description goes here"),
			FieldDefinitionValues: map[string]*string{
				"AutomationCertificateName": to.Ptr("myCertificateName"),
				"SubscriptionID":            to.Ptr("b5e4748c-f69a-467c-8749-e2f9c8cd3009"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Connection = armautomation.Connection{
	// 	Name: to.Ptr("myConnection"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount29/connections/myConnection"),
	// 	Properties: &armautomation.ConnectionProperties{
	// 		Description: to.Ptr("my description goes here"),
	// 		ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
	// 			Name: to.Ptr("Azure"),
	// 		},
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T08:21:29.313+00:00"); return t}()),
	// 		FieldDefinitionValues: map[string]*string{
	// 			"AutomationCertificateName": to.Ptr("myCertificateName"),
	// 			"SubscriptionID": to.Ptr("b5e4748c-f69a-467c-8749-e2f9c8cd3009"),
	// 		},
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T08:21:30.093+00:00"); return t}()),
	// 	},
	// }
}
