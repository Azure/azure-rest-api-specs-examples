package armappconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration/v3"
)

// Generated from example definition: 2025-06-01-preview/ConfigurationStoresGetReplica.json
func ExampleReplicasClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappconfiguration.NewClientFactory("c80fb759-c965-4c6a-9110-9b2b2d038882", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReplicasClient().Get(ctx, "myResourceGroup", "contoso", "myReplicaEus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armappconfiguration.ReplicasClientGetResponse{
	// 	Replica: &armappconfiguration.Replica{
	// 		Name: to.Ptr("myReplicaEus"),
	// 		Type: to.Ptr("Microsoft.AppConfiguration/configurationStores/replicas"),
	// 		ID: to.Ptr("/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourceGroups/myResourceGroup/providers/Microsoft.AppConfiguration/configurationStores/contoso/replicas/myReplicaEus"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armappconfiguration.ReplicaProperties{
	// 			Endpoint: to.Ptr("https://contoso-myreplicaeus.azconfig.io"),
	// 			ProvisioningState: to.Ptr(armappconfiguration.ReplicaProvisioningStateSucceeded),
	// 		},
	// 		SystemData: &armappconfiguration.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55+00:00"); return t}()),
	// 			CreatedBy: to.Ptr("foo@contoso.com"),
	// 			CreatedByType: to.Ptr(armappconfiguration.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55+00:00"); return t}()),
	// 			LastModifiedBy: to.Ptr("foo@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armappconfiguration.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
