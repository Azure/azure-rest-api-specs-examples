package armiotsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotsecurity/armiotsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/iotsecurity/resource-manager/Microsoft.IoTSecurity/preview/2021-02-01-preview/examples/DefenderSettings/Get.json
func ExampleDefenderSettingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDefenderSettingsClient().Get(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DefenderSettingsModel = armiotsecurity.DefenderSettingsModel{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.IoTSecurity/defenderSettings"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/providers/Microsoft.IoTSecurity/defenderSettings/default"),
	// 	Properties: &armiotsecurity.DefenderSettingsProperties{
	// 		DeviceQuota: to.Ptr[int32](2000),
	// 		EvaluationEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-29T00:00:00.000Z"); return t}()),
	// 		MdeIntegration: &armiotsecurity.DefenderSettingsPropertiesMdeIntegration{
	// 			Status: to.Ptr(armiotsecurity.MdeIntegrationEnabled),
	// 		},
	// 		OnboardingKind: to.Ptr(armiotsecurity.OnboardingKindEvaluation),
	// 		SentinelWorkspaceResourceIDs: []*string{
	// 			to.Ptr("/subscriptions/c4930e90-cd72-4aa5-93e9-2d081d129569/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace1")},
	// 		},
	// 	}
}
