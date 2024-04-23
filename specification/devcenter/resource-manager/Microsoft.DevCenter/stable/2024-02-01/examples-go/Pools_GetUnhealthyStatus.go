package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/Pools_GetUnhealthyStatus.json
func ExamplePoolsClient_Get_poolsGetUnhealthyStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPoolsClient().Get(ctx, "rg1", "DevProject", "DevPool", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Pool = armdevcenter.Pool{
	// 	Name: to.Ptr("DevPool"),
	// 	Type: to.Ptr("Microsoft.DevCenter/pools"),
	// 	ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/projects/DevProject/pools/DevPool"),
	// 	SystemData: &armdevcenter.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:00:36.993Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:30:36.993Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user1"),
	// 		LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("centralus"),
	// 	Properties: &armdevcenter.PoolProperties{
	// 		DevBoxDefinitionName: to.Ptr("WebDevBox"),
	// 		DisplayName: to.Ptr("Developer Pool"),
	// 		LicenseType: to.Ptr(armdevcenter.LicenseTypeWindowsClient),
	// 		LocalAdministrator: to.Ptr(armdevcenter.LocalAdminStatusEnabled),
	// 		ManagedVirtualNetworkRegions: []*string{
	// 			to.Ptr("centralus")},
	// 			NetworkConnectionName: to.Ptr("Network1-westus2"),
	// 			SingleSignOnStatus: to.Ptr(armdevcenter.SingleSignOnStatusDisabled),
	// 			StopOnDisconnect: &armdevcenter.StopOnDisconnectConfiguration{
	// 				GracePeriodMinutes: to.Ptr[int32](60),
	// 				Status: to.Ptr(armdevcenter.StopOnDisconnectEnableStatusEnabled),
	// 			},
	// 			VirtualNetworkType: to.Ptr(armdevcenter.VirtualNetworkTypeManaged),
	// 			DevBoxCount: to.Ptr[int32](1),
	// 			HealthStatus: to.Ptr(armdevcenter.HealthStatusUnhealthy),
	// 			HealthStatusDetails: []*armdevcenter.HealthStatusDetail{
	// 				{
	// 					Code: to.Ptr("NetworkConnectionUnhealthy"),
	// 					Message: to.Ptr("The Pool's Network Connection is in an unhealthy state. Check the Network Connection's health status for more details."),
	// 				},
	// 				{
	// 					Code: to.Ptr("ImageValidationFailed"),
	// 					Message: to.Ptr("Image validation has failed. Check the Dev Box Definition's image validation status for more details."),
	// 				},
	// 				{
	// 					Code: to.Ptr("IntuneValidationFailed"),
	// 					Message: to.Ptr("Intune license validation has failed. The tenant does not have a valid Intune license."),
	// 			}},
	// 			ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateSucceeded),
	// 		},
	// 	}
}
