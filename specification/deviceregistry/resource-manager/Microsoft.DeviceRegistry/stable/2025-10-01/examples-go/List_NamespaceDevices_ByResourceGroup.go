package armdeviceregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceregistry/armdeviceregistry/v2"
)

// Generated from example definition: 2025-10-01/List_NamespaceDevices_ByResourceGroup.json
func ExampleNamespaceDevicesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNamespaceDevicesClient().NewListByResourceGroupPager("myResourceGroup", "adr-namespace-gbk0925-n01", nil)
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
		// page = armdeviceregistry.NamespaceDevicesClientListByResourceGroupResponse{
		// 	NamespaceDeviceListResult: armdeviceregistry.NamespaceDeviceListResult{
		// 		Value: []*armdeviceregistry.NamespaceDevice{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-1366-430f-0000-cc873bcf2d27/resourceGroups/gbktestRG/providers/Microsoft.DeviceRegistry/namespaces/my-namespace-1/devices/adr-smart-device3-f191f536-f652-4eb4-b9a0-1a9d43300cab"),
		// 				Name: to.Ptr("adr-smart-device3-f191f536-f652-4eb4-b9a0-1a9d43300cab"),
		// 				Type: to.Ptr("Microsoft.DeviceRegistry/namespaces/devices"),
		// 				Location: to.Ptr("North Europe"),
		// 				Etag: to.Ptr("\"c80325f9-0000-0800-0000-66fcb0290000\""),
		// 				Tags: map[string]*string{
		// 					"sensor": to.Ptr("temperature,humidity"),
		// 				},
		// 				SystemData: &armdeviceregistry.SystemData{
		// 					CreatedBy: to.Ptr("00003442-0000-0000-0000-494059220000"),
		// 					CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-02T02:30:01.6394987Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("00003442-0000-0000-0000-494059220000"),
		// 					LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-02T02:30:01.6394987Z"); return t}()),
		// 				},
		// 				Properties: &armdeviceregistry.NamespaceDeviceProperties{
		// 					ExternalDeviceID: to.Ptr("adr-smart-device3-f191f536-f652-4eb4-b9a0-1a9d43300cab"),
		// 					Enabled: to.Ptr(true),
		// 					Attributes: map[string]any{
		// 						"deviceType": "sensor",
		// 						"deviceOwner": "IT",
		// 						"deviceCategory": 16,
		// 					},
		// 					ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-1366-430f-0000-cc873bcf2d27/resourceGroups/gbktestRG/providers/Microsoft.DeviceRegistry/namespaces/my-namespace-1/devices/adr-smart-device3-67b8b41c-25e9-479a-884e-5502a5e9afbd"),
		// 				Name: to.Ptr("adr-smart-device3-67b8b41c-25e9-479a-884e-5502a5e9afbd"),
		// 				Type: to.Ptr("Microsoft.DeviceRegistry/namespaces/devices"),
		// 				Location: to.Ptr("North Europe"),
		// 				Etag: to.Ptr("\"e103c7dc-0000-0800-0000-66fda99c0000\""),
		// 				Tags: map[string]*string{
		// 					"sensor": to.Ptr("temperature,humidity"),
		// 				},
		// 				SystemData: &armdeviceregistry.SystemData{
		// 					CreatedBy: to.Ptr("00003442-0000-0000-0000-494059220000"),
		// 					CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-02T20:14:20.0968397Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("00003442-0000-0000-0000-494059220000"),
		// 					LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-02T20:14:20.0968397Z"); return t}()),
		// 				},
		// 				Properties: &armdeviceregistry.NamespaceDeviceProperties{
		// 					Endpoints: &armdeviceregistry.MessagingEndpoints{
		// 						Outbound: &armdeviceregistry.OutboundEndpoints{
		// 							Assigned: map[string]*armdeviceregistry.DeviceMessagingEndpoint{
		// 								"eventGridEndpoint": &armdeviceregistry.DeviceMessagingEndpoint{
		// 									EndpointType: to.Ptr("Microsoft.Devices/IoTHubs"),
		// 									Address: to.Ptr("https://myeventgridtopic.westeurope-1.eventgrid.azure.net/api/events"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					ExternalDeviceID: to.Ptr("adr-smart-device3-67b8b41c-25e9-479a-884e-5502a5e9afbd"),
		// 					Enabled: to.Ptr(true),
		// 					Attributes: map[string]any{
		// 						"deviceType": "sensor",
		// 						"deviceOwner": "IT",
		// 						"deviceCategory": 16,
		// 					},
		// 					ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
