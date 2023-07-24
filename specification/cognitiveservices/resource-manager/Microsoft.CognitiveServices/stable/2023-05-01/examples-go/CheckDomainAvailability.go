package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/CheckDomainAvailability.json
func ExampleManagementClient_CheckDomainAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagementClient().CheckDomainAvailability(ctx, armcognitiveservices.CheckDomainAvailabilityParameter{
		Type:          to.Ptr("Microsoft.CognitiveServices/accounts"),
		SubdomainName: to.Ptr("contosodemoapp1"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DomainAvailability = armcognitiveservices.DomainAvailability{
	// 	Type: to.Ptr("Microsoft.CognitiveServices/accounts"),
	// 	IsSubdomainAvailable: to.Ptr(false),
	// 	Reason: to.Ptr("Sub domain name 'contosodemoapp1' is not valid"),
	// 	SubdomainName: to.Ptr("contosodemoapp1"),
	// }
}
