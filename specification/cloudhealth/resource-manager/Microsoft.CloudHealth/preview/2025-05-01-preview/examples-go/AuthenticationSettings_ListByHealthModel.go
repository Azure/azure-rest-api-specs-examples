package armcloudhealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cloudhealth/armcloudhealth"
)

// Generated from example definition: 2025-05-01-preview/AuthenticationSettings_ListByHealthModel.json
func ExampleAuthenticationSettingsClient_NewListByHealthModelPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcloudhealth.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAuthenticationSettingsClient().NewListByHealthModelPager("my-resource-group", "my-health-model", nil)
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
		// page = armcloudhealth.AuthenticationSettingsClientListByHealthModelResponse{
		// 	AuthenticationSettingListResult: armcloudhealth.AuthenticationSettingListResult{
		// 		Value: []*armcloudhealth.AuthenticationSetting{
		// 			{
		// 				Properties: &armcloudhealth.ManagedIdentityAuthenticationSettingProperties{
		// 					ManagedIdentityName: to.Ptr("SystemAssigned"),
		// 					ProvisioningState: to.Ptr(armcloudhealth.HealthModelProvisioningStateSucceeded),
		// 					DisplayName: to.Ptr("my-display-name"),
		// 					AuthenticationKind: to.Ptr(armcloudhealth.AuthenticationKindManagedIdentity),
		// 				},
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/my-resource-group/providers/Microsoft.CloudHealth/healthModels/my-health-model/authenticationSettings/my-name"),
		// 				Name: to.Ptr("my-name"),
		// 				Type: to.Ptr("Microsoft.CloudHealth/healthModels/authenticationSettings"),
		// 				SystemData: &armcloudhealth.SystemData{
		// 					CreatedBy: to.Ptr("my-user"),
		// 					CreatedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-18T14:04:09.327Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("my-user"),
		// 					LastModifiedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-18T14:04:09.328Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/ahgxpg"),
		// 	},
		// }
	}
}
