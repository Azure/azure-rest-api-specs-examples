package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/preview/2022-08-01-preview/examples/GetAlertById.json
func ExampleAlertsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAlertsClient().Get(ctx, "subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f", "TooManyOwnersAssignedToResource", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Alert = armauthorization.Alert{
	// 	Name: to.Ptr("TooManyOwnersAssignedToResource"),
	// 	Type: to.Ptr("Microsoft.Authorization/roleManagementAlerts"),
	// 	ID: to.Ptr("/subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleManagementAlerts/TooManyOwnersAssignedToResource"),
	// 	Properties: &armauthorization.AlertProperties{
	// 		AlertConfiguration: &armauthorization.AlertConfiguration{
	// 			Name: to.Ptr("TooManyOwnersAssignedToResource"),
	// 			Type: to.Ptr("Microsoft.Authorization/roleManagementAlertConfigurations"),
	// 			ID: to.Ptr("/subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleManagementAlertConfigurations/TooManyOwnersAssignedToResource"),
	// 			Properties: &armauthorization.TooManyOwnersAssignedToResourceAlertConfigurationProperties{
	// 				AlertConfigurationType: to.Ptr("TooManyOwnersAssignedToResourceAlertConfiguration"),
	// 				AlertDefinitionID: to.Ptr("TooManyOwnersAssignedToResource"),
	// 				IsEnabled: to.Ptr(true),
	// 				Scope: to.Ptr("subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f"),
	// 				ThresholdNumberOfOwners: to.Ptr[int32](2),
	// 				ThresholdPercentageOfOwnersOutOfAllRoleMembers: to.Ptr[int32](3),
	// 			},
	// 		},
	// 		AlertDefinition: &armauthorization.AlertDefinition{
	// 			Name: to.Ptr("TooManyOwnersAssignedToResource"),
	// 			Type: to.Ptr("Microsoft.Authorization/roleManagementAlertDefinitions"),
	// 			ID: to.Ptr("/subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleManagementAlertDefinitions/TooManyOwnersAssignedToResource"),
	// 			Properties: &armauthorization.AlertDefinitionProperties{
	// 				Description: to.Ptr("The number of users with the Owner role is too high. We recommend assigning these individuals to less privileged roles or roles more suitable to their daily needs. Take a moment to review the current assignments, and suggested changes here."),
	// 				DisplayName: to.Ptr("Too many owners assigned to a resource"),
	// 				HowToPrevent: to.Ptr("Choose a role that provides the fewest privileges necessary for a user or group to complete their tasks."),
	// 				IsConfigurable: to.Ptr(true),
	// 				IsRemediatable: to.Ptr(true),
	// 				MitigationSteps: to.Ptr("To mitigate this issue, reduce the number of users in the Owner role. Review the list of users in the list, and reassign them to a less privileged role such as Contributor."),
	// 				Scope: to.Ptr("subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f"),
	// 				SecurityImpact: to.Ptr("As the number of users with the owner role increases, so does the potential for malicious or mistaken actions affecting your resource."),
	// 				SeverityLevel: to.Ptr(armauthorization.SeverityLevelMedium),
	// 			},
	// 		},
	// 		AlertIncidents: []*armauthorization.AlertIncident{
	// 			{
	// 				Name: to.Ptr("a9f38501-74ec-43ea-8663-6c538602150d"),
	// 				Type: to.Ptr("Microsoft.Authorization/roleManagementAlerts/alertIncidents"),
	// 				ID: to.Ptr("/subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleManagementAlerts/TooManyOwnersAssignedToResource/alertIncidents/a9f38501-74ec-43ea-8663-6c538602150d"),
	// 				Properties: &armauthorization.TooManyOwnersAssignedToResourceAlertIncidentProperties{
	// 					AlertIncidentType: to.Ptr("TooManyOwnersAssignedToResourceAlertIncident"),
	// 					AssigneeName: to.Ptr("testUser"),
	// 					AssigneeType: to.Ptr("User"),
	// 				},
	// 		}},
	// 		IncidentCount: to.Ptr[int32](1),
	// 		IsActive: to.Ptr(true),
	// 		LastModifiedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-05T03:04:06.467Z"); return t}()),
	// 		LastScannedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-06T18:25:00.380Z"); return t}()),
	// 		Scope: to.Ptr("subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f"),
	// 	},
	// }
}
