package armprivatetrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatetrafficmanager/armprivatetrafficmanager"
)

// Generated from example definition: 2026-10-01-preview/HealthPolicies_ListByParent_MaximumSet_Gen.json
func ExampleHealthPoliciesClient_NewListByParentPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprivatetrafficmanager.NewClientFactory("10B6D88D-ADF4-4281-B3D0-B5A6702DEEDA", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewHealthPoliciesClient().NewListByParentPager("rgprivateTrafficManager", "myProfile", nil)
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
		// page = armprivatetrafficmanager.HealthPoliciesClientListByParentResponse{
		// 	HealthPolicyListResult: armprivatetrafficmanager.HealthPolicyListResult{
		// 		Value: []armprivatetrafficmanager.HealthPolicyClassification{
		// 			&armprivatetrafficmanager.ProbeHealthPolicy{
		// 				Properties: &armprivatetrafficmanager.HealthPolicyProperties{
		// 					ProbeConfig: &armprivatetrafficmanager.ProbeConfig{
		// 						Protocol: to.Ptr(armprivatetrafficmanager.ProtocolHTTPS),
		// 						Port: to.Ptr[int64](443),
		// 						Path: to.Ptr("/health"),
		// 						IntervalInSeconds: to.Ptr[int64](30),
		// 						TimeoutInSeconds: to.Ptr[int64](10),
		// 						ToleratedNumberOfFailures: to.Ptr[int64](3),
		// 					},
		// 					ProvisioningState: to.Ptr(armprivatetrafficmanager.ProvisioningStateSucceeded),
		// 				},
		// 				Kind: to.Ptr(armprivatetrafficmanager.HealthPolicyKindProbe),
		// 				ID: to.Ptr("/subscriptions/10B6D88D-ADF4-4281-B3D0-B5A6702DEEDA/resourceGroups/rgprivateTrafficManager/providers/Microsoft.Network/privateTrafficManagerProfiles/myProfile/healthPolicies/myHealthPolicy"),
		// 				Name: to.Ptr("myHealthPolicy"),
		// 				Type: to.Ptr("Microsoft.Network/privateTrafficManagerProfiles/healthPolicies"),
		// 				SystemData: &armprivatetrafficmanager.SystemData{
		// 					CreatedBy: to.Ptr("user@contoso.com"),
		// 					CreatedByType: to.Ptr(armprivatetrafficmanager.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2026, time.February, 18, 10, 5, 58, 138000000, time.UTC)),
		// 					LastModifiedBy: to.Ptr("user@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armprivatetrafficmanager.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2026, time.February, 18, 10, 5, 58, 139000000, time.UTC)),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/10B6D88D-ADF4-4281-B3D0-B5A6702DEEDA/resourceGroups/rgprivateTrafficManager/providers/Microsoft.Network/privateTrafficManagerProfiles/myProfile/healthPolicies?api-version=2026-10-01-preview&$skipToken=eyJuZXh0UGFnZSI6IjIifQ"),
		// 	},
		// }
	}
}
