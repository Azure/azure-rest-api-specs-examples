package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-01-20-preview/ParameterGroups_Get.json
func ExampleParameterGroupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewParameterGroupsClient().Get(ctx, "exampleresourcegroup", "exampleparametergroup", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhorizondb.ParameterGroupsClientGetResponse{
	// 	ParameterGroup: &armhorizondb.ParameterGroup{
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.HorizonDb/parameterGroups/exampleparametergroup"),
	// 		Name: to.Ptr("exampleparametergroup"),
	// 		Type: to.Ptr("Microsoft.HorizonDb/parameterGroups"),
	// 		Location: to.Ptr("westus2"),
	// 		Tags: map[string]*string{
	// 			"env": to.Ptr("dev"),
	// 			"team": to.Ptr("data-platform"),
	// 		},
	// 		Properties: &armhorizondb.ParameterGroupProperties{
	// 			Parameters: []*armhorizondb.ParameterProperties{
	// 				{
	// 					Name: to.Ptr("max_connections"),
	// 					Description: to.Ptr("Sets the maximum number of concurrent connections"),
	// 					Value: to.Ptr("200"),
	// 					DataType: to.Ptr("integer"),
	// 					AllowedValues: to.Ptr("100,200,500,1000"),
	// 					IsDynamic: to.Ptr(true),
	// 					IsReadOnly: to.Ptr(false),
	// 					DocumentationLink: to.Ptr("https://docs.example.com/max_connections"),
	// 					Unit: to.Ptr("connections"),
	// 				},
	// 				{
	// 					Name: to.Ptr("log_min_error_statement"),
	// 					Description: to.Ptr("Controls which statements are logged"),
	// 					Value: to.Ptr("error"),
	// 					DataType: to.Ptr("string"),
	// 					AllowedValues: to.Ptr("debug,info,notice,warning,error,log,fatal,panic"),
	// 					IsDynamic: to.Ptr(true),
	// 					IsReadOnly: to.Ptr(false),
	// 					DocumentationLink: to.Ptr("https://docs.example.com/log_min_error_statement"),
	// 				},
	// 				{
	// 					Name: to.Ptr("shared_buffers"),
	// 					Description: to.Ptr("Sets the amount of memory the database server uses for shared memory buffers"),
	// 					Value: to.Ptr("2000"),
	// 					DataType: to.Ptr("integer"),
	// 					AllowedValues: to.Ptr("1000,2000,4000,8000"),
	// 					IsDynamic: to.Ptr(false),
	// 					IsReadOnly: to.Ptr(false),
	// 					DocumentationLink: to.Ptr("https://docs.example.com/shared_buffers"),
	// 					Unit: to.Ptr("MB"),
	// 				},
	// 			},
	// 			Description: to.Ptr("Parameter group for high-throughput workloads"),
	// 			PgVersion: to.Ptr[int32](17),
	// 			Version: to.Ptr[int32](22),
	// 			ApplyImmediately: to.Ptr(true),
	// 			ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
