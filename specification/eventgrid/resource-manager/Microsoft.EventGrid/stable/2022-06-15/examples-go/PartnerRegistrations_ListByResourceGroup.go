package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f88928d723133dc392e3297e6d61b7f6d10501fd/specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/PartnerRegistrations_ListByResourceGroup.json
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
		// 			ID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/amh/providers/Microsoft.EventGrid/partnerRegistrations/ContosoCorpAccount1"),
		// 			Location: to.Ptr("global"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("Value2"),
		// 				"key3": to.Ptr("Value3"),
		// 			},
		// 			Properties: &armeventgrid.PartnerRegistrationProperties{
		// 				PartnerRegistrationImmutableID: to.Ptr("6a1e637f-1495-4938-bf46-ff468b9a75d2"),
		// 				ProvisioningState: to.Ptr(armeventgrid.PartnerRegistrationProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
