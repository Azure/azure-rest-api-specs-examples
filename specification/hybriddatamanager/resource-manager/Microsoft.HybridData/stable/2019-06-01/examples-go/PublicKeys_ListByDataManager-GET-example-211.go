package armhybriddatamanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/PublicKeys_ListByDataManager-GET-example-211.json
func ExamplePublicKeysClient_NewListByDataManagerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybriddatamanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPublicKeysClient().NewListByDataManagerPager("ResourceGroupForSDKTest", "TestAzureSDKOperations", nil)
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
		// page.PublicKeyList = armhybriddatamanager.PublicKeyList{
		// 	Value: []*armhybriddatamanager.PublicKey{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers/publicKeys"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/publicKeys/default"),
		// 			Properties: &armhybriddatamanager.PublicKeyProperties{
		// 				DataServiceLevel1Key: &armhybriddatamanager.Key{
		// 					EncryptionChunkSizeInBytes: to.Ptr[int32](245),
		// 					KeyExponent: to.Ptr("AQAB"),
		// 					KeyModulus: to.Ptr("mjDxqxGmawhC15TUM/oKCe2rRg6nM+IEqujgn17vc1litm3TPmv7rtDr4Y/L/t+tYCug7aXxJtwGA9ETOUF9iUoGPE3zBKMGPJhO5nRF3IW27OzYUNTdHgUjlV0ba5QlZQ/f5ideEboJvdlw05ofPVQKZ9Hh95/9sOFYuNBKP0LPwKz1VrrhvM7tVgdIhZdekuIOt4S+7WjRV5J+XT0jlhwUBEIxx8knRPagmxygSZM3h/FbX+mEbduIwVy+y1HwtfwVq3PyR9YIDjVDuc3+6VNZd69TEIHqRQlbwb2jkitgEHx/Vs32KtDyfRZgkhA6ZGZRlnEiX3R0YRzjCt5xCw=="),
		// 				},
		// 				DataServiceLevel2Key: &armhybriddatamanager.Key{
		// 					EncryptionChunkSizeInBytes: to.Ptr[int32](245),
		// 					KeyExponent: to.Ptr("AQAB"),
		// 					KeyModulus: to.Ptr("rtCsQNdCaDwLHIvgkhkKldvUNjili+rsj8CVaKGTwAyQGvAIwKOe3zfwW3TFaeycTFAQ1payBjY2tW9uWSWDTJRpHZVYTTX/1mjnXHTqcZYsgSkblt0PhLWbbYATGMmyBie0XM3Xfy1ilwAMYNHu+YaW56NyFpepyNcheZbmkD4/Vveh+5JStwObqEp1vsagraQ/IqUDCAETRxFc3iIWJZnqW2yfIWZshky20fkmyBnRrpe5fexpj6Xz4VHT76Lj+7bTEbsFSq7fNUjRcCIf2gat8bBN4HD2w//GZVCKarQG6G0kilA0bDZHFAzVeTs2+UYB+1GA+r+Uy3SOEIdLPw=="),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
