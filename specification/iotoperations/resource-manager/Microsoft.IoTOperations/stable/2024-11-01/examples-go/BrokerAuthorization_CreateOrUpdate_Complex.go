package armiotoperations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotoperations/armiotoperations"
)

// Generated from example definition: 2024-11-01/BrokerAuthorization_CreateOrUpdate_Complex.json
func ExampleBrokerAuthorizationClient_BeginCreateOrUpdate_brokerAuthorizationCreateOrUpdateComplex() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBrokerAuthorizationClient().BeginCreateOrUpdate(ctx, "rgiotoperations", "resource-name123", "resource-name123", "resource-name123", armiotoperations.BrokerAuthorizationResource{
		Properties: &armiotoperations.BrokerAuthorizationProperties{
			AuthorizationPolicies: &armiotoperations.AuthorizationConfig{
				Cache: to.Ptr(armiotoperations.OperationalModeEnabled),
				Rules: []*armiotoperations.AuthorizationRule{
					{
						Principals: &armiotoperations.PrincipalDefinition{
							Usernames: []*string{
								to.Ptr("temperature-sensor"),
								to.Ptr("humidity-sensor"),
							},
							Attributes: []map[string]*string{
								{
									"building":     to.Ptr("17"),
									"organization": to.Ptr("contoso"),
								},
							},
						},
						BrokerResources: []*armiotoperations.BrokerResourceRule{
							{
								Method: to.Ptr(armiotoperations.BrokerResourceDefinitionMethodsConnect),
								ClientIDs: []*string{
									to.Ptr("{principal.attributes.building}*"),
								},
							},
							{
								Method: to.Ptr(armiotoperations.BrokerResourceDefinitionMethodsPublish),
								Topics: []*string{
									to.Ptr("sensors/{principal.attributes.building}/{principal.clientId}/telemetry/*"),
								},
							},
							{
								Method: to.Ptr(armiotoperations.BrokerResourceDefinitionMethodsSubscribe),
								Topics: []*string{
									to.Ptr("commands/{principal.attributes.organization}"),
								},
							},
						},
						StateStoreResources: []*armiotoperations.StateStoreResourceRule{
							{
								Method:  to.Ptr(armiotoperations.StateStoreResourceDefinitionMethodsRead),
								KeyType: to.Ptr(armiotoperations.StateStoreResourceKeyTypesPattern),
								Keys: []*string{
									to.Ptr("myreadkey"),
									to.Ptr("myotherkey?"),
									to.Ptr("mynumerickeysuffix[0-9]"),
									to.Ptr("clients:{principal.clientId}:*"),
								},
							},
							{
								Method:  to.Ptr(armiotoperations.StateStoreResourceDefinitionMethodsReadWrite),
								KeyType: to.Ptr(armiotoperations.StateStoreResourceKeyTypesBinary),
								Keys: []*string{
									to.Ptr("MTE2IDEwMSAxMTUgMTE2"),
								},
							},
						},
					},
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
	// res = armiotoperations.BrokerAuthorizationClientCreateOrUpdateResponse{
	// 	BrokerAuthorizationResource: &armiotoperations.BrokerAuthorizationResource{
	// 		Properties: &armiotoperations.BrokerAuthorizationProperties{
	// 			AuthorizationPolicies: &armiotoperations.AuthorizationConfig{
	// 				Cache: to.Ptr(armiotoperations.OperationalModeEnabled),
	// 				Rules: []*armiotoperations.AuthorizationRule{
	// 					{
	// 						Principals: &armiotoperations.PrincipalDefinition{
	// 							Usernames: []*string{
	// 								to.Ptr("temperature-sensor"),
	// 								to.Ptr("humidity-sensor"),
	// 							},
	// 							Attributes: []map[string]*string{
	// 								map[string]*string{
	// 									"building": to.Ptr("17"),
	// 									"organization": to.Ptr("contoso"),
	// 								},
	// 							},
	// 						},
	// 						BrokerResources: []*armiotoperations.BrokerResourceRule{
	// 							{
	// 								Method: to.Ptr(armiotoperations.BrokerResourceDefinitionMethodsConnect),
	// 								ClientIDs: []*string{
	// 									to.Ptr("{principal.attributes.building}*"),
	// 								},
	// 							},
	// 							{
	// 								Method: to.Ptr(armiotoperations.BrokerResourceDefinitionMethodsPublish),
	// 								Topics: []*string{
	// 									to.Ptr("sensors/{principal.attributes.building}/{principal.clientId}/telemetry/*"),
	// 								},
	// 							},
	// 							{
	// 								Method: to.Ptr(armiotoperations.BrokerResourceDefinitionMethodsSubscribe),
	// 								Topics: []*string{
	// 									to.Ptr("commands/{principal.attributes.organization}"),
	// 								},
	// 							},
	// 						},
	// 						StateStoreResources: []*armiotoperations.StateStoreResourceRule{
	// 							{
	// 								Method: to.Ptr(armiotoperations.StateStoreResourceDefinitionMethodsRead),
	// 								KeyType: to.Ptr(armiotoperations.StateStoreResourceKeyTypesPattern),
	// 								Keys: []*string{
	// 									to.Ptr("myreadkey"),
	// 									to.Ptr("myotherkey?"),
	// 									to.Ptr("mynumerickeysuffix[0-9]"),
	// 									to.Ptr("clients:{principal.clientId}:*"),
	// 								},
	// 							},
	// 							{
	// 								Method: to.Ptr(armiotoperations.StateStoreResourceDefinitionMethodsReadWrite),
	// 								KeyType: to.Ptr(armiotoperations.StateStoreResourceKeyTypesBinary),
	// 								Keys: []*string{
	// 									to.Ptr("MTE2IDEwMSAxMTUgMTE2"),
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armiotoperations.ProvisioningStateSucceeded),
	// 		},
	// 		ExtendedLocation: &armiotoperations.ExtendedLocation{
	// 			Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
	// 			Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
	// 		},
	// 		ID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.IoTOperations/instances/resource-name123/brokers/resource-name123/authorizations/resource-name123"),
	// 		Name: to.Ptr("anqrqsvrjmlvzkrbuav"),
	// 		Type: to.Ptr("yjlsfarshqoxojvgmy"),
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
