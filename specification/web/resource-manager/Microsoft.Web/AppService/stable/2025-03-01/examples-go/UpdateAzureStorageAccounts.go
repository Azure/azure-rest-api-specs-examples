package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/UpdateAzureStorageAccounts.json
func ExampleWebAppsClient_UpdateAzureStorageAccounts() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWebAppsClient().UpdateAzureStorageAccounts(ctx, "testrg123", "sitef6141", armappservice.AzureStoragePropertyDictionaryResource{
		Properties: map[string]*armappservice.AzureStorageInfoValue{
			"account1": {
				Type:        to.Ptr(armappservice.AzureStorageTypeAzureFiles),
				AccessKey:   to.Ptr("26515^%@#*"),
				AccountName: to.Ptr("testsa"),
				MountPath:   to.Ptr("/mounts/a/files"),
				ShareName:   to.Ptr("web"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AzureStoragePropertyDictionaryResource = armappservice.AzureStoragePropertyDictionaryResource{
	// 	Name: to.Ptr("web"),
	// 	Type: to.Ptr("Microsoft.Web/sites/config"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/sites/sitef6141/config/web"),
	// 	Kind: to.Ptr("app"),
	// 	Properties: map[string]*armappservice.AzureStorageInfoValue{
	// 		"account1": &armappservice.AzureStorageInfoValue{
	// 			Type: to.Ptr(armappservice.AzureStorageTypeAzureFiles),
	// 			AccountName: to.Ptr("testsa"),
	// 			MountPath: to.Ptr("/mounts/a/files"),
	// 			ShareName: to.Ptr("web"),
	// 			State: to.Ptr(armappservice.AzureStorageStateOk),
	// 		},
	// 	},
	// }
}
