package armdomainregistration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/domainregistration/armdomainregistration"
)

// Generated from example definition: 2024-11-01/CheckDomainAvailability.json
func ExampleDomainsClient_CheckAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdomainregistration.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDomainsClient().CheckAvailability(ctx, armdomainregistration.NameIdentifier{
		Name: to.Ptr("abcd.com"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdomainregistration.DomainsClientCheckAvailabilityResponse{
	// 	DomainAvailabilityCheckResult: &armdomainregistration.DomainAvailabilityCheckResult{
	// 		Name: to.Ptr("abcd.com"),
	// 		Available: to.Ptr(true),
	// 		DomainType: to.Ptr(armdomainregistration.DomainTypeRegular),
	// 	},
	// }
}
