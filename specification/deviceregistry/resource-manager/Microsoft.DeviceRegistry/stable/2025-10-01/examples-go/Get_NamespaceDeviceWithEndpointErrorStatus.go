package armdeviceregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceregistry/armdeviceregistry/v2"
)

// Generated from example definition: 2025-10-01/Get_NamespaceDeviceWithEndpointErrorStatus.json
func ExampleNamespaceDevicesClient_Get_getNamespaceDeviceWithEndpointErrorStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNamespaceDevicesClient().Get(ctx, "myResourceGroup", "my-namespace-1", "my-device-name", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdeviceregistry.NamespaceDevicesClientGetResponse{
	// 	NamespaceDevice: &armdeviceregistry.NamespaceDevice{
	// 		ID: to.Ptr("/subscriptions/00000000-1366-430f-0000-cc873bcf2d27/resourceGroups/gbktestRG/providers/Microsoft.DeviceRegistry/namespaces/my-namespace-1/devices/adr-smart-device3-f191f536-f652-4eb4-b9a0-1a9d43300cab"),
	// 		Name: to.Ptr("adr-smart-device3-f191f536-f652-4eb4-b9a0-1a9d43300cab"),
	// 		Type: to.Ptr("Microsoft.DeviceRegistry/namespaces/devices"),
	// 		Location: to.Ptr("North Europe"),
	// 		ExtendedLocation: &armdeviceregistry.ExtendedLocation{
	// 			Type: to.Ptr("CustomLocation"),
	// 			Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
	// 		},
	// 		Etag: to.Ptr("\"c80325f9-0000-0800-0000-66fcb0290000\""),
	// 		Tags: map[string]*string{
	// 			"sensor": to.Ptr("temperature,humidity"),
	// 		},
	// 		SystemData: &armdeviceregistry.SystemData{
	// 			CreatedBy: to.Ptr("00003442-0000-0000-0000-494059220000"),
	// 			CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-02T02:30:01.6394987Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("00003442-0000-0000-0000-494059220000"),
	// 			LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-02T02:30:01.6394987Z"); return t}()),
	// 		},
	// 		Properties: &armdeviceregistry.NamespaceDeviceProperties{
	// 			Endpoints: &armdeviceregistry.MessagingEndpoints{
	// 				Inbound: map[string]*armdeviceregistry.InboundEndpoints{
	// 					"myOPCUABroker": &armdeviceregistry.InboundEndpoints{
	// 						Address: to.Ptr("opc.tcp://foo"),
	// 						EndpointType: to.Ptr("microsoft.opcua"),
	// 						Version: to.Ptr("2"),
	// 						Authentication: &armdeviceregistry.HostAuthentication{
	// 							Method: to.Ptr(armdeviceregistry.AuthenticationMethodCertificate),
	// 							X509Credentials: &armdeviceregistry.X509CertificateCredentials{
	// 								CertificateSecretName: to.Ptr("certificateSecretName"),
	// 								KeySecretName: to.Ptr("keySecretName"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			Version: to.Ptr[int64](1),
	// 			Status: &armdeviceregistry.DeviceStatus{
	// 				Config: &armdeviceregistry.StatusConfig{
	// 					Version: to.Ptr[int64](2),
	// 					LastTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-02T02:30:01.6394987Z"); return t}()),
	// 				},
	// 				Endpoints: &armdeviceregistry.DeviceStatusEndpoints{
	// 					Inbound: map[string]*armdeviceregistry.DeviceStatusEndpoint{
	// 						"myOPCUABroker": &armdeviceregistry.DeviceStatusEndpoint{
	// 							Error: &armdeviceregistry.StatusError{
	// 								Code: to.Ptr("InvalidCertificate"),
	// 								Message: to.Ptr("The client certificate is invalid."),
	// 								Details: []*armdeviceregistry.ErrorDetails{
	// 									{
	// 										Code: to.Ptr("InvalidCertificate"),
	// 										Message: to.Ptr("The certificate is invalid."),
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
