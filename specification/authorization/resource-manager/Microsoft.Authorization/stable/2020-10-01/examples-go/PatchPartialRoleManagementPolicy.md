Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fauthorization%2Farmauthorization%2Fv1.0.0/sdk/resourcemanager/authorization/armauthorization/README.md) on how to add the SDK to your project and authenticate.

```go
package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/PatchPartialRoleManagementPolicy.json
func ExampleRoleManagementPoliciesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armauthorization.NewRoleManagementPoliciesClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"providers/Microsoft.Subscription/subscriptions/129ff972-28f8-46b8-a726-e497be039368",
		"570c3619-7688-4b34-b290-2b8bb3ccab2a",
		armauthorization.RoleManagementPolicy{
			Properties: &armauthorization.RoleManagementPolicyProperties{
				Rules: []armauthorization.RoleManagementPolicyRuleClassification{
					&armauthorization.RoleManagementPolicyExpirationRule{
						ID:       to.Ptr("Expiration_Admin_Eligibility"),
						RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyExpirationRule),
						Target: &armauthorization.RoleManagementPolicyRuleTarget{
							Caller: to.Ptr("Admin"),
							Level:  to.Ptr("Eligibility"),
							Operations: []*string{
								to.Ptr("All")},
						},
						IsExpirationRequired: to.Ptr(false),
						MaximumDuration:      to.Ptr("P180D"),
					},
					&armauthorization.RoleManagementPolicyNotificationRule{
						ID:       to.Ptr("Notification_Admin_Admin_Eligibility"),
						RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
						Target: &armauthorization.RoleManagementPolicyRuleTarget{
							Caller: to.Ptr("Admin"),
							Level:  to.Ptr("Eligibility"),
							Operations: []*string{
								to.Ptr("All")},
						},
						IsDefaultRecipientsEnabled: to.Ptr(false),
						NotificationLevel:          to.Ptr(armauthorization.NotificationLevelCritical),
						NotificationRecipients: []*string{
							to.Ptr("admin_admin_eligible@test.com")},
						NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
						RecipientType:    to.Ptr(armauthorization.RecipientTypeAdmin),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
