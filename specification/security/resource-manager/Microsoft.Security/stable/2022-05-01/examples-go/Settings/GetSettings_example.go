package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/stable/2022-05-01/examples/Settings/GetSettings_example.json
func ExampleSettingsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSettingsClient().NewListPager(nil)
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
		// page.SettingsList = armsecurity.SettingsList{
		// 	Value: []armsecurity.SettingClassification{
		// 		&armsecurity.DataExportSettings{
		// 			Name: to.Ptr("WDATP"),
		// 			Type: to.Ptr("Microsoft.Security/settings"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/settings/WDATP"),
		// 			Kind: to.Ptr(armsecurity.SettingKindDataExportSettings),
		// 			Properties: &armsecurity.DataExportSettingProperties{
		// 				Enabled: to.Ptr(false),
		// 			},
		// 		},
		// 		&armsecurity.DataExportSettings{
		// 			Name: to.Ptr("WDATP_EXCLUDE_LINUX_PUBLIC_PREVIEW"),
		// 			Type: to.Ptr("Microsoft.Security/settings"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/settings/WDATP_EXCLUDE_LINUX_PUBLIC_PREVIEW"),
		// 			Kind: to.Ptr(armsecurity.SettingKindDataExportSettings),
		// 			Properties: &armsecurity.DataExportSettingProperties{
		// 				Enabled: to.Ptr(false),
		// 			},
		// 		},
		// 		&armsecurity.DataExportSettings{
		// 			Name: to.Ptr("WDATP_UNIFIED_SOLUTION"),
		// 			Type: to.Ptr("Microsoft.Security/settings"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/settings/WDATP_UNIFIED_SOLUTION"),
		// 			Kind: to.Ptr(armsecurity.SettingKindDataExportSettings),
		// 			Properties: &armsecurity.DataExportSettingProperties{
		// 				Enabled: to.Ptr(false),
		// 			},
		// 		},
		// 		&armsecurity.AlertSyncSettings{
		// 			Name: to.Ptr("Sentinel"),
		// 			Type: to.Ptr("Microsoft.Security/settings"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/settings/Sentinel"),
		// 			Kind: to.Ptr(armsecurity.SettingKindAlertSyncSettings),
		// 			Properties: &armsecurity.AlertSyncSettingProperties{
		// 				Enabled: to.Ptr(false),
		// 			},
		// 	}},
		// }
	}
}
