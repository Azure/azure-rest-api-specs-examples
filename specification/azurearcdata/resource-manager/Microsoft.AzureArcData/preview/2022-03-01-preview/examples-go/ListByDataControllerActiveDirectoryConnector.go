package armazurearcdata_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/ListByDataControllerActiveDirectoryConnector.json
func ExampleActiveDirectoryConnectorsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurearcdata.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewActiveDirectoryConnectorsClient().NewListPager("testrg", "testdataController", nil)
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
		// page.ActiveDirectoryConnectorListResult = armazurearcdata.ActiveDirectoryConnectorListResult{
		// 	Value: []*armazurearcdata.ActiveDirectoryConnectorResource{
		// 		{
		// 			Name: to.Ptr("testADConnector1"),
		// 			Type: to.Ptr("Microsoft.AzureArcData/dataControllers/activeDirectoryConnectors"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureArcData/dataControllers/testdataController/activeDirectoryConnectors/testADConnector1"),
		// 			SystemData: &armazurearcdata.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
		// 			},
		// 			Properties: &armazurearcdata.ActiveDirectoryConnectorProperties{
		// 				Spec: &armazurearcdata.ActiveDirectoryConnectorSpec{
		// 					ActiveDirectory: &armazurearcdata.ActiveDirectoryConnectorDomainDetails{
		// 						DomainControllers: &armazurearcdata.ActiveDirectoryDomainControllers{
		// 							PrimaryDomainController: &armazurearcdata.ActiveDirectoryDomainController{
		// 								Hostname: to.Ptr("dc1.contoso.local"),
		// 							},
		// 							SecondaryDomainControllers: []*armazurearcdata.ActiveDirectoryDomainController{
		// 								{
		// 									Hostname: to.Ptr("dc2.contoso.local"),
		// 								},
		// 								{
		// 									Hostname: to.Ptr("dc3.contoso.local"),
		// 							}},
		// 						},
		// 						NetbiosDomainName: to.Ptr("CONTOSO"),
		// 						Realm: to.Ptr("CONTOSO.LOCAL"),
		// 						ServiceAccountProvisioning: to.Ptr(armazurearcdata.AccountProvisioningModeManual),
		// 					},
		// 					DNS: &armazurearcdata.ActiveDirectoryConnectorDNSDetails{
		// 						DomainName: to.Ptr("contoso.local"),
		// 						NameserverIPAddresses: []*string{
		// 							to.Ptr("11.11.111.111"),
		// 							to.Ptr("22.22.222.222")},
		// 							PreferK8SDNSForPtrLookups: to.Ptr(false),
		// 							Replicas: to.Ptr[int64](1),
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("testADConnector2"),
		// 				Type: to.Ptr("Microsoft.AzureArcData/dataControllers/activeDirectoryConnectors"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureArcData/dataControllers/testdataController/activeDirectoryConnectors/testADConnector2"),
		// 				SystemData: &armazurearcdata.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 					CreatedBy: to.Ptr("user1"),
		// 					CreatedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user2"),
		// 					LastModifiedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
		// 				},
		// 				Properties: &armazurearcdata.ActiveDirectoryConnectorProperties{
		// 					Spec: &armazurearcdata.ActiveDirectoryConnectorSpec{
		// 						ActiveDirectory: &armazurearcdata.ActiveDirectoryConnectorDomainDetails{
		// 							DomainControllers: &armazurearcdata.ActiveDirectoryDomainControllers{
		// 								PrimaryDomainController: &armazurearcdata.ActiveDirectoryDomainController{
		// 									Hostname: to.Ptr("dc4.contoso.local"),
		// 								},
		// 								SecondaryDomainControllers: []*armazurearcdata.ActiveDirectoryDomainController{
		// 									{
		// 										Hostname: to.Ptr("dc5.contoso.local"),
		// 									},
		// 									{
		// 										Hostname: to.Ptr("dc6.contoso.local"),
		// 								}},
		// 							},
		// 							NetbiosDomainName: to.Ptr("CONTOSO"),
		// 							Realm: to.Ptr("CONTOSO.LOCAL"),
		// 							ServiceAccountProvisioning: to.Ptr(armazurearcdata.AccountProvisioningModeManual),
		// 						},
		// 						DNS: &armazurearcdata.ActiveDirectoryConnectorDNSDetails{
		// 							DomainName: to.Ptr("contoso.local"),
		// 							NameserverIPAddresses: []*string{
		// 								to.Ptr("11.11.111.111"),
		// 								to.Ptr("22.22.222.222")},
		// 								PreferK8SDNSForPtrLookups: to.Ptr(false),
		// 								Replicas: to.Ptr[int64](1),
		// 							},
		// 						},
		// 					},
		// 			}},
		// 		}
	}
}
