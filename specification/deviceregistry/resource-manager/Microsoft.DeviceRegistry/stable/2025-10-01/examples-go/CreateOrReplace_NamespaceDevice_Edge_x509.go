package armdeviceregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceregistry/armdeviceregistry/v2"
)

// Generated from example definition: 2025-10-01/CreateOrReplace_NamespaceDevice_Edge_x509.json
func ExampleNamespaceDevicesClient_BeginCreateOrReplace_createEdgeEnabledDeviceWithX509InboundAuthentication() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNamespaceDevicesClient().BeginCreateOrReplace(ctx, "myResourceGroup", "adr-namespace-gbk0925-n01", "namespace-device-on-edge", armdeviceregistry.NamespaceDevice{
		ExtendedLocation: &armdeviceregistry.ExtendedLocation{
			Type: to.Ptr("CustomLocation"),
			Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
		},
		Location: to.Ptr("West Europe"),
		Properties: &armdeviceregistry.NamespaceDeviceProperties{
			Endpoints: &armdeviceregistry.MessagingEndpoints{
				Inbound: map[string]*armdeviceregistry.InboundEndpoints{
					"theV1OPCUAEndpoint": {
						Address:      to.Ptr("opc.tcp://192.168.86.23:51211/UA/SampleServer"),
						EndpointType: to.Ptr("microsoft.opcua"),
						Version:      to.Ptr("2"),
						Authentication: &armdeviceregistry.HostAuthentication{
							Method: to.Ptr(armdeviceregistry.AuthenticationMethodCertificate),
							X509Credentials: &armdeviceregistry.X509CertificateCredentials{
								CertificateSecretName:              to.Ptr("cert-secret"),
								KeySecretName:                      to.Ptr("key-secret"),
								IntermediateCertificatesSecretName: to.Ptr("intermediate-certs-secret"),
							},
						},
					},
					"theV2OPCUAEndpoint": {
						Address:      to.Ptr("opc.tcp://192.168.86.23:51211/UA/SampleServer"),
						EndpointType: to.Ptr("microsoft.opcua"),
						Version:      to.Ptr("2"),
						Authentication: &armdeviceregistry.HostAuthentication{
							Method: to.Ptr(armdeviceregistry.AuthenticationMethodCertificate),
							X509Credentials: &armdeviceregistry.X509CertificateCredentials{
								CertificateSecretName: to.Ptr("cert-secret"),
							},
						},
						TrustSettings: &armdeviceregistry.TrustSettings{
							TrustList: to.Ptr("trust-secret-reference"),
						},
					},
				},
			},
			ExternalDeviceID: to.Ptr("unique-edge-device-identifier"),
			Enabled:          to.Ptr(true),
			Attributes: map[string]any{
				"deviceType":     "OPCUAServers",
				"deviceOwner":    "OT",
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
	// 		Name: to.Ptr("namespace-device-on-edge"),
	// 		Type: to.Ptr("private.deviceregistry/namespaces/devices"),
	// 		Location: to.Ptr("West Europe"),
	// 		Tags: map[string]*string{
	// 			"sensor": to.Ptr("temperature,humidity,rotation"),
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
	// 			UUID: to.Ptr("3575e690-31d7-4168-a6c9-9ff1b9eed282"),
	// 			Endpoints: &armdeviceregistry.MessagingEndpoints{
	// 				Inbound: map[string]*armdeviceregistry.InboundEndpoints{
	// 					"theV1OPCUAEndpoint": &armdeviceregistry.InboundEndpoints{
	// 						Address: to.Ptr("opc.tcp://192.168.86.23:51211/UA/SampleServer"),
	// 						EndpointType: to.Ptr("microsoft.opcua"),
	// 						Version: to.Ptr("2"),
	// 						Authentication: &armdeviceregistry.HostAuthentication{
	// 							Method: to.Ptr(armdeviceregistry.AuthenticationMethodCertificate),
	// 							X509Credentials: &armdeviceregistry.X509CertificateCredentials{
	// 								CertificateSecretName: to.Ptr("cert-secret"),
	// 								KeySecretName: to.Ptr("key-secret"),
	// 								IntermediateCertificatesSecretName: to.Ptr("intermediate-certs-secret"),
	// 							},
	// 						},
	// 					},
	// 					"theV2OPCUAEndpoint": &armdeviceregistry.InboundEndpoints{
	// 						Address: to.Ptr("opc.tcp://192.168.86.23:51211/UA/SampleServer"),
	// 						EndpointType: to.Ptr("microsoft.opcua"),
	// 						Version: to.Ptr("2"),
	// 						Authentication: &armdeviceregistry.HostAuthentication{
	// 							Method: to.Ptr(armdeviceregistry.AuthenticationMethodCertificate),
	// 							X509Credentials: &armdeviceregistry.X509CertificateCredentials{
	// 								CertificateSecretName: to.Ptr("cert-secret"),
	// 							},
	// 						},
	// 						TrustSettings: &armdeviceregistry.TrustSettings{
	// 							TrustList: to.Ptr("trust-secret-reference"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ExternalDeviceID: to.Ptr("unique-edge-device-identifier"),
	// 			Enabled: to.Ptr(true),
	// 			Attributes: map[string]any{
	// 				"deviceType": "OPCUAServers",
	// 				"deviceOwner": "OT",
	// 			},
	// 			ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
