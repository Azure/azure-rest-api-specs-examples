package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/58be094c6b365f8d4d73a91e293dfb4818e57cf6/specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/ContainerApps_ListBySubscription.json
func ExampleContainerAppsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewContainerAppsClient().NewListBySubscriptionPager(nil)
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
		// page.ContainerAppCollection = armappcontainers.ContainerAppCollection{
		// 	Value: []*armappcontainers.ContainerApp{
		// 		{
		// 			Name: to.Ptr("testcontainerApp0"),
		// 			Type: to.Ptr("Microsoft.App/containerApps"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/containerApps/testcontainerApp0"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappcontainers.ContainerAppProperties{
		// 				Configuration: &armappcontainers.Configuration{
		// 					Dapr: &armappcontainers.Dapr{
		// 						AppPort: to.Ptr[int32](3000),
		// 						AppProtocol: to.Ptr(armappcontainers.AppProtocolHTTP),
		// 						EnableAPILogging: to.Ptr(true),
		// 						Enabled: to.Ptr(true),
		// 						HTTPMaxRequestSize: to.Ptr[int32](10),
		// 						HTTPReadBufferSize: to.Ptr[int32](30),
		// 						LogLevel: to.Ptr(armappcontainers.LogLevelDebug),
		// 					},
		// 					Ingress: &armappcontainers.Ingress{
		// 						CustomDomains: []*armappcontainers.CustomDomain{
		// 							{
		// 								Name: to.Ptr("www.my-name.com"),
		// 								BindingType: to.Ptr(armappcontainers.BindingTypeSniEnabled),
		// 								CertificateID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-name-dot-com"),
		// 							},
		// 							{
		// 								Name: to.Ptr("www.my--other-name.com"),
		// 								BindingType: to.Ptr(armappcontainers.BindingTypeSniEnabled),
		// 								CertificateID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-other-name-dot-com"),
		// 						}},
		// 						External: to.Ptr(true),
		// 						Fqdn: to.Ptr("testcontainerApp0.demokube-t24clv0g.eastus.containerApps.k4apps.io"),
		// 						IPSecurityRestrictions: []*armappcontainers.IPSecurityRestrictionRule{
		// 							{
		// 								Name: to.Ptr("Allow work IP A subnet"),
		// 								Description: to.Ptr("Allowing all IP's within the subnet below to access containerapp"),
		// 								Action: to.Ptr(armappcontainers.ActionAllow),
		// 								IPAddressRange: to.Ptr("192.168.1.1/32"),
		// 							},
		// 							{
		// 								Name: to.Ptr("Allow work IP B subnet"),
		// 								Description: to.Ptr("Allowing all IP's within the subnet below to access containerapp"),
		// 								Action: to.Ptr(armappcontainers.ActionAllow),
		// 								IPAddressRange: to.Ptr("192.168.1.1/8"),
		// 						}},
		// 						StickySessions: &armappcontainers.IngressStickySessions{
		// 							Affinity: to.Ptr(armappcontainers.AffinitySticky),
		// 						},
		// 						TargetPort: to.Ptr[int32](3000),
		// 						Traffic: []*armappcontainers.TrafficWeight{
		// 							{
		// 								RevisionName: to.Ptr("testcontainerApp0-ab1234"),
		// 								Weight: to.Ptr[int32](80),
		// 							},
		// 							{
		// 								Label: to.Ptr("staging"),
		// 								RevisionName: to.Ptr("testcontainerApp0-ab4321"),
		// 								Weight: to.Ptr[int32](20),
		// 						}},
		// 						Transport: to.Ptr(armappcontainers.IngressTransportMethodAuto),
		// 					},
		// 					MaxInactiveRevisions: to.Ptr[int32](10),
		// 					Service: &armappcontainers.Service{
		// 						Type: to.Ptr("redis"),
		// 					},
		// 				},
		// 				EventStreamEndpoint: to.Ptr("testEndpoint"),
		// 				LatestReadyRevisionName: to.Ptr("testcontainerApp0-pjxhsye"),
		// 				LatestRevisionFqdn: to.Ptr("testcontainerApp0-pjxhsye.demokube-t24clv0g.eastus.containerApps.k4apps.io"),
		// 				ManagedEnvironmentID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube"),
		// 				ProvisioningState: to.Ptr(armappcontainers.ContainerAppProvisioningStateSucceeded),
		// 				Template: &armappcontainers.Template{
		// 					Containers: []*armappcontainers.Container{
		// 						{
		// 							Name: to.Ptr("testcontainerApp0"),
		// 							Image: to.Ptr("repo/testcontainerApp0:v4"),
		// 							Resources: &armappcontainers.ContainerResources{
		// 								CPU: to.Ptr[float64](0.2),
		// 								Memory: to.Ptr("100Mi"),
		// 							},
		// 					}},
		// 					InitContainers: []*armappcontainers.InitContainer{
		// 						{
		// 							Name: to.Ptr("testinitcontainerApp0"),
		// 							Image: to.Ptr("repo/testcontainerApp0:v4"),
		// 							Resources: &armappcontainers.ContainerResources{
		// 								CPU: to.Ptr[float64](0.2),
		// 								Memory: to.Ptr("100Mi"),
		// 							},
		// 					}},
		// 					Scale: &armappcontainers.Scale{
		// 						MaxReplicas: to.Ptr[int32](5),
		// 						MinReplicas: to.Ptr[int32](1),
		// 						Rules: []*armappcontainers.ScaleRule{
		// 							{
		// 								Name: to.Ptr("httpscalingrule"),
		// 								HTTP: &armappcontainers.HTTPScaleRule{
		// 									Metadata: map[string]*string{
		// 										"concurrentRequests": to.Ptr("50"),
		// 									},
		// 								},
		// 						}},
		// 					},
		// 					ServiceBinds: []*armappcontainers.ServiceBind{
		// 						{
		// 							Name: to.Ptr("service"),
		// 							ServiceID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/containerApps/service"),
		// 					}},
		// 				},
		// 				WorkloadProfileName: to.Ptr("My-GP-01"),
		// 			},
		// 	}},
		// }
	}
}
