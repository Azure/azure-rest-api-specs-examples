package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/be0d1c383d683a8eb22ab5fd5b75e380ac3c2eea/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_GetConnectionInfo.json
func ExampleIntegrationRuntimesClient_GetConnectionInfo() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIntegrationRuntimesClient().GetConnectionInfo(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleIntegrationRuntime", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IntegrationRuntimeConnectionInfo = armdatafactory.IntegrationRuntimeConnectionInfo{
	// 	HostServiceURI: to.Ptr("https://yanzhang-dt.fareast.corp.microsoft.com:8050/HostServiceRemote.svc/"),
	// 	IdentityCertThumbprint: to.Ptr("**********"),
	// 	IsIdentityCertExprired: to.Ptr(false),
	// 	PublicKey: to.Ptr("**********"),
	// 	ServiceToken: to.Ptr("**********"),
	// 	Version: to.Ptr("3.8.6730.2"),
	// }
}
