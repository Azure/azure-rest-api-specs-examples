package armdomainregistration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/domainregistration/armdomainregistration"
)

// Generated from example definition: 2024-11-01/ListDomainOwnershipIdentifiers.json
func ExampleDomainsClient_NewListOwnershipIdentifiersPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdomainregistration.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDomainsClient().NewListOwnershipIdentifiersPager("testrg123", "example.com", nil)
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
		// page = armdomainregistration.DomainsClientListOwnershipIdentifiersResponse{
		// 	DomainOwnershipIdentifierCollection: armdomainregistration.DomainOwnershipIdentifierCollection{
		// 		Value: []*armdomainregistration.DomainOwnershipIdentifier{
		// 			{
		// 				Name: to.Ptr("SampleOwnershipId"),
		// 				Type: to.Ptr("Microsoft.DomainRegistration/domains/domainownershipidentifiers"),
		// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.DomainRegistration/domains/example.com/domainownershipidentifiers/SampleOwnershipId"),
		// 				Properties: &armdomainregistration.DomainOwnershipIdentifierProperties{
		// 					OwnershipID: to.Ptr("SampleOwnershipId"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
