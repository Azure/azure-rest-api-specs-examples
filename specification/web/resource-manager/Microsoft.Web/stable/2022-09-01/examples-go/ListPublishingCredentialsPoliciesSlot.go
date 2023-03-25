package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/10c740b7224771c889cfb92f128168f5a0568c26/specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/ListPublishingCredentialsPoliciesSlot.json
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
