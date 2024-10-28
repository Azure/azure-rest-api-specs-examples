package armiotoperations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotoperations/armiotoperations"
)

// Generated from example definition: 2024-11-01/DataflowEndpoint_CreateOrUpdate_Kafka.json
func ExampleDataflowEndpointClient_BeginCreateOrUpdate_dataflowEndpointCreateOrUpdateKafka() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDataflowEndpointClient().BeginCreateOrUpdate(ctx, "rgiotoperations", "resource-name123", "generic-kafka-endpoint", armiotoperations.DataflowEndpointResource{
		Properties: &armiotoperations.DataflowEndpointProperties{
			EndpointType: to.Ptr(armiotoperations.EndpointTypeKafka),
			KafkaSettings: &armiotoperations.DataflowEndpointKafka{
				Host: to.Ptr("example.kafka.local:9093"),
				Authentication: &armiotoperations.DataflowEndpointKafkaAuthentication{
					Method: to.Ptr(armiotoperations.KafkaAuthMethodSasl),
					SaslSettings: &armiotoperations.DataflowEndpointAuthenticationSasl{
						SaslType:  to.Ptr(armiotoperations.DataflowEndpointAuthenticationSaslTypePlain),
						SecretRef: to.Ptr("my-secret"),
					},
				},
				TLS: &armiotoperations.TLSProperties{
					Mode:                             to.Ptr(armiotoperations.OperationalModeEnabled),
					TrustedCaCertificateConfigMapRef: to.Ptr("ca-certificates"),
				},
				ConsumerGroupID: to.Ptr("dataflows"),
				Compression:     to.Ptr(armiotoperations.DataflowEndpointKafkaCompressionGzip),
				Batching: &armiotoperations.DataflowEndpointKafkaBatching{
					Mode:        to.Ptr(armiotoperations.OperationalModeEnabled),
					LatencyMs:   to.Ptr[int32](5),
					MaxBytes:    to.Ptr[int32](1000000),
					MaxMessages: to.Ptr[int32](100000),
				},
				PartitionStrategy:    to.Ptr(armiotoperations.DataflowEndpointKafkaPartitionStrategyDefault),
				KafkaAcks:            to.Ptr(armiotoperations.DataflowEndpointKafkaAcksAll),
				CopyMqttProperties:   to.Ptr(armiotoperations.OperationalModeEnabled),
				CloudEventAttributes: to.Ptr(armiotoperations.CloudEventAttributeTypePropagate),
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
	// 			EndpointType: to.Ptr(armiotoperations.EndpointTypeKafka),
	// 			KafkaSettings: &armiotoperations.DataflowEndpointKafka{
	// 				Host: to.Ptr("example.kafka.local:9093"),
	// 				Authentication: &armiotoperations.DataflowEndpointKafkaAuthentication{
	// 					Method: to.Ptr(armiotoperations.KafkaAuthMethodSasl),
	// 					SaslSettings: &armiotoperations.DataflowEndpointAuthenticationSasl{
	// 						SaslType: to.Ptr(armiotoperations.DataflowEndpointAuthenticationSaslTypePlain),
	// 						SecretRef: to.Ptr("my-secret"),
	// 					},
	// 				},
	// 				TLS: &armiotoperations.TLSProperties{
	// 					Mode: to.Ptr(armiotoperations.OperationalModeEnabled),
	// 					TrustedCaCertificateConfigMapRef: to.Ptr("ca-certificates"),
	// 				},
	// 				ConsumerGroupID: to.Ptr("dataflows"),
	// 				Compression: to.Ptr(armiotoperations.DataflowEndpointKafkaCompressionGzip),
	// 				Batching: &armiotoperations.DataflowEndpointKafkaBatching{
	// 					Mode: to.Ptr(armiotoperations.OperationalModeEnabled),
	// 					LatencyMs: to.Ptr[int32](5),
	// 					MaxBytes: to.Ptr[int32](1000000),
	// 					MaxMessages: to.Ptr[int32](100000),
	// 				},
	// 				PartitionStrategy: to.Ptr(armiotoperations.DataflowEndpointKafkaPartitionStrategyDefault),
	// 				KafkaAcks: to.Ptr(armiotoperations.DataflowEndpointKafkaAcksAll),
	// 				CopyMqttProperties: to.Ptr(armiotoperations.OperationalModeEnabled),
	// 				CloudEventAttributes: to.Ptr(armiotoperations.CloudEventAttributeTypePropagate),
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
