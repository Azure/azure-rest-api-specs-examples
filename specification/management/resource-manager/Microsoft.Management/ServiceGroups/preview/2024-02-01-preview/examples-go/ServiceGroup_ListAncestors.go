package armservicegroups_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicegroups/armservicegroups"
)

// Generated from example definition: 2024-02-01-preview/ServiceGroup_ListAncestors.json
func ExampleClient_ListAncestors() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicegroups.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().ListAncestors(ctx, "20000000-0001-0000-0000-000000000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armservicegroups.ClientListAncestorsResponse{
	// 	ServiceGroupCollectionResponse: &armservicegroups.ServiceGroupCollectionResponse{
	// 		NextLink: to.Ptr("https://abc.def.com/providers/Microsoft.Management/serviceGroups/20000000-0001-0000-0000-000000000000/listAncestors?api-version=2024-02-01-preview&skipToken=xyz"),
	// 		Value: []*armservicegroups.ServiceGroup{
	// 			{
	// 				Name: to.Ptr("Ancestor1"),
	// 				Type: to.Ptr("Microsoft.Management/serviceGroups"),
	// 				ID: to.Ptr("/providers/Microsoft.Management/serviceGroups/Ancestor1"),
	// 				Properties: &armservicegroups.ServiceGroupProperties{
	// 					DisplayName: to.Ptr("ServiceGroup Ancestor1"),
	// 				},
	// 			},
	// 			{
	// 				Name: to.Ptr("Ancestor2"),
	// 				Type: to.Ptr("Microsoft.Management/serviceGroups"),
	// 				ID: to.Ptr("/providers/Microsoft.Management/serviceGroups/Ancestor2"),
	// 				Properties: &armservicegroups.ServiceGroupProperties{
	// 					DisplayName: to.Ptr("ServiceGroup Ancestor2"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
