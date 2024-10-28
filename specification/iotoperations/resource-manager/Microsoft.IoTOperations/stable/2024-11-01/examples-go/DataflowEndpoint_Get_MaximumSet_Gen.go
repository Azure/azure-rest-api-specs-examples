package armiotoperations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotoperations/armiotoperations"
)

// Generated from example definition: 2024-11-01/DataflowEndpoint_Get_MaximumSet_Gen.json
func ExampleDataflowEndpointClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataflowEndpointClient().Get(ctx, "rgiotoperations", "resource-name123", "resource-name123", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armiotoperations.DataflowEndpointClientGetResponse{
	// 	DataflowEndpointResource: &armiotoperations.DataflowEndpointResource{
	// 		Properties: &armiotoperations.DataflowEndpointProperties{
	// 			EndpointType: to.Ptr(armiotoperations.EndpointTypeDataExplorer),
	// 			DataExplorerSettings: &armiotoperations.DataflowEndpointDataExplorer{
	// 				Authentication: &armiotoperations.DataflowEndpointDataExplorerAuthentication{
	// 					Method: to.Ptr(armiotoperations.DataExplorerAuthMethodSystemAssignedManagedIdentity),
	// 					SystemAssignedManagedIdentitySettings: &armiotoperations.DataflowEndpointAuthenticationSystemAssignedManagedIdentity{
	// 						Audience: to.Ptr("psxomrfbhoflycm"),
	// 					},
	// 					UserAssignedManagedIdentitySettings: &armiotoperations.DataflowEndpointAuthenticationUserAssignedManagedIdentity{
	// 						ClientID: to.Ptr("fb90f267-8872-431a-a76a-a1cec5d3c4d2"),
	// 						Scope: to.Ptr("zop"),
	// 						TenantID: to.Ptr("ed060aa2-71ff-4d3f-99c4-a9138356fdec"),
	// 					},
	// 				},
	// 				Database: to.Ptr("yqcdpjsifm"),
	// 				Host: to.Ptr("<cluster>.<region>.kusto.windows.net"),
	// 				Batching: &armiotoperations.BatchingConfiguration{
	// 					LatencySeconds: to.Ptr[int32](1228),
	// 					MaxMessages: to.Ptr[int32](171),
	// 				},
	// 			},
	// 			DataLakeStorageSettings: &armiotoperations.DataflowEndpointDataLakeStorage{
	// 				Authentication: &armiotoperations.DataflowEndpointDataLakeStorageAuthentication{
	// 					Method: to.Ptr(armiotoperations.DataLakeStorageAuthMethodSystemAssignedManagedIdentity),
	// 					AccessTokenSettings: &armiotoperations.DataflowEndpointAuthenticationAccessToken{
	// 						SecretRef: to.Ptr("sevriyphcvnlrnfudqzejecwa"),
	// 					},
	// 					SystemAssignedManagedIdentitySettings: &armiotoperations.DataflowEndpointAuthenticationSystemAssignedManagedIdentity{
	// 						Audience: to.Ptr("psxomrfbhoflycm"),
	// 					},
	// 					UserAssignedManagedIdentitySettings: &armiotoperations.DataflowEndpointAuthenticationUserAssignedManagedIdentity{
	// 						ClientID: to.Ptr("fb90f267-8872-431a-a76a-a1cec5d3c4d2"),
	// 						Scope: to.Ptr("zop"),
	// 						TenantID: to.Ptr("ed060aa2-71ff-4d3f-99c4-a9138356fdec"),
	// 					},
	// 				},
	// 				Host: to.Ptr("<account>.blob.core.windows.net"),
	// 				Batching: &armiotoperations.BatchingConfiguration{
	// 					LatencySeconds: to.Ptr[int32](1228),
	// 					MaxMessages: to.Ptr[int32](171),
	// 				},
	// 			},
	// 			FabricOneLakeSettings: &armiotoperations.DataflowEndpointFabricOneLake{
	// 				Authentication: &armiotoperations.DataflowEndpointFabricOneLakeAuthentication{
	// 					Method: to.Ptr(armiotoperations.FabricOneLakeAuthMethodSystemAssignedManagedIdentity),
	// 					SystemAssignedManagedIdentitySettings: &armiotoperations.DataflowEndpointAuthenticationSystemAssignedManagedIdentity{
	// 						Audience: to.Ptr("psxomrfbhoflycm"),
	// 					},
	// 					UserAssignedManagedIdentitySettings: &armiotoperations.DataflowEndpointAuthenticationUserAssignedManagedIdentity{
	// 						ClientID: to.Ptr("fb90f267-8872-431a-a76a-a1cec5d3c4d2"),
	// 						Scope: to.Ptr("zop"),
	// 						TenantID: to.Ptr("ed060aa2-71ff-4d3f-99c4-a9138356fdec"),
	// 					},
	// 				},
	// 				Names: &armiotoperations.DataflowEndpointFabricOneLakeNames{
	// 					LakehouseName: to.Ptr("wpeathi"),
	// 					WorkspaceName: to.Ptr("nwgmitkbljztgms"),
	// 				},
	// 				OneLakePathType: to.Ptr(armiotoperations.DataflowEndpointFabricPathTypeFiles),
	// 				Host: to.Ptr("https://<host>.fabric.microsoft.com"),
	// 				Batching: &armiotoperations.BatchingConfiguration{
	// 					LatencySeconds: to.Ptr[int32](1228),
	// 					MaxMessages: to.Ptr[int32](171),
	// 				},
	// 			},
	// 			KafkaSettings: &armiotoperations.DataflowEndpointKafka{
	// 				Authentication: &armiotoperations.DataflowEndpointKafkaAuthentication{
	// 					Method: to.Ptr(armiotoperations.KafkaAuthMethodSystemAssignedManagedIdentity),
	// 					SystemAssignedManagedIdentitySettings: &armiotoperations.DataflowEndpointAuthenticationSystemAssignedManagedIdentity{
	// 						Audience: to.Ptr("psxomrfbhoflycm"),
	// 					},
	// 					UserAssignedManagedIdentitySettings: &armiotoperations.DataflowEndpointAuthenticationUserAssignedManagedIdentity{
	// 						ClientID: to.Ptr("fb90f267-8872-431a-a76a-a1cec5d3c4d2"),
	// 						Scope: to.Ptr("zop"),
	// 						TenantID: to.Ptr("ed060aa2-71ff-4d3f-99c4-a9138356fdec"),
	// 					},
	// 					SaslSettings: &armiotoperations.DataflowEndpointAuthenticationSasl{
	// 						SaslType: to.Ptr(armiotoperations.DataflowEndpointAuthenticationSaslTypePlain),
	// 						SecretRef: to.Ptr("visyxoztqnylvbyokhtmpdkwes"),
	// 					},
	// 					X509CertificateSettings: &armiotoperations.DataflowEndpointAuthenticationX509{
	// 						SecretRef: to.Ptr("afwizrystfslkfqd"),
	// 					},
	// 				},
	// 				ConsumerGroupID: to.Ptr("ukkzcjiyenhxokat"),
	// 				Host: to.Ptr("pwcqfiqclcgneolpewnyavoulbip"),
	// 				Batching: &armiotoperations.DataflowEndpointKafkaBatching{
	// 					Mode: to.Ptr(armiotoperations.OperationalModeEnabled),
	// 					LatencyMs: to.Ptr[int32](3679),
	// 					MaxBytes: to.Ptr[int32](8887),
	// 					MaxMessages: to.Ptr[int32](2174),
	// 				},
	// 				CopyMqttProperties: to.Ptr(armiotoperations.OperationalModeEnabled),
	// 				Compression: to.Ptr(armiotoperations.DataflowEndpointKafkaCompressionNone),
	// 				KafkaAcks: to.Ptr(armiotoperations.DataflowEndpointKafkaAcksZero),
	// 				PartitionStrategy: to.Ptr(armiotoperations.DataflowEndpointKafkaPartitionStrategyDefault),
	// 				TLS: &armiotoperations.TLSProperties{
	// 					Mode: to.Ptr(armiotoperations.OperationalModeEnabled),
	// 					TrustedCaCertificateConfigMapRef: to.Ptr("tectjjvukvelsreihwadh"),
	// 				},
	// 			},
	// 			LocalStorageSettings: &armiotoperations.DataflowEndpointLocalStorage{
	// 				PersistentVolumeClaimRef: to.Ptr("jjwqwvd"),
	// 			},
	// 			MqttSettings: &armiotoperations.DataflowEndpointMqtt{
	// 				Authentication: &armiotoperations.DataflowEndpointMqttAuthentication{
	// 					Method: to.Ptr(armiotoperations.MqttAuthMethodSystemAssignedManagedIdentity),
	// 					SystemAssignedManagedIdentitySettings: &armiotoperations.DataflowEndpointAuthenticationSystemAssignedManagedIdentity{
	// 						Audience: to.Ptr("psxomrfbhoflycm"),
	// 					},
	// 					UserAssignedManagedIdentitySettings: &armiotoperations.DataflowEndpointAuthenticationUserAssignedManagedIdentity{
	// 						ClientID: to.Ptr("fb90f267-8872-431a-a76a-a1cec5d3c4d2"),
	// 						Scope: to.Ptr("zop"),
	// 						TenantID: to.Ptr("ed060aa2-71ff-4d3f-99c4-a9138356fdec"),
	// 					},
	// 					ServiceAccountTokenSettings: &armiotoperations.DataflowEndpointAuthenticationServiceAccountToken{
	// 						Audience: to.Ptr("ejbklrbxgjaqleoycgpje"),
	// 					},
	// 					X509CertificateSettings: &armiotoperations.DataflowEndpointAuthenticationX509{
	// 						SecretRef: to.Ptr("afwizrystfslkfqd"),
	// 					},
	// 				},
	// 				ClientIDPrefix: to.Ptr("kkljsdxdirfhwxtkavldekeqhv"),
	// 				Host: to.Ptr("nyhnxqnbspstctl"),
	// 				Protocol: to.Ptr(armiotoperations.BrokerProtocolTypeMqtt),
	// 				KeepAliveSeconds: to.Ptr[int32](0),
	// 				Retain: to.Ptr(armiotoperations.MqttRetainTypeKeep),
	// 				MaxInflightMessages: to.Ptr[int32](0),
	// 				Qos: to.Ptr[int32](1),
	// 				SessionExpirySeconds: to.Ptr[int32](0),
	// 				TLS: &armiotoperations.TLSProperties{
	// 					Mode: to.Ptr(armiotoperations.OperationalModeEnabled),
	// 					TrustedCaCertificateConfigMapRef: to.Ptr("tectjjvukvelsreihwadh"),
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
