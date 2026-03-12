package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: 2025-07-15-preview/Domains_ListByResourceGroup.json
func ExampleDomainsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("8f6b6269-84f2-4d09-9e31-1127efcd1e40", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDomainsClient().NewListByResourceGroupPager("examplerg", nil)
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
		// page = armeventgrid.DomainsClientListByResourceGroupResponse{
		// 	DomainsListResult: armeventgrid.DomainsListResult{
		// 		Value: []*armeventgrid.Domain{
		// 			{
		// 				Name: to.Ptr("exampledomain1"),
		// 				Type: to.Ptr("Microsoft.EventGrid/domains"),
		// 				ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/domains/exampledomain1"),
		// 				Location: to.Ptr("westus2"),
		// 				Properties: &armeventgrid.DomainProperties{
		// 					Endpoint: to.Ptr("https://exampledomain1.westus2-1.eventgrid.azure.net/api/events"),
		// 					ProvisioningState: to.Ptr(armeventgrid.DomainProvisioningStateSucceeded),
		// 				},
		// 				Tags: map[string]*string{
		// 					"tag1": to.Ptr("value1"),
		// 					"tag2": to.Ptr("value2"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("exampledomain2"),
		// 				Type: to.Ptr("Microsoft.EventGrid/domains"),
		// 				ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/domains/exampledomain2"),
		// 				Location: to.Ptr("westcentralus"),
		// 				Properties: &armeventgrid.DomainProperties{
		// 					Endpoint: to.Ptr("https://exampledomain2.westcentralus-1.eventgrid.azure.net/api/events"),
		// 					ProvisioningState: to.Ptr(armeventgrid.DomainProvisioningStateSucceeded),
		// 				},
		// 				Tags: map[string]*string{
		// 					"tag1": to.Ptr("value1"),
		// 					"tag2": to.Ptr("value2"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
