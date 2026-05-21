package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-01-20-preview/ParameterGroups_ListVersions.json
func ExampleParameterGroupsClient_NewListVersionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewParameterGroupsClient().NewListVersionsPager("exampleresourcegroup", "exampleparametergroup", &armhorizondb.ParameterGroupsClientListVersionsOptions{
		Version: to.Ptr[int32](22)})
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
		// page = armhorizondb.ParameterGroupsClientListVersionsResponse{
		// 	ParameterGroupListResult: armhorizondb.ParameterGroupListResult{
		// 		Value: []*armhorizondb.ParameterGroup{
		// 			{
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.HorizonDb/parameterGroups/exampleparametergroup"),
		// 				Name: to.Ptr("exampleparametergroup"),
		// 				Type: to.Ptr("Microsoft.HorizonDb/parameterGroups"),
		// 				Location: to.Ptr("westus2"),
		// 				Tags: map[string]*string{
		// 					"env": to.Ptr("dev"),
		// 					"team": to.Ptr("data-platform"),
		// 				},
		// 				Properties: &armhorizondb.ParameterGroupProperties{
		// 					Parameters: []*armhorizondb.ParameterProperties{
		// 						{
		// 							Name: to.Ptr("max_connections"),
		// 							Description: to.Ptr("Maximum number of concurrent connections to the database"),
		// 							Value: to.Ptr("200"),
		// 							DataType: to.Ptr("integer"),
		// 							AllowedValues: to.Ptr("100,200,300,400,500"),
		// 							IsDynamic: to.Ptr(true),
		// 							IsReadOnly: to.Ptr(false),
		// 							DocumentationLink: to.Ptr("https://docs.postgresql.org/current/runtime-config-connection.html#GUC-MAX-CONNECTIONS"),
		// 							Unit: to.Ptr("connections"),
		// 						},
		// 					},
		// 					Description: to.Ptr("Parameter group for high-throughput workloads"),
		// 					PgVersion: to.Ptr[int32](17),
		// 					Version: to.Ptr[int32](22),
		// 					ApplyImmediately: to.Ptr(true),
		// 					ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/anotherResourceGroup/providers/Microsoft.HorizonDb/parameterGroups/exampleparametergroup3"),
		// 				Name: to.Ptr("exampleparametergroup3"),
		// 				Type: to.Ptr("Microsoft.HorizonDb/parameterGroups"),
		// 				Location: to.Ptr("westus2"),
		// 				Tags: map[string]*string{
		// 					"env": to.Ptr("prod"),
		// 				},
		// 				Properties: &armhorizondb.ParameterGroupProperties{
		// 					Parameters: []*armhorizondb.ParameterProperties{
		// 						{
		// 							Name: to.Ptr("shared_buffers"),
		// 							Description: to.Ptr("Amount of memory used for shared memory buffers"),
		// 							Value: to.Ptr("3000"),
		// 							DataType: to.Ptr("integer"),
		// 							AllowedValues: to.Ptr("128,256,512,1024,2048,4096"),
		// 							IsDynamic: to.Ptr(false),
		// 							IsReadOnly: to.Ptr(false),
		// 							DocumentationLink: to.Ptr("https://docs.postgresql.org/current/runtime-config-resource.html#GUC-SHARED-BUFFERS"),
		// 							Unit: to.Ptr("MB"),
		// 						},
		// 					},
		// 					Description: to.Ptr("Parameter group for production workloads"),
		// 					PgVersion: to.Ptr[int32](17),
		// 					Version: to.Ptr[int32](22),
		// 					ApplyImmediately: to.Ptr(true),
		// 					ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
