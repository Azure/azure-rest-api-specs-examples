package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-01-20-preview/ParameterGroups_ListBySubscription.json
func ExampleParameterGroupsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewParameterGroupsClient().NewListBySubscriptionPager(nil)
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
		// page = armhorizondb.ParameterGroupsClientListBySubscriptionResponse{
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
		// 							Description: to.Ptr("Sets the maximum number of concurrent connections"),
		// 							Value: to.Ptr("200"),
		// 							DataType: to.Ptr("integer"),
		// 							AllowedValues: to.Ptr("100,200,500,1000"),
		// 							IsDynamic: to.Ptr(true),
		// 							IsReadOnly: to.Ptr(false),
		// 							DocumentationLink: to.Ptr("https://docs.example.com/max_connections"),
		// 							Unit: to.Ptr("connections"),
		// 						},
		// 						{
		// 							Name: to.Ptr("log_min_error_statement"),
		// 							Description: to.Ptr("Controls which statements are logged"),
		// 							Value: to.Ptr("error"),
		// 							DataType: to.Ptr("string"),
		// 							AllowedValues: to.Ptr("debug,info,notice,warning,error,log,fatal,panic"),
		// 							IsDynamic: to.Ptr(true),
		// 							IsReadOnly: to.Ptr(false),
		// 							DocumentationLink: to.Ptr("https://docs.example.com/log_min_error_statement"),
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
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/anotherResourceGroup/providers/Microsoft.HorizonDb/parameterGroups/exampleparametergroup2"),
		// 				Name: to.Ptr("exampleparametergroup2"),
		// 				Type: to.Ptr("Microsoft.HorizonDb/parameterGroups"),
		// 				Location: to.Ptr("westus2"),
		// 				Tags: map[string]*string{
		// 					"env": to.Ptr("prod"),
		// 					"team": to.Ptr("data-platform"),
		// 				},
		// 				Properties: &armhorizondb.ParameterGroupProperties{
		// 					Parameters: []*armhorizondb.ParameterProperties{
		// 						{
		// 							Name: to.Ptr("max_connections"),
		// 							Description: to.Ptr("Sets the maximum number of concurrent connections"),
		// 							Value: to.Ptr("500"),
		// 							DataType: to.Ptr("integer"),
		// 							AllowedValues: to.Ptr("100,200,500,1000"),
		// 							IsDynamic: to.Ptr(true),
		// 							IsReadOnly: to.Ptr(false),
		// 							DocumentationLink: to.Ptr("https://docs.example.com/max_connections"),
		// 							Unit: to.Ptr("connections"),
		// 						},
		// 					},
		// 					Description: to.Ptr("Parameter group for production workloads"),
		// 					PgVersion: to.Ptr[int32](17),
		// 					Version: to.Ptr[int32](25),
		// 					ApplyImmediately: to.Ptr(true),
		// 					ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testResourceGroup/providers/Microsoft.HorizonDb/parameterGroups/exampleparametergrouptest"),
		// 				Name: to.Ptr("exampleparametergrouptest"),
		// 				Type: to.Ptr("Microsoft.HorizonDb/parameterGroups"),
		// 				Location: to.Ptr("centralus"),
		// 				Tags: map[string]*string{
		// 					"env": to.Ptr("test"),
		// 				},
		// 				Properties: &armhorizondb.ParameterGroupProperties{
		// 					Parameters: []*armhorizondb.ParameterProperties{
		// 						{
		// 							Name: to.Ptr("shared_buffers"),
		// 							Description: to.Ptr("Sets the amount of memory the database server uses for shared memory buffers"),
		// 							Value: to.Ptr("1000"),
		// 							DataType: to.Ptr("integer"),
		// 							AllowedValues: to.Ptr("1000,2000,4000,8000"),
		// 							IsDynamic: to.Ptr(false),
		// 							IsReadOnly: to.Ptr(false),
		// 							DocumentationLink: to.Ptr("https://docs.example.com/shared_buffers"),
		// 							Unit: to.Ptr("MB"),
		// 						},
		// 					},
		// 					Description: to.Ptr("Parameter group for testing environments"),
		// 					PgVersion: to.Ptr[int32](17),
		// 					Version: to.Ptr[int32](10),
		// 					ApplyImmediately: to.Ptr(false),
		// 					ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
