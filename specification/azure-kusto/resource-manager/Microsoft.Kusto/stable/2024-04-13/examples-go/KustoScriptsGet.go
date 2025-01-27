package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8adce17dc500f338f86f18af30aac61dcb71c5f/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoScriptsGet.json
func ExampleScriptsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScriptsClient().Get(ctx, "kustorptest", "kustoCluster", "Kustodatabase8", "kustoScript", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Script = armkusto.Script{
	// 	Name: to.Ptr("kustoCluster/Kustodatabase8/kustoScript"),
	// 	Type: to.Ptr("Microsoft.Kusto/Clusters/Databases/Scripts"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/Kustodatabase8/Scripts/kustoScript"),
	// 	Properties: &armkusto.ScriptProperties{
	// 		ContinueOnErrors: to.Ptr(true),
	// 		ForceUpdateTag: to.Ptr("2bcf3c21-ffd1-4444-b9dd-e52e00ee53fe"),
	// 		PrincipalPermissionsAction: to.Ptr(armkusto.PrincipalPermissionsActionRemovePermissionOnScriptCompletion),
	// 		ScriptLevel: to.Ptr(armkusto.ScriptLevelDatabase),
	// 		ScriptURL: to.Ptr("https://mysa.blob.core.windows.net/container/script.txt"),
	// 	},
	// 	SystemData: &armkusto.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-29T15:06:54.275Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@microsoft.com"),
	// 		CreatedByType: to.Ptr(armkusto.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-29T15:06:54.275Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@microsoft.com"),
	// 		LastModifiedByType: to.Ptr(armkusto.CreatedByTypeUser),
	// 	},
	// }
}
