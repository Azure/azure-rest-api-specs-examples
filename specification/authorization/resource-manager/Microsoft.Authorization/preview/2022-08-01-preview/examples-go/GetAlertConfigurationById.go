package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/preview/2022-08-01-preview/examples/GetAlertConfigurationById.json
func ExampleAlertConfigurationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAlertConfigurationsClient().Get(ctx, "subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f", "TooManyOwnersAssignedToResource", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AlertConfiguration = armauthorization.AlertConfiguration{
	// 	Name: to.Ptr("TooManyOwnersAssignedToResource"),
	// 	Type: to.Ptr("Microsoft.Authorization/roleManagementAlertConfigurations"),
	// 	ID: to.Ptr("/subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleManagementAlertConfigurations/TooManyOwnersAssignedToResource"),
	// 	Properties: &armauthorization.TooManyOwnersAssignedToResourceAlertConfigurationProperties{
	// 		AlertConfigurationType: to.Ptr("TooManyOwnersAssignedToResourceAlertConfiguration"),
	// 		AlertDefinitionID: to.Ptr("TooManyOwnersAssignedToResource"),
	// 		IsEnabled: to.Ptr(true),
	// 		Scope: to.Ptr("subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f"),
	// 		ThresholdNumberOfOwners: to.Ptr[int32](3),
	// 		ThresholdPercentageOfOwnersOutOfAllRoleMembers: to.Ptr[int32](40),
	// 	},
	// }
}
