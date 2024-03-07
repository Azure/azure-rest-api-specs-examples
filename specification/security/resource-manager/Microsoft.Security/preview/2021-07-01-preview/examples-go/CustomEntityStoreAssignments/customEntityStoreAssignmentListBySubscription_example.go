package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/CustomEntityStoreAssignments/customEntityStoreAssignmentListBySubscription_example.json
func ExampleCustomEntityStoreAssignmentsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomEntityStoreAssignmentsClient().NewListBySubscriptionPager(nil)
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
		// page.CustomEntityStoreAssignmentsListResult = armsecurity.CustomEntityStoreAssignmentsListResult{
		// 	Value: []*armsecurity.CustomEntityStoreAssignment{
		// 		{
		// 			Name: to.Ptr("33e7cc6e-a139-4723-a0e5-76993aee0771"),
		// 			Type: to.Ptr("Microsoft.Security/customEntityStoreAssignments"),
		// 			ID: to.Ptr("/subscriptions/e5d1b86c-3051-44d5-8802-aa65d45a279b/resourcegroups/TestResourceGroup/providers/Microsoft.Security/customEntityStoreAssignments/33e7cc6e-a139-4723-a0e5-76993aee0771"),
		// 			Properties: &armsecurity.CustomEntityStoreAssignmentProperties{
		// 				EntityStoreDatabaseLink: to.Ptr("https://dataexplorer.azure.com/clusters/securitydatastore.centralus/databases/DiscoveryAwsKedamari?query=H4sIAAAAAAAAAwtILC4uzy9KCcjPyUyu5OWqUShJzE5VMAQAlMJzABgAAAA="),
		// 				Principal: to.Ptr("aaduser=f3923a3e-ad57-4752-b1a9-fbf3c8e5e082;72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 			},
		// 			SystemData: &armsecurity.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@contoso.com"),
		// 				CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@contoso.com"),
		// 				LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("a400695c-4728-d5cc-8e19-4b5a76f209df"),
		// 			Type: to.Ptr("Microsoft.Security/customEntityStoreAssignments"),
		// 			ID: to.Ptr("/subscriptions/e5d1b86c-3051-44d5-8802-aa65d45a279b/resourcegroups/TestResourceGroup/providers/Microsoft.Security/customEntityStoreAssignments/a400695c-4728-d5cc-8e19-4b5a76f209df"),
		// 			Properties: &armsecurity.CustomEntityStoreAssignmentProperties{
		// 				EntityStoreDatabaseLink: to.Ptr("https://dataexplorer.azure.com/clusters/securitydatastore.centralus/databases/DiscoveryAwsKedamari?query=H4sIAAAAAAAAAwvIz8lMrgzKz0nlqlFIrShJzUtR8Cz2SE3MKcmoVLBVUE9LzClOVQcA1IFnficAAAA="),
		// 				Principal: to.Ptr("aaduser=f6e2564c-f34a-9b61-416c-5e5e7e521118;72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 			},
		// 			SystemData: &armsecurity.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:01:50.328Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@contoso.com"),
		// 				CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:01:50.328Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@contoso.com"),
		// 				LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
