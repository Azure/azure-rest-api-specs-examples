package armappconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4a2bb0762eaad11e725516708483598e0c12cabb/specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2024-06-01/examples/ConfigurationStoresListReplicas.json
func ExampleReplicasClient_NewListByConfigurationStorePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappconfiguration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReplicasClient().NewListByConfigurationStorePager("myResourceGroup", "contoso", &armappconfiguration.ReplicasClientListByConfigurationStoreOptions{SkipToken: nil})
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
		// page.ReplicaListResult = armappconfiguration.ReplicaListResult{
		// 	Value: []*armappconfiguration.Replica{
		// 		{
		// 			Name: to.Ptr("myReplicaEus"),
		// 			Type: to.Ptr("Microsoft.AppConfiguration/configurationStores/replicas"),
		// 			ID: to.Ptr("/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourceGroups/myResourceGroup/providers/Microsoft.AppConfiguration/configurationStores/contoso/replicas/myReplicaEus"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armappconfiguration.ReplicaProperties{
		// 				Endpoint: to.Ptr("https://contoso-myreplicaeus.azconfig.io"),
		// 				ProvisioningState: to.Ptr(armappconfiguration.ReplicaProvisioningStateSucceeded),
		// 			},
		// 			SystemData: &armappconfiguration.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("foo@contoso.com"),
		// 				CreatedByType: to.Ptr(armappconfiguration.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("foo@contoso.com"),
		// 				LastModifiedByType: to.Ptr(armappconfiguration.CreatedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myReplicaWestEurope"),
		// 			Type: to.Ptr("Microsoft.AppConfiguration/configurationStores/replicas"),
		// 			ID: to.Ptr("/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourceGroups/myResourceGroup/providers/Microsoft.AppConfiguration/configurationStores/contoso/replicas/myReplicaWestEurope"),
		// 			Location: to.Ptr("westeurope"),
		// 			Properties: &armappconfiguration.ReplicaProperties{
		// 				Endpoint: to.Ptr("https://contoso-myreplicawesteurope.azconfig.io"),
		// 				ProvisioningState: to.Ptr(armappconfiguration.ReplicaProvisioningStateSucceeded),
		// 			},
		// 			SystemData: &armappconfiguration.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("foo@contoso.com"),
		// 				CreatedByType: to.Ptr(armappconfiguration.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("foo@contoso.com"),
		// 				LastModifiedByType: to.Ptr(armappconfiguration.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
