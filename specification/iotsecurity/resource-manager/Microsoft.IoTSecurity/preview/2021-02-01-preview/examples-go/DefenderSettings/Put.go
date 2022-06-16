package armiotsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotsecurity/armiotsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iotsecurity/resource-manager/Microsoft.IoTSecurity/preview/2021-02-01-preview/examples/DefenderSettings/Put.json
func ExampleDefenderSettingsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armiotsecurity.NewDefenderSettingsClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		armiotsecurity.DefenderSettingsModel{
			Properties: &armiotsecurity.DefenderSettingsProperties{
				DeviceQuota: to.Ptr[int32](2000),
				MdeIntegration: &armiotsecurity.DefenderSettingsPropertiesMdeIntegration{
					Status: to.Ptr(armiotsecurity.MdeIntegrationEnabled),
				},
				OnboardingKind: to.Ptr(armiotsecurity.OnboardingKindDefault),
				SentinelWorkspaceResourceIDs: []*string{
					to.Ptr("/subscriptions/c4930e90-cd72-4aa5-93e9-2d081d129569/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace1")},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
