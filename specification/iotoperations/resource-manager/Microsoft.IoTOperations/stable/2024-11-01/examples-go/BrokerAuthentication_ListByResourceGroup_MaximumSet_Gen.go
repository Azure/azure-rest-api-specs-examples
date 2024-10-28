package armiotoperations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotoperations/armiotoperations"
)

// Generated from example definition: 2024-11-01/BrokerAuthentication_ListByResourceGroup_MaximumSet_Gen.json
func ExampleBrokerAuthenticationClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBrokerAuthenticationClient().NewListByResourceGroupPager("rgiotoperations", "resource-name123", "resource-name123", nil)
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
		// page = armiotoperations.BrokerAuthenticationClientListByResourceGroupResponse{
		// 	BrokerAuthenticationResourceListResult: armiotoperations.BrokerAuthenticationResourceListResult{
		// 		Value: []*armiotoperations.BrokerAuthenticationResource{
		// 			{
		// 				Properties: &armiotoperations.BrokerAuthenticationProperties{
		// 					AuthenticationMethods: []*armiotoperations.BrokerAuthenticatorMethods{
		// 						{
		// 							Method: to.Ptr(armiotoperations.BrokerAuthenticationMethodCustom),
		// 							CustomSettings: &armiotoperations.BrokerAuthenticatorMethodCustom{
		// 								Auth: &armiotoperations.BrokerAuthenticatorCustomAuth{
		// 									X509: &armiotoperations.X509ManualCertificate{
		// 										SecretRef: to.Ptr("secret-name"),
		// 									},
		// 								},
		// 								CaCertConfigMap: to.Ptr("pdecudefqyolvncbus"),
		// 								Endpoint: to.Ptr("https://www.example.com"),
		// 								Headers: map[string]*string{
		// 									"key8518": to.Ptr("bwityjy"),
		// 								},
		// 							},
		// 							ServiceAccountTokenSettings: &armiotoperations.BrokerAuthenticatorMethodSat{
		// 								Audiences: []*string{
		// 									to.Ptr("jqyhyqatuydg"),
		// 								},
		// 							},
		// 							X509Settings: &armiotoperations.BrokerAuthenticatorMethodX509{
		// 								AuthorizationAttributes: map[string]*armiotoperations.BrokerAuthenticatorMethodX509Attributes{
		// 									"key3384": &armiotoperations.BrokerAuthenticatorMethodX509Attributes{
		// 										Attributes: map[string]*string{
		// 											"key186": to.Ptr("ucpajramsz"),
		// 										},
		// 										Subject: to.Ptr("jpgwctfeixitptfgfnqhua"),
		// 									},
		// 								},
		// 								TrustedClientCaCert: to.Ptr("vlctsqddl"),
		// 							},
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armiotoperations.ProvisioningStateSucceeded),
		// 				},
		// 				ExtendedLocation: &armiotoperations.ExtendedLocation{
		// 					Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
		// 					Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
		// 				},
		// 				ID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.IoTOperations/instances/resource-name123/brokers/resource-name123/authentications/resource-name123"),
		// 				Name: to.Ptr("lwucizfvtsdpx"),
		// 				Type: to.Ptr("kvtilkgcxanlfozrd"),
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
