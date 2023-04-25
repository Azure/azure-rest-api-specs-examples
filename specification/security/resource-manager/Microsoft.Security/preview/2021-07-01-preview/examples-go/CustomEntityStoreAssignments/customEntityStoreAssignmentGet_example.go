package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e716082ac474f182e2220e4f38f1d6191e7636cf/specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/CustomEntityStoreAssignments/customEntityStoreAssignmentGet_example.json
func ExampleCustomEntityStoreAssignmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCustomEntityStoreAssignmentsClient().Get(ctx, "TestResourceGroup", "33e7cc6e-a139-4723-a0e5-76993aee0771", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomEntityStoreAssignment = armsecurity.CustomEntityStoreAssignment{
	// 	Name: to.Ptr("33e7cc6e-a139-4723-a0e5-76993aee0771"),
	// 	Type: to.Ptr("Microsoft.Security/customEntityStoreAssignments"),
	// 	ID: to.Ptr("/subscriptions/e5d1b86c-3051-44d5-8802-aa65d45a279b/resourcegroups/TestResourceGroup/providers/Microsoft.Security/customEntityStoreAssignments/33e7cc6e-a139-4723-a0e5-76993aee0771"),
	// 	Properties: &armsecurity.CustomEntityStoreAssignmentProperties{
	// 		EntityStoreDatabaseLink: to.Ptr("https://dataexplorer.azure.com/clusters/securitydatastore.centralus/databases/DiscoveryAwsKedamari?query=H4sIAAAAAAAAAwtILC4uzy9KCcjPyUyu5OWqUShJzE5VMAQAlMJzABgAAAA="),
	// 		Principal: to.Ptr("aaduser=f3923a3e-ad57-4752-b1a9-fbf3c8e5e082;72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 	},
	// 	SystemData: &armsecurity.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@contoso.com"),
	// 		CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@contoso.com"),
	// 		LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
	// 	},
	// }
}
