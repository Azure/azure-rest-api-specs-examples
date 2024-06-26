package armsaas_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/saas/armsaas"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/saas/resource-manager/Microsoft.SaaS/preview/2018-03-01-beta/examples/saasV2/SaasPatch.json
func ExampleClient_BeginUpdateResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsaas.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginUpdateResource(ctx, "115c3523-1fae-757f-af86-7b27cfd29805", armsaas.ResourceCreation{
		Properties: &armsaas.CreationProperties{
			SKUID: to.Ptr("premium"),
		},
		Tags: map[string]*string{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Resource = armsaas.Resource{
	// 	Name: to.Ptr("MyContosoSubscription"),
	// 	Type: to.Ptr("Microsoft.SaaS/saasresources"),
	// 	ID: to.Ptr("/providers/Microsoft.SaaS/saasresources/115c3523-1fae-757f-af86-7b27cfd29805"),
	// 	Properties: &armsaas.ResourceProperties{
	// 		AutoRenew: to.Ptr(true),
	// 		OfferID: to.Ptr("contosoOffer"),
	// 		PaymentChannelMetadata: map[string]*string{
	// 			"azureSubscriptionId": to.Ptr("c825645b-e31b-9cf4-1cee-2aba9e58bc7c"),
	// 			"resourceId": to.Ptr("263ebe8c-3621-4ac0-a6ba-f1419bfb9166"),
	// 		},
	// 		PaymentChannelType: to.Ptr(armsaas.PaymentChannelTypeSubscriptionDelegated),
	// 		PublisherID: to.Ptr("microsoft-contoso"),
	// 		SaasResourceName: to.Ptr("MyContosoSubscription"),
	// 		SKUID: to.Ptr("premium"),
	// 		Created: to.Ptr("2021-01-01T08:30:10.1234567Z"),
	// 		IsFreeTrial: to.Ptr(false),
	// 		LastModified: to.Ptr("2020-01-01T08:35:05.7654321Z"),
	// 		Status: to.Ptr(armsaas.SaasResourceStatusSubscribed),
	// 		Term: &armsaas.PropertiesTerm{
	// 			EndDate: to.Ptr("2021-02-31T00:00:00Z"),
	// 			StartDate: to.Ptr("2021-01-01T00:00:00Z"),
	// 			TermUnit: to.Ptr("P1M"),
	// 		},
	// 	},
	// 	Tags: map[string]*string{
	// 	},
	// }
}
