package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createOrUpdateConnection.json
func ExampleConnectionClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
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
				"SubscriptionID":            to.Ptr("subid"),
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
	// 	Name: to.Ptr("mysConnection"),
	// 	Properties: &armautomation.ConnectionProperties{
	// 		Description: to.Ptr("my description goes here"),
	// 		ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
	// 			Name: to.Ptr("Azure"),
	// 		},
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:29.290Z"); return t}()),
	// 		FieldDefinitionValues: map[string]*string{
	// 			"AutomationCertificateName": to.Ptr("mysCertificateName"),
	// 			"SubscriptionID": to.Ptr("subid"),
	// 		},
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T07:52:29.290Z"); return t}()),
	// 	},
	// }
}
