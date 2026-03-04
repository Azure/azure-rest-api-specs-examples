package armdomainregistration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/domainregistration/armdomainregistration"
)

// Generated from example definition: 2024-11-01/ListTopLevelDomains.json
func ExampleTopLevelDomainsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdomainregistration.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTopLevelDomainsClient().NewListPager(nil)
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
		// page = armdomainregistration.TopLevelDomainsClientListResponse{
		// 	TopLevelDomainCollection: armdomainregistration.TopLevelDomainCollection{
		// 		Value: []*armdomainregistration.TopLevelDomain{
		// 			{
		// 				Name: to.Ptr("com"),
		// 				Type: to.Ptr("Microsoft.DomainRegistration/topLevelDomains"),
		// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.DomainRegistration/topLevelDomains/com"),
		// 				Properties: &armdomainregistration.TopLevelDomainProperties{
		// 					Privacy: to.Ptr(true),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("net"),
		// 				Type: to.Ptr("Microsoft.DomainRegistration/topLevelDomains"),
		// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.DomainRegistration/topLevelDomains/net"),
		// 				Properties: &armdomainregistration.TopLevelDomainProperties{
		// 					Privacy: to.Ptr(true),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("co.uk"),
		// 				Type: to.Ptr("Microsoft.DomainRegistration/topLevelDomains"),
		// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.DomainRegistration/topLevelDomains/co.uk"),
		// 				Properties: &armdomainregistration.TopLevelDomainProperties{
		// 					Privacy: to.Ptr(false),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("org"),
		// 				Type: to.Ptr("Microsoft.DomainRegistration/topLevelDomains"),
		// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.DomainRegistration/topLevelDomains/org"),
		// 				Properties: &armdomainregistration.TopLevelDomainProperties{
		// 					Privacy: to.Ptr(true),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("nl"),
		// 				Type: to.Ptr("Microsoft.DomainRegistration/topLevelDomains"),
		// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.DomainRegistration/topLevelDomains/nl"),
		// 				Properties: &armdomainregistration.TopLevelDomainProperties{
		// 					Privacy: to.Ptr(true),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("in"),
		// 				Type: to.Ptr("Microsoft.DomainRegistration/topLevelDomains"),
		// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.DomainRegistration/topLevelDomains/in"),
		// 				Properties: &armdomainregistration.TopLevelDomainProperties{
		// 					Privacy: to.Ptr(false),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("biz"),
		// 				Type: to.Ptr("Microsoft.DomainRegistration/topLevelDomains"),
		// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.DomainRegistration/topLevelDomains/biz"),
		// 				Properties: &armdomainregistration.TopLevelDomainProperties{
		// 					Privacy: to.Ptr(true),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("org.uk"),
		// 				Type: to.Ptr("Microsoft.DomainRegistration/topLevelDomains"),
		// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.DomainRegistration/topLevelDomains/org.uk"),
		// 				Properties: &armdomainregistration.TopLevelDomainProperties{
		// 					Privacy: to.Ptr(false),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("co.in"),
		// 				Type: to.Ptr("Microsoft.DomainRegistration/topLevelDomains"),
		// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.DomainRegistration/topLevelDomains/co.in"),
		// 				Properties: &armdomainregistration.TopLevelDomainProperties{
		// 					Privacy: to.Ptr(false),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
