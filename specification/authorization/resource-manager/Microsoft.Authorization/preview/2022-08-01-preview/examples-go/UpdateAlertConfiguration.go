package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/preview/2022-08-01-preview/examples/UpdateAlertConfiguration.json
func ExampleAlertConfigurationsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAlertConfigurationsClient().Update(ctx, "subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f", "TooManyOwnersAssignedToResource", armauthorization.AlertConfiguration{
		Properties: &armauthorization.TooManyOwnersAssignedToResourceAlertConfigurationProperties{
			AlertConfigurationType:  to.Ptr("TooManyOwnersAssignedToResourceAlertConfiguration"),
			IsEnabled:               to.Ptr(true),
			ThresholdNumberOfOwners: to.Ptr[int32](2),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
