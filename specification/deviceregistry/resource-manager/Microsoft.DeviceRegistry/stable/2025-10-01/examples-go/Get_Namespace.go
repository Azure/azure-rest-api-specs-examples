package armdeviceregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceregistry/armdeviceregistry/v2"
)

// Generated from example definition: 2025-10-01/Get_Namespace.json
func ExampleNamespacesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNamespacesClient().Get(ctx, "myResourceGroup", "adr-namespace-gbk0925-n01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdeviceregistry.NamespacesClientGetResponse{
	// 	Namespace: &armdeviceregistry.Namespace{
	// 		ID: to.Ptr("/subscriptions/00000000-1366-430f-0000-cc873bcf2d27/resourceGroups/gbktestRG/providers/Microsoft.DeviceRegistry/namespaces/adr-namespace-gbk1001-n01"),
	// 		Name: to.Ptr("adr-namespace-gbk1001-n01"),
	// 		Type: to.Ptr("Microsoft.DeviceRegistry/namespaces"),
	// 		Location: to.Ptr("North Europe"),
	// 		SystemData: &armdeviceregistry.SystemData{
	// 			CreatedBy: to.Ptr("00003442-0000-0000-0000-494059220000"),
	// 			CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-01T21:39:27.7050447Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("00003442-0000-0000-0000-494059220000"),
	// 			LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-01T21:39:27.7050447Z"); return t}()),
	// 		},
	// 		Identity: &armdeviceregistry.SystemAssignedServiceIdentity{
	// 			PrincipalID: to.Ptr("4a06b859-8c07-4d9f-822d-3348a09f72c2"),
	// 			TenantID: to.Ptr("0006f47a-0000-0000-0000-99be82dea000"),
	// 			Type: to.Ptr(armdeviceregistry.SystemAssignedServiceIdentityTypeSystemAssigned),
	// 		},
	// 		Properties: &armdeviceregistry.NamespaceProperties{
	// 			UUID: to.Ptr("cfbef47a-6971-4c90-a7a9-99be82dea167"),
	// 			Messaging: &armdeviceregistry.Messaging{
	// 				Endpoints: map[string]*armdeviceregistry.MessagingEndpoint{
	// 					"myPrimaryEventGridEndpoint": &armdeviceregistry.MessagingEndpoint{
	// 						Address: to.Ptr("https://myeventgridtopic1.westeurope-1.eventgrid.azure.net"),
	// 						EndpointType: to.Ptr("Microsoft.Devices/IoTHubs"),
	// 					},
	// 					"mySecondaryEventGridEndpoint": &armdeviceregistry.MessagingEndpoint{
	// 						Address: to.Ptr("https://myeventgridtopic2.westeurope-1.eventgrid.azure.net"),
	// 						EndpointType: to.Ptr("Microsoft.Devices/IoTHubs"),
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
