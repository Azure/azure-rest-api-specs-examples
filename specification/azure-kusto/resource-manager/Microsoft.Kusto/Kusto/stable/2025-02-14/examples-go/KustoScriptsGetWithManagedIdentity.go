package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: 2025-02-14/KustoScriptsGetWithManagedIdentity.json
func ExampleScriptsClient_Get_kustoScriptsGetWithManagedIdentity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("12345678-1234-1234-1234-123456789098", cred, nil)
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
	// res = armkusto.ScriptsClientGetResponse{
	// 	Script: armkusto.Script{
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/Kustodatabase8/Scripts/kustoScript"),
	// 		Name: to.Ptr("kustoCluster/Kustodatabase8/kustoScript"),
	// 		Type: to.Ptr("Microsoft.Kusto/Clusters/Databases/Scripts"),
	// 		Properties: &armkusto.ScriptProperties{
	// 			ScriptURL: to.Ptr("https://mysa.blob.core.windows.net/container/script.txt"),
	// 			ForceUpdateTag: to.Ptr("2bcf3c21-ffd1-4444-b9dd-e52e00ee53fe"),
	// 			ContinueOnErrors: to.Ptr(true),
	// 			ScriptLevel: to.Ptr(armkusto.ScriptLevelDatabase),
	// 			PrincipalPermissionsAction: to.Ptr(armkusto.PrincipalPermissionsActionRemovePermissionOnScriptCompletion),
	// 			ManagedIdentityResourceID: to.Ptr("/subscriptions/12345666-1234-1277-1234-123456789098/resourceGroups/kustomis/providers/Microsoft.ManagedIdentity/userAssignedIdentities/kustoMi1"),
	// 		},
	// 		SystemData: &armkusto.SystemData{
	// 			CreatedBy: to.Ptr("user@microsoft.com"),
	// 			CreatedByType: to.Ptr(armkusto.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-29T15:06:54.2757906Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armkusto.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-29T15:06:54.2757906Z"); return t}()),
	// 		},
	// 	},
	// }
}
