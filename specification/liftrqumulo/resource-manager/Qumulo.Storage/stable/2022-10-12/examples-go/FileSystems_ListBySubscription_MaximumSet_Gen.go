package armqumulo_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/liftrqumulo/armqumulo"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2022-10-12/examples/FileSystems_ListBySubscription_MaximumSet_Gen.json
func ExampleFileSystemsClient_NewListBySubscriptionPager_fileSystemsListBySubscriptionMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armqumulo.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFileSystemsClient().NewListBySubscriptionPager(nil)
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
		// page.FileSystemResourceListResult = armqumulo.FileSystemResourceListResult{
		// 	Value: []*armqumulo.FileSystemResource{
		// 		{
		// 			Name: to.Ptr("bii"),
		// 			Type: to.Ptr("qtvxrqwpoistduq"),
		// 			ID: to.Ptr("tvelgpobdtazrweunifqzaxkgjauyx"),
		// 			SystemData: &armqumulo.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-14T04:40:17.991Z"); return t}()),
		// 				CreatedBy: to.Ptr("mtdhqooysjhueaojwpmvophkgntl"),
		// 				CreatedByType: to.Ptr(armqumulo.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-14T04:40:17.991Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("jcywglomzuamsxltnoegdrkzlscxl"),
		// 				LastModifiedByType: to.Ptr(armqumulo.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("przdlsmlzsszphnixq"),
		// 			Tags: map[string]*string{
		// 				"key6565": to.Ptr("cgdhmupta"),
		// 			},
		// 			Identity: &armqumulo.ManagedServiceIdentity{
		// 				Type: to.Ptr(armqumulo.ManagedServiceIdentityTypeNone),
		// 				PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 				TenantID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 				UserAssignedIdentities: map[string]*armqumulo.UserAssignedIdentity{
		// 					"key4522": &armqumulo.UserAssignedIdentity{
		// 						ClientID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 						PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 					},
		// 				},
		// 			},
		// 			Properties: &armqumulo.FileSystemResourceProperties{
		// 				AvailabilityZone: to.Ptr("maseyqhlnhoiwbabcqabtedbjpip"),
		// 				ClusterLoginURL: to.Ptr("jjqhgevy"),
		// 				DelegatedSubnetID: to.Ptr("neqctctqdmjezfgt"),
		// 				InitialCapacity: to.Ptr[int32](9),
		// 				MarketplaceDetails: &armqumulo.MarketplaceDetails{
		// 					MarketplaceSubscriptionID: to.Ptr("ujrcqvxfnhxxheoth"),
		// 					MarketplaceSubscriptionStatus: to.Ptr(armqumulo.MarketplaceSubscriptionStatusPendingFulfillmentStart),
		// 					OfferID: to.Ptr("eiyhbmpwgezcmzrrfoiskuxlcvwojf"),
		// 					PlanID: to.Ptr("x"),
		// 					PublisherID: to.Ptr("wfmokfdjbwpjhz"),
		// 				},
		// 				PrivateIPs: []*string{
		// 					to.Ptr("kslguxrwbwkrj")},
		// 					ProvisioningState: to.Ptr(armqumulo.ProvisioningStateSucceeded),
		// 					StorageSKU: to.Ptr(armqumulo.StorageSKUStandard),
		// 					UserDetails: &armqumulo.UserDetails{
		// 					},
		// 				},
		// 		}},
		// 	}
	}
}
