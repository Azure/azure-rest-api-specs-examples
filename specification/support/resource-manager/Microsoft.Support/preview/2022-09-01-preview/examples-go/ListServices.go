package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/ListServices.json
func ExampleServicesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServicesClient().NewListPager(nil)
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
		// page.ServicesListResult = armsupport.ServicesListResult{
		// 	Value: []*armsupport.Service{
		// 		{
		// 			Name: to.Ptr("service_guid_1"),
		// 			Type: to.Ptr("Microsoft.Support/services"),
		// 			ID: to.Ptr("/providers/Microsoft.Support/services/service_guid_1"),
		// 			Properties: &armsupport.ServiceProperties{
		// 				DisplayName: to.Ptr("Billing"),
		// 				ResourceTypes: []*string{
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("service_guid_2"),
		// 			Type: to.Ptr("Microsoft.Support/services"),
		// 			ID: to.Ptr("/providers/Microsoft.Support/services/service_guid_2"),
		// 			Properties: &armsupport.ServiceProperties{
		// 				DisplayName: to.Ptr("Service and subscription limits (quotas)"),
		// 				ResourceTypes: []*string{
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("service_guid_3"),
		// 			Type: to.Ptr("Microsoft.Support/services"),
		// 			ID: to.Ptr("/providers/Microsoft.Support/services/service_guid_3"),
		// 			Properties: &armsupport.ServiceProperties{
		// 				DisplayName: to.Ptr("Subscription management"),
		// 				ResourceTypes: []*string{
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("service_guid_4"),
		// 			Type: to.Ptr("Microsoft.Support/services"),
		// 			ID: to.Ptr("/providers/Microsoft.Support/services/service_guid_4"),
		// 			Properties: &armsupport.ServiceProperties{
		// 				DisplayName: to.Ptr("Data Explorer"),
		// 				ResourceTypes: []*string{
		// 					to.Ptr("MICROSOFT.KUSTO/CLUSTERS"),
		// 					to.Ptr("MICROSOFT.KUSTO/DATABASES")},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("service_guid_5"),
		// 				Type: to.Ptr("Microsoft.Support/services"),
		// 				ID: to.Ptr("/providers/Microsoft.Support/services/service_guid_5"),
		// 				Properties: &armsupport.ServiceProperties{
		// 					DisplayName: to.Ptr("Virtual Machine running Windows"),
		// 					ResourceTypes: []*string{
		// 						to.Ptr("MICROSOFT.CLASSICCOMPUTE/VIRTUALMACHINES"),
		// 						to.Ptr("MICROSOFT.COMPUTE/VIRTUALMACHINES")},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("service_guid_6"),
		// 					Type: to.Ptr("Microsoft.Support/services"),
		// 					ID: to.Ptr("/providers/Microsoft.Support/services/service_guid_6"),
		// 					Properties: &armsupport.ServiceProperties{
		// 						DisplayName: to.Ptr("Virtual Machine running Linux"),
		// 						ResourceTypes: []*string{
		// 							to.Ptr("MICROSOFT.CLASSICCOMPUTE/VIRTUALMACHINES"),
		// 							to.Ptr("MICROSOFT.COMPUTE/VIRTUALMACHINES")},
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("service_guid_7"),
		// 						Type: to.Ptr("Microsoft.Support/services"),
		// 						ID: to.Ptr("/providers/Microsoft.Support/services/service_guid_7"),
		// 						Properties: &armsupport.ServiceProperties{
		// 							DisplayName: to.Ptr("Virtual Network"),
		// 							ResourceTypes: []*string{
		// 								to.Ptr("MICROSOFT.NETWORK/VIRTUALNETWORKS"),
		// 								to.Ptr("MICROSOFT.CLASSICNETWORK/VIRTUALNETWORKS")},
		// 							},
		// 					}},
		// 				}
	}
}
