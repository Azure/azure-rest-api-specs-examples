package armdeviceregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceregistry/armdeviceregistry/v2"
)

// Generated from example definition: 2025-10-01/CreateOrReplace_NamespaceDevice.json
func ExampleNamespaceDevicesClient_BeginCreateOrReplace_createOrReplaceNamespaceDevices() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNamespaceDevicesClient().BeginCreateOrReplace(ctx, "myResourceGroup", "adr-namespace-gbk0925-n01", "dev-namespace-gbk0925-n01", armdeviceregistry.NamespaceDevice{
		Location: to.Ptr("West Europe"),
		Properties: &armdeviceregistry.NamespaceDeviceProperties{
			Endpoints: &armdeviceregistry.MessagingEndpoints{
				Outbound: &armdeviceregistry.OutboundEndpoints{
					Assigned: map[string]*armdeviceregistry.DeviceMessagingEndpoint{
						"eventGridEndpoint": {
							EndpointType: to.Ptr("Microsoft.Devices/IoTHubs"),
							Address:      to.Ptr("https://myeventgridtopic.westeurope-1.eventgrid.azure.net/api/events"),
						},
					},
				},
			},
			ExternalDeviceID: to.Ptr("adr-smart-device3-7a848b15-af47-40a7-8c06-a3f43314d44f"),
			Enabled:          to.Ptr(true),
			Attributes: map[string]any{
				"deviceType":     "sensor",
				"deviceOwner":    "IT",
				"deviceCategory": 16,
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
	// res = armdeviceregistry.NamespaceDevicesClientCreateOrReplaceResponse{
	// 	NamespaceDevice: &armdeviceregistry.NamespaceDevice{
	// 		ID: to.Ptr("/subscriptions/00000000-1366-430f-0000-cc873bcf2d27/resourceGroups/gbktestRG/providers/private.deviceregistry/namespaces/my-namespace-1/devices/adr-smart-device3-7a848b15-af47-40a7-8c06-a3f43314d44f"),
	// 		Name: to.Ptr("adr-smart-device3-7a848b15-af47-40a7-8c06-a3f43314d44f"),
	// 		Type: to.Ptr("private.deviceregistry/namespaces/devices"),
	// 		Location: to.Ptr("North Europe"),
	// 		Tags: map[string]*string{
	// 			"sensor": to.Ptr("temperature,humidity"),
	// 		},
	// 		SystemData: &armdeviceregistry.SystemData{
	// 			CreatedBy: to.Ptr("00003442-0000-0000-0000-494059220000"),
	// 			CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-02T20:15:21.8874648Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("00003442-0000-0000-0000-494059220000"),
	// 			LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-02T20:15:21.8874648Z"); return t}()),
	// 		},
	// 		Properties: &armdeviceregistry.NamespaceDeviceProperties{
	// 			Endpoints: &armdeviceregistry.MessagingEndpoints{
	// 				Outbound: &armdeviceregistry.OutboundEndpoints{
	// 					Assigned: map[string]*armdeviceregistry.DeviceMessagingEndpoint{
	// 						"eventGridEndpoint": &armdeviceregistry.DeviceMessagingEndpoint{
	// 							EndpointType: to.Ptr("Microsoft.Devices/IoTHubs"),
	// 							Address: to.Ptr("https://myeventgridtopic.westeurope-1.eventgrid.azure.net/api/events"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ExternalDeviceID: to.Ptr("adr-smart-device3-7a848b15-af47-40a7-8c06-a3f43314d44f"),
	// 			Enabled: to.Ptr(true),
	// 			Attributes: map[string]any{
	// 				"deviceType": "sensor",
	// 				"deviceOwner": "IT",
	// 				"deviceCategory": 16,
	// 			},
	// 			ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
