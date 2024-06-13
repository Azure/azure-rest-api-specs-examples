package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/settings/UpdateEyesOnSetting.json
func ExampleProductSettingsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProductSettingsClient().Update(ctx, "myRg", "myWorkspace", "EyesOn", &armsecurityinsights.EyesOn{
		Etag:       to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		Kind:       to.Ptr(armsecurityinsights.SettingKindEyesOn),
		Properties: &armsecurityinsights.EyesOnSettingsProperties{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.ProductSettingsClientUpdateResponse{
	// 	                            SettingsClassification: &armsecurityinsights.EyesOn{
	// 		Name: to.Ptr("EyesOn"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/settings"),
	// 		ID: to.Ptr("/subscriptions/bd794837-4d29-4647-9105-6339bfdb4e6a/resourceGroups/mms-eus/providers/Microsoft.OperationalInsights/workspaces/avdvirInt/providers/Microsoft.SecurityInsights/settings/EyesOn"),
	// 		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 		Kind: to.Ptr(armsecurityinsights.SettingKindEyesOn),
	// 		Properties: &armsecurityinsights.EyesOnSettingsProperties{
	// 			IsEnabled: to.Ptr(true),
	// 		},
	// 	},
	// 	                        }
}
