package armautomanage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automanage/armautomanage"
)

// x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/preview/2020-06-30-preview/examples/createOrUpdateConfigurationProfilePreference.json
func ExampleConfigurationProfilePreferencesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomanage.NewConfigurationProfilePreferencesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<configuration-profile-preference-name>",
		"<resource-group-name>",
		armautomanage.ConfigurationProfilePreference{
			TrackedResource: armautomanage.TrackedResource{
				Location: to.StringPtr("<location>"),
				Tags: map[string]*string{
					"Organization": to.StringPtr("Administration"),
				},
			},
			Properties: &armautomanage.ConfigurationProfilePreferenceProperties{
				AntiMalware: &armautomanage.ConfigurationProfilePreferenceAntiMalware{
					EnableRealTimeProtection: armautomanage.EnableRealTimeProtectionTrue.ToPtr(),
				},
				VMBackup: &armautomanage.ConfigurationProfilePreferenceVMBackup{
					TimeZone: to.StringPtr("<time-zone>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ConfigurationProfilePreference.ID: %s\n", *res.ID)
}
