package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/82e9c6f9fbfa2d6d47d5e2a6a11c0ad2eb345c43/specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/ListPublishingCredentialsPoliciesSlot.json
func ExampleWebAppsClient_NewListBasicPublishingCredentialsPoliciesSlotPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWebAppsClient().NewListBasicPublishingCredentialsPoliciesSlotPager("testrg123", "testsite", "staging", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.PublishingCredentialsPoliciesCollection = armappservice.PublishingCredentialsPoliciesCollection{
		// 	Value: []*armappservice.CsmPublishingCredentialsPoliciesEntity{
		// 		{
		// 			Name: to.Ptr("ftp"),
		// 			Type: to.Ptr("Microsoft.Web/sites/slots/basicPublishingCredentialsPolicies"),
		// 			ID: to.Ptr("/subscriptions/3fb8d758-2e2c-42e9-a528-a8acdfe87237/resourceGroups/testrg123/providers/Microsoft.Web/sites/testsite/slots/staging/basicPublishingCredentialsPolicies/ftp"),
		// 			Properties: &armappservice.CsmPublishingCredentialsPoliciesEntityProperties{
		// 				Allow: to.Ptr(true),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("scm"),
		// 			Type: to.Ptr("Microsoft.Web/sites/slots/basicPublishingCredentialsPolicies"),
		// 			ID: to.Ptr("/subscriptions/3fb8d758-2e2c-42e9-a528-a8acdfe87237/resourceGroups/testrg123/providers/Microsoft.Web/sites/testsite/slots/staging/basicPublishingCredentialsPolicies/scm"),
		// 			Properties: &armappservice.CsmPublishingCredentialsPoliciesEntityProperties{
		// 				Allow: to.Ptr(true),
		// 			},
		// 	}},
		// }
	}
}
