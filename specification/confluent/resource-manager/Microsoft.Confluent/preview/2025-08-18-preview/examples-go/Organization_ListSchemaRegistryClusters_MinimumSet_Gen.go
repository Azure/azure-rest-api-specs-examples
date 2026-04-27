package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent/v2"
)

// Generated from example definition: 2025-08-18-preview/Organization_ListSchemaRegistryClusters_MinimumSet_Gen.json
func ExampleOrganizationClient_NewListSchemaRegistryClustersPager_organizationListSchemaRegistryClustersMinimumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("DC34558A-05D3-4370-AED8-75E60B381F94", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOrganizationClient().NewListSchemaRegistryClustersPager("rgconfluent", "npek", "tdtxr", nil)
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
		// page = armconfluent.OrganizationClientListSchemaRegistryClustersResponse{
		// 	ListSchemaRegistryClustersResponse: armconfluent.ListSchemaRegistryClustersResponse{
		// 		Value: []*armconfluent.SchemaRegistryClusterRecord{
		// 			{
		// 			},
		// 		},
		// 	},
		// }
	}
}
