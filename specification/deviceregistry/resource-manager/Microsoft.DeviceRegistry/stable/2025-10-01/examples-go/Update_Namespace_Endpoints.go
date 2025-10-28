package armdeviceregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceregistry/armdeviceregistry/v2"
)

// Generated from example definition: 2025-10-01/Update_Namespace_Endpoints.json
func ExampleNamespacesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNamespacesClient().BeginUpdate(ctx, "myResourceGroup", "adr-namespace-gbk0925-n01", armdeviceregistry.NamespaceUpdate{
		Properties: &armdeviceregistry.NamespaceUpdateProperties{
			Messaging: &armdeviceregistry.Messaging{
				Endpoints: map[string]*armdeviceregistry.MessagingEndpoint{
					"eventGridEndpoint": {
						EndpointType: to.Ptr("Microsoft.Devices/IoTHubs"),
						Address:      to.Ptr("https://myeventgridtopic.westeurope-1.eventgrid.azure.net/api/events"),
					},
				},
			},
		},
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
	// res = armdeviceregistry.NamespacesClientUpdateResponse{
	// 	Namespace: &armdeviceregistry.Namespace{
	// 		ID: to.Ptr("/subscriptions/00000000-1366-430f-0000-cc873bcf2d27/resourceGroups/gbktestRG/providers/Microsoft.DeviceRegistry/namespaces/adr-namespace-gbk0925-n01"),
	// 		Name: to.Ptr("adr-namespace-gbk0925-n01"),
	// 		Type: to.Ptr("Microsoft.DeviceRegistry/namespaces"),
	// 		Location: to.Ptr("North Europe"),
	// 		SystemData: &armdeviceregistry.SystemData{
	// 			CreatedBy: to.Ptr("00003442-0000-0000-0000-494059220000"),
	// 			CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-09-25T23:41:41.8591157Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("00003442-0000-0000-0000-494059220000"),
	// 			LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-01T21:03:33.5993192Z"); return t}()),
	// 		},
	// 		Identity: &armdeviceregistry.SystemAssignedServiceIdentity{
	// 			PrincipalID: to.Ptr("00000000-0000-0000-9d20-8a5570c3eb6e"),
	// 			TenantID: to.Ptr("0006f47a-0000-0000-0000-99be82dea000"),
	// 			Type: to.Ptr(armdeviceregistry.SystemAssignedServiceIdentityTypeSystemAssigned),
	// 		},
	// 		Properties: &armdeviceregistry.NamespaceProperties{
	// 			UUID: to.Ptr("cbfe124a-6971-4c90-a7a9-99be82def1ab"),
	// 			ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
	// 			Messaging: &armdeviceregistry.Messaging{
	// 				Endpoints: map[string]*armdeviceregistry.MessagingEndpoint{
	// 					"eventGridEndpoint1": &armdeviceregistry.MessagingEndpoint{
	// 						EndpointType: to.Ptr("Microsoft.Devices/IoTHubs"),
	// 						Address: to.Ptr("https://myeventgridtopic1.westeurope-1.eventgrid.azure.net/api/events"),
	// 					},
	// 					"eventGridEndpoint2": &armdeviceregistry.MessagingEndpoint{
	// 						EndpointType: to.Ptr("Microsoft.Devices/IoTHubs"),
	// 						Address: to.Ptr("https://myeventgridtopic2.westeurope-1.eventgrid.azure.net/api/events"),
	// 					},
	// 					"eventGridEndpoint3": &armdeviceregistry.MessagingEndpoint{
	// 						EndpointType: to.Ptr("Microsoft.Devices/IoTHubs"),
	// 						Address: to.Ptr("https://myeventgridtopic3.westeurope-1.eventgrid.azure.net/api/events"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
