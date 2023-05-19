package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/173bb3b6fd5b1809fdbf347f67fccfa0440ac126/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/PartnerRegistrations_ListByResourceGroup.json
func ExamplePartnerRegistrationsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPartnerRegistrationsClient().NewListByResourceGroupPager("examplerg", &armeventgrid.PartnerRegistrationsClientListByResourceGroupOptions{Filter: nil,
		Top: nil,
	})
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
		// page.PartnerRegistrationsListResult = armeventgrid.PartnerRegistrationsListResult{
		// 	Value: []*armeventgrid.PartnerRegistration{
		// 		{
		// 			Name: to.Ptr("ContosoCorpAccount1"),
		// 			Type: to.Ptr("Microsoft.EventGrid/partnerRegistrations"),
		// 			ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/amh/providers/Microsoft.EventGrid/partnerRegistrations/ContosoCorpAccount1"),
		// 			Location: to.Ptr("global"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("Value2"),
		// 				"key3": to.Ptr("Value3"),
		// 			},
		// 			Properties: &armeventgrid.PartnerRegistrationProperties{
		// 				PartnerRegistrationImmutableID: to.Ptr("cda82399-79fe-4d5a-bc6d-b05a437204d9"),
		// 				ProvisioningState: to.Ptr(armeventgrid.PartnerRegistrationProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
