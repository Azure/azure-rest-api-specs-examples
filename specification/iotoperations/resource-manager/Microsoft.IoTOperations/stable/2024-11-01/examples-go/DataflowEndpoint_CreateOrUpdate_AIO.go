package armiotoperations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotoperations/armiotoperations"
)

// Generated from example definition: 2024-11-01/DataflowEndpoint_CreateOrUpdate_AIO.json
func ExampleDataflowEndpointClient_BeginCreateOrUpdate_dataflowEndpointCreateOrUpdateAio() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDataflowEndpointClient().BeginCreateOrUpdate(ctx, "rgiotoperations", "resource-name123", "aio-builtin-broker-endpoint", armiotoperations.DataflowEndpointResource{
		Properties: &armiotoperations.DataflowEndpointProperties{
			EndpointType: to.Ptr(armiotoperations.EndpointTypeMqtt),
			MqttSettings: &armiotoperations.DataflowEndpointMqtt{
				Host: to.Ptr("aio-broker:18883"),
				Authentication: &armiotoperations.DataflowEndpointMqttAuthentication{
					ServiceAccountTokenSettings: &armiotoperations.DataflowEndpointAuthenticationServiceAccountToken{
						Audience: to.Ptr("aio-internal"),
					},
				},
				TLS: &armiotoperations.TLSProperties{
					Mode:                             to.Ptr(armiotoperations.OperationalModeEnabled),
					TrustedCaCertificateConfigMapRef: to.Ptr("aio-ca-trust-bundle-test-only"),
				},
			},
		},
		ExtendedLocation: &armiotoperations.ExtendedLocation{
			Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
			Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
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
	// res = armiotoperations.DataflowEndpointClientCreateOrUpdateResponse{
	// 	DataflowEndpointResource: &armiotoperations.DataflowEndpointResource{
	// 		Properties: &armiotoperations.DataflowEndpointProperties{
	// 			EndpointType: to.Ptr(armiotoperations.EndpointTypeMqtt),
	// 			MqttSettings: &armiotoperations.DataflowEndpointMqtt{
	// 				Host: to.Ptr("aio-broker:18883"),
	// 				Authentication: &armiotoperations.DataflowEndpointMqttAuthentication{
	// 					ServiceAccountTokenSettings: &armiotoperations.DataflowEndpointAuthenticationServiceAccountToken{
	// 						Audience: to.Ptr("aio-internal"),
	// 					},
	// 				},
	// 				TLS: &armiotoperations.TLSProperties{
	// 					Mode: to.Ptr(armiotoperations.OperationalModeEnabled),
	// 					TrustedCaCertificateConfigMapRef: to.Ptr("aio-ca-trust-bundle-test-only"),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armiotoperations.ProvisioningStateSucceeded),
	// 		},
	// 		ExtendedLocation: &armiotoperations.ExtendedLocation{
	// 			Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
	// 			Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
	// 		},
	// 		ID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.IoTOperations/instances/resource-name123/dataflowEndpoints/resource-name123"),
	// 		Name: to.Ptr("zyhxscudobzfacetvgyjiav"),
	// 		Type: to.Ptr("iay"),
	// 		SystemData: &armiotoperations.SystemData{
	// 			CreatedBy: to.Ptr("ssvaslsmudloholronopqyxjcu"),
	// 			CreatedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("gnicpuszwd"),
	// 			LastModifiedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 		},
	// 	},
	// }
}
