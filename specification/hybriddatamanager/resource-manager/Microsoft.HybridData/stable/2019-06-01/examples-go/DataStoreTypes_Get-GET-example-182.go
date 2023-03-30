package armhybriddatamanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/DataStoreTypes_Get-GET-example-182.json
func ExampleDataStoreTypesClient_Get_dataStoreTypesGetGet182() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybriddatamanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataStoreTypesClient().Get(ctx, "StorSimple8000Series", "ResourceGroupForSDKTest", "TestAzureSDKOperations", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataStoreType = armhybriddatamanager.DataStoreType{
	// 	Name: to.Ptr("StorSimple8000Series"),
	// 	Type: to.Ptr("Microsoft.HybridData/dataManagers/dataStoreTypes"),
	// 	ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStoreTypes/StorSimple8000Series"),
	// 	Properties: &armhybriddatamanager.DataStoreTypeProperties{
	// 		RepositoryType: to.Ptr("Microsoft.StorSimple/managers"),
	// 		State: to.Ptr(armhybriddatamanager.StateEnabled),
	// 		SupportedDataServicesAsSink: []*string{
	// 		},
	// 		SupportedDataServicesAsSource: []*string{
	// 			to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataServices/DataTransformation"),
	// 			to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataServices/IndexingAndSearch")},
	// 		},
	// 	}
}
