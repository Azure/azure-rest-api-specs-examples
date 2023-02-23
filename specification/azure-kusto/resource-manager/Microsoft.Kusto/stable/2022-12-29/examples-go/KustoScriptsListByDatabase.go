package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-12-29/examples/KustoScriptsListByDatabase.json
func ExampleScriptsClient_NewListByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armkusto.NewScriptsClient("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByDatabasePager("kustorptest", "kustoCluster", "Kustodatabase8", nil)
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
		// page.ScriptListResult = armkusto.ScriptListResult{
		// 	Value: []*armkusto.Script{
		// 		{
		// 			Name: to.Ptr("kustoCluster/Kustodatabase8/kustoScript1"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/Databases/Scripts"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/Kustodatabase8/Scripts/kustoScript1"),
		// 			Properties: &armkusto.ScriptProperties{
		// 				ContinueOnErrors: to.Ptr(true),
		// 				ForceUpdateTag: to.Ptr("2bcf3c21-ffd1-4444-b9dd-e52e00ee53fe"),
		// 				ScriptURL: to.Ptr("https://mysa.blob.core.windows.net/container/script.txt"),
		// 			},
		// 			SystemData: &armkusto.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-29T15:06:54.2757906Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@microsoft.com"),
		// 				CreatedByType: to.Ptr(armkusto.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-29T15:06:54.2757906Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@microsoft.com"),
		// 				LastModifiedByType: to.Ptr(armkusto.CreatedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("kustoCluster/Kustodatabase8/kustoScript2"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/Databases/Scripts"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/Kustodatabase8/Scripts/kustoScript2"),
		// 			Properties: &armkusto.ScriptProperties{
		// 				ContinueOnErrors: to.Ptr(true),
		// 				ForceUpdateTag: to.Ptr("2bcf3c21-ffd1-4444-b9dd-e52e00ee53fe"),
		// 				ScriptURL: to.Ptr("https://mysa.blob.core.windows.net/container/script2.txt"),
		// 			},
		// 			SystemData: &armkusto.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-29T15:06:54.2757906Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@microsoft.com"),
		// 				CreatedByType: to.Ptr(armkusto.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-30T16:07:55.2831451Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2@microsoft.com"),
		// 				LastModifiedByType: to.Ptr(armkusto.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
