package armiotoperations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotoperations/armiotoperations"
)

// Generated from example definition: 2024-11-01/Broker_ListByResourceGroup_MaximumSet_Gen.json
func ExampleBrokerClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBrokerClient().NewListByResourceGroupPager("rgiotoperations", "resource-name123", nil)
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
		// page = armiotoperations.BrokerClientListByResourceGroupResponse{
		// 	BrokerResourceListResult: armiotoperations.BrokerResourceListResult{
		// 		Value: []*armiotoperations.BrokerResource{
		// 			{
		// 				Properties: &armiotoperations.BrokerProperties{
		// 					Advanced: &armiotoperations.AdvancedSettings{
		// 						Clients: &armiotoperations.ClientConfig{
		// 							MaxSessionExpirySeconds: to.Ptr[int32](3859),
		// 							MaxMessageExpirySeconds: to.Ptr[int32](3263),
		// 							MaxPacketSizeBytes: to.Ptr[int32](3029),
		// 							SubscriberQueueLimit: &armiotoperations.SubscriberQueueLimit{
		// 								Length: to.Ptr[int64](6),
		// 								Strategy: to.Ptr(armiotoperations.SubscriberMessageDropStrategyNone),
		// 							},
		// 							MaxReceiveMaximum: to.Ptr[int32](2365),
		// 							MaxKeepAliveSeconds: to.Ptr[int32](3744),
		// 						},
		// 						EncryptInternalTraffic: to.Ptr(armiotoperations.OperationalModeEnabled),
		// 						InternalCerts: &armiotoperations.CertManagerCertOptions{
		// 							Duration: to.Ptr("bchrc"),
		// 							RenewBefore: to.Ptr("xkafmpgjfifkwwrhkswtopdnne"),
		// 							PrivateKey: &armiotoperations.CertManagerPrivateKey{
		// 								Algorithm: to.Ptr(armiotoperations.PrivateKeyAlgorithmEc256),
		// 								RotationPolicy: to.Ptr(armiotoperations.PrivateKeyRotationPolicyAlways),
		// 							},
		// 						},
		// 					},
		// 					Cardinality: &armiotoperations.Cardinality{
		// 						BackendChain: &armiotoperations.BackendChain{
		// 							Partitions: to.Ptr[int32](11),
		// 							RedundancyFactor: to.Ptr[int32](5),
		// 							Workers: to.Ptr[int32](15),
		// 						},
		// 						Frontend: &armiotoperations.Frontend{
		// 							Replicas: to.Ptr[int32](2),
		// 							Workers: to.Ptr[int32](6),
		// 						},
		// 					},
		// 					Diagnostics: &armiotoperations.BrokerDiagnostics{
		// 						Logs: &armiotoperations.DiagnosticsLogs{
		// 							Level: to.Ptr("rnmwokumdmebpmfxxxzvvjfdywotav"),
		// 						},
		// 						Metrics: &armiotoperations.Metrics{
		// 							PrometheusPort: to.Ptr[int32](7581),
		// 						},
		// 						SelfCheck: &armiotoperations.SelfCheck{
		// 							Mode: to.Ptr(armiotoperations.OperationalModeEnabled),
		// 							IntervalSeconds: to.Ptr[int32](158),
		// 							TimeoutSeconds: to.Ptr[int32](14),
		// 						},
		// 						Traces: &armiotoperations.Traces{
		// 							Mode: to.Ptr(armiotoperations.OperationalModeEnabled),
		// 							CacheSizeMegabytes: to.Ptr[int32](28),
		// 							SelfTracing: &armiotoperations.SelfTracing{
		// 								Mode: to.Ptr(armiotoperations.OperationalModeEnabled),
		// 								IntervalSeconds: to.Ptr[int32](22),
		// 							},
		// 							SpanChannelCapacity: to.Ptr[int32](1000),
		// 						},
		// 					},
		// 					DiskBackedMessageBuffer: &armiotoperations.DiskBackedMessageBuffer{
		// 						MaxSize: to.Ptr("500M"),
		// 						EphemeralVolumeClaimSpec: &armiotoperations.VolumeClaimSpec{
		// 							VolumeName: to.Ptr("c"),
		// 							VolumeMode: to.Ptr("rxvpksjuuugqnqzeiprocknbn"),
		// 							StorageClassName: to.Ptr("sseyhrjptkhrqvpdpjmornkqvon"),
		// 							AccessModes: []*string{
		// 								to.Ptr("nuluhigrbb"),
		// 							},
		// 							DataSource: &armiotoperations.LocalKubernetesReference{
		// 								APIGroup: to.Ptr("npqapyksvvpkohujx"),
		// 								Kind: to.Ptr("wazgyb"),
		// 								Name: to.Ptr("cwhsgxxcxsyppoefm"),
		// 							},
		// 							DataSourceRef: &armiotoperations.KubernetesReference{
		// 								APIGroup: to.Ptr("mnfnykznjjsoqpfsgdqioupt"),
		// 								Kind: to.Ptr("odynqzekfzsnawrctaxg"),
		// 								Name: to.Ptr("envszivbbmixbyddzg"),
		// 								Namespace: to.Ptr("etcfzvxqd"),
		// 							},
		// 							Resources: &armiotoperations.VolumeClaimResourceRequirements{
		// 								Limits: map[string]*string{
		// 									"key2719": to.Ptr("hmphcrgctu"),
		// 								},
		// 								Requests: map[string]*string{
		// 									"key2909": to.Ptr("txocprnyrsgvhfrg"),
		// 								},
		// 							},
		// 							Selector: &armiotoperations.VolumeClaimSpecSelector{
		// 								MatchExpressions: []*armiotoperations.VolumeClaimSpecSelectorMatchExpressions{
		// 									{
		// 										Key: to.Ptr("e"),
		// 										Operator: to.Ptr(armiotoperations.OperatorValuesIn),
		// 										Values: []*string{
		// 											to.Ptr("slmpajlywqvuyknipgztsonqyybt"),
		// 										},
		// 									},
		// 								},
		// 								MatchLabels: map[string]*string{
		// 									"key6673": to.Ptr("wlngfalznwxnurzpgxomcxhbqefpr"),
		// 								},
		// 							},
		// 						},
		// 						PersistentVolumeClaimSpec: &armiotoperations.VolumeClaimSpec{
		// 							VolumeName: to.Ptr("c"),
		// 							VolumeMode: to.Ptr("rxvpksjuuugqnqzeiprocknbn"),
		// 							StorageClassName: to.Ptr("sseyhrjptkhrqvpdpjmornkqvon"),
		// 							AccessModes: []*string{
		// 								to.Ptr("nuluhigrbb"),
		// 							},
		// 							DataSource: &armiotoperations.LocalKubernetesReference{
		// 								APIGroup: to.Ptr("npqapyksvvpkohujx"),
		// 								Kind: to.Ptr("wazgyb"),
		// 								Name: to.Ptr("cwhsgxxcxsyppoefm"),
		// 							},
		// 							DataSourceRef: &armiotoperations.KubernetesReference{
		// 								APIGroup: to.Ptr("mnfnykznjjsoqpfsgdqioupt"),
		// 								Kind: to.Ptr("odynqzekfzsnawrctaxg"),
		// 								Name: to.Ptr("envszivbbmixbyddzg"),
		// 								Namespace: to.Ptr("etcfzvxqd"),
		// 							},
		// 							Resources: &armiotoperations.VolumeClaimResourceRequirements{
		// 								Limits: map[string]*string{
		// 									"key2719": to.Ptr("hmphcrgctu"),
		// 								},
		// 								Requests: map[string]*string{
		// 									"key2909": to.Ptr("txocprnyrsgvhfrg"),
		// 								},
		// 							},
		// 							Selector: &armiotoperations.VolumeClaimSpecSelector{
		// 								MatchExpressions: []*armiotoperations.VolumeClaimSpecSelectorMatchExpressions{
		// 									{
		// 										Key: to.Ptr("e"),
		// 										Operator: to.Ptr(armiotoperations.OperatorValuesIn),
		// 										Values: []*string{
		// 											to.Ptr("slmpajlywqvuyknipgztsonqyybt"),
		// 										},
		// 									},
		// 								},
		// 								MatchLabels: map[string]*string{
		// 									"key6673": to.Ptr("wlngfalznwxnurzpgxomcxhbqefpr"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					GenerateResourceLimits: &armiotoperations.GenerateResourceLimits{
		// 						CPU: to.Ptr(armiotoperations.OperationalModeEnabled),
		// 					},
		// 					MemoryProfile: to.Ptr(armiotoperations.BrokerMemoryProfileTiny),
		// 					ProvisioningState: to.Ptr(armiotoperations.ProvisioningStateSucceeded),
		// 				},
		// 				ExtendedLocation: &armiotoperations.ExtendedLocation{
		// 					Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
		// 					Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
		// 				},
		// 				ID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.IoTOperations/instances/resource-name123/brokers/resource-name123"),
		// 				Name: to.Ptr("dowrkel"),
		// 				Type: to.Ptr("xshjnsdgadygb"),
		// 				SystemData: &armiotoperations.SystemData{
		// 					CreatedBy: to.Ptr("ssvaslsmudloholronopqyxjcu"),
		// 					CreatedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("gnicpuszwd"),
		// 					LastModifiedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
