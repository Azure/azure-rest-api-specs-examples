package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01a71545e82bb98b8137d3038150c436d46a59ed/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Factories_Update.json
func ExampleFactoriesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFactoriesClient().Update(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.FactoryUpdateParameters{
		Tags: map[string]*string{
			"exampleTag": to.Ptr("exampleValue"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Factory = armdatafactory.Factory{
	// 	Name: to.Ptr("exampleFactoryName"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories"),
	// 	ETag: to.Ptr("\"00003f04-0000-0000-0000-5b28979e0000\""),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 		"exampleTag": to.Ptr("exampleValue"),
	// 	},
	// 	Properties: &armdatafactory.FactoryProperties{
	// 		CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-19T05:41:50.004Z"); return t}()),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Version: to.Ptr("2018-06-01"),
	// 	},
	// }
}
