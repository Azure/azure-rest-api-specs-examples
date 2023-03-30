package armhybriddatamanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/DataStores_CreateOrUpdate_DataSource-PUT-example-162.json
func ExampleDataStoresClient_BeginCreateOrUpdate_dataStoresCreateOrUpdateDataSourcePut162() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybriddatamanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDataStoresClient().BeginCreateOrUpdate(ctx, "TestStorSimpleSource1", "ResourceGroupForSDKTest", "TestAzureSDKOperations", armhybriddatamanager.DataStore{
		Properties: &armhybriddatamanager.DataStoreProperties{
			CustomerSecrets: []*armhybriddatamanager.CustomerSecret{
				{
					Algorithm:     to.Ptr(armhybriddatamanager.SupportedAlgorithmRSA15),
					KeyIdentifier: to.Ptr("ServiceEncryptionKey"),
					KeyValue:      to.Ptr("EVuEBV40qv23xDRL4NZBuMms4e3So6ikHjrQYRvG9NloqxdgPOg+ZYzpho5lytI4fmv0ANmRIvDiDboRXcUVSjbB9T2gm19fMIuwZa4FK2+LYEgMqKK1GaLkk7xC8f5IeFUXLo6KyBBpaAiayTnWDcHuYEpMiGrV7trDDcbhMRefO3CHecmH3Z7ye8L0RQ/e7WW8GlCKZj3m0BaG7OrJgjai8gyDfMfGAG5rTqEhDVh2hLQ+TjvUjcOFwHvJusqKTENtbJTNQYmL9wZXsnwBvUwxqrGieILNB7V3GD1Ow9OiV0UCDW1e9LnMueukg+l7YJCU9FUhIPh/nSif6p32zw==:jCfio+pDtY3BSPZDpDJ0L6QdXLYMeOmxaFWtYTOZkNqNTgT8Loc/KSQRjtWS5K4N4btbznuSJ/dzg0aZEzlXgKDSuZgMfd4Ch92ZwAC/BkeCmVrTjiKJsoQXO1IICCUf7GHGBbYnnpsNJcEn4vyc9NXyKwOBjeU+I9AyK7PtYiC03RLpTS6xttFCICteBV0uoBHAiV0chZ5VIIUUMjO9u8EhHqRY7NNcGbWdVJeAb6J3vH4E/DHkQj+DXlpjcLvmK/uqBwxfNju30RJhR04Nmz6zcv/zTcvS0uN5hEPQoHLyv84hjnc4omg/gmNjo2cDW64QxA3RTJ5Sl///4xTBkg=="),
				}},
			DataStoreTypeID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStoreTypes/StorSimple8000Series"),
			ExtendedProperties: map[string]any{
				"extendedSaKey": nil,
				"resourceId":    "/subscriptions/c5fc377d-0085-41b9-86b7-cc96dc56d1e9/resourceGroups/ForDMS/providers/Microsoft.StorSimple/managers/BLR8600",
			},
			RepositoryID: to.Ptr("/subscriptions/c5fc377d-0085-41b9-86b7-cc96dc56d1e9/resourceGroups/ForDMS/providers/Microsoft.StorSimple/managers/BLR8600"),
			State:        to.Ptr(armhybriddatamanager.StateEnabled),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataStore = armhybriddatamanager.DataStore{
	// 	Name: to.Ptr("TestStorSimpleSource1"),
	// 	Type: to.Ptr("Microsoft.HybridData/dataManagers/dataStores"),
	// 	ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStores/TestStorSimpleSource1"),
	// 	Properties: &armhybriddatamanager.DataStoreProperties{
	// 		DataStoreTypeID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStoreTypes/StorSimple8000Series"),
	// 		ExtendedProperties: map[string]any{
	// 			"extendedSaKey": nil,
	// 			"resourceId": "/subscriptions/c5fc377d-0085-41b9-86b7-cc96dc56d1e9/resourceGroups/ForDMS/providers/Microsoft.StorSimple/managers/BLR8600",
	// 		},
	// 		RepositoryID: to.Ptr("/subscriptions/c5fc377d-0085-41b9-86b7-cc96dc56d1e9/resourceGroups/ForDMS/providers/Microsoft.StorSimple/managers/BLR8600"),
	// 		State: to.Ptr(armhybriddatamanager.StateEnabled),
	// 	},
	// }
}
