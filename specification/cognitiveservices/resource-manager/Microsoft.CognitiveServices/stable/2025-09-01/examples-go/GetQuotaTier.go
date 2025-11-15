package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6bbb6daca7175b2cab69b20b2dd01094e3c6a534/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/GetQuotaTier.json
func ExampleQuotaTiersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewQuotaTiersClient().Get(ctx, "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.QuotaTier = armcognitiveservices.QuotaTier{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/quotaTiers"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.CognitiveServices/quotaTiers/default"),
	// 	Properties: &armcognitiveservices.QuotaTierProperties{
	// 		AssignmentDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-09T18:13:29.389Z"); return t}()),
	// 		CurrentTierName: to.Ptr("Free-Tier"),
	// 		TierUpgradeEligibilityInfo: &armcognitiveservices.QuotaTierUpgradeEligibilityInfo{
	// 			NextTierName: to.Ptr("Tier-1"),
	// 			UpgradeApplicableDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-15T18:13:29.389Z"); return t}()),
	// 			UpgradeAvailabilityStatus: to.Ptr(armcognitiveservices.UpgradeAvailabilityStatusAvailable),
	// 		},
	// 		TierUpgradePolicy: to.Ptr(armcognitiveservices.TierUpgradePolicyOnceUpgradeIsAvailable),
	// 	},
	// }
}
