package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v11"
)

// Generated from example definition: 2018-06-01/Factories_CreateOrUpdate.json
func ExampleFactoriesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("12345678-1234-1234-1234-123456789012", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFactoriesClient().CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.Factory{
		Location: to.Ptr("East US"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatafactory.FactoriesClientCreateOrUpdateResponse{
	// 	Factory: armdatafactory.Factory{
	// 		Name: to.Ptr("exampleFactoryName"),
	// 		Type: to.Ptr("Microsoft.DataFactory/factories"),
	// 		ETag: to.Ptr("\"00003e04-0000-0000-0000-5b28979e0000\""),
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName"),
	// 		Location: to.Ptr("East US"),
	// 		Properties: &armdatafactory.FactoryProperties{
	// 			CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-19T05:41:50.0041314Z"); return t}()),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			Version: to.Ptr("2018-06-01"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
