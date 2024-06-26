package armhybriddatamanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/DataManagers_Get-GET-example-41.json
func ExampleDataManagersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybriddatamanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataManagersClient().Get(ctx, "ResourceGroupForSDKTest", "TestAzureSDKOperations", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataManager = armhybriddatamanager.DataManager{
	// 	Name: to.Ptr("TestAzureSDKOperations"),
	// 	Type: to.Ptr("Microsoft.HybridData/dataManagers"),
	// 	ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations"),
	// 	Location: to.Ptr("westus"),
	// 	SKU: &armhybriddatamanager.SKU{
	// 		Name: to.Ptr("DS0"),
	// 		Tier: to.Ptr("Standard"),
	// 	},
	// 	Tags: map[string]*string{
	// 	},
	// 	Etag: to.Ptr("W/\"datetime'2020-02-05T04%3A22%3A40.6354864Z'\"_W/\"datetime'2020-02-05T04%3A22%3A40.7912864Z'\""),
	// }
}
