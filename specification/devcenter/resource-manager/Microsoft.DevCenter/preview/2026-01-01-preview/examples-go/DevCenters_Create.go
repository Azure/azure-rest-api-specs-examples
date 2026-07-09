package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v3"
)

// Generated from example definition: 2026-01-01-preview/DevCenters_Create.json
func ExampleDevCentersClient_BeginCreateOrUpdate_devCentersCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("0ac520ee-14c0-480f-b6c9-0a90c58fffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDevCentersClient().BeginCreateOrUpdate(ctx, "rg1", "Contoso", armdevcenter.DevCenter{
		Location: to.Ptr("centralus"),
		Properties: &armdevcenter.Properties{
			DevBoxProvisioningSettings: &armdevcenter.DevBoxProvisioningSettings{
				InstallAzureMonitorAgentEnableStatus: to.Ptr(armdevcenter.InstallAzureMonitorAgentEnableStatusEnabled),
			},
			DisplayName: to.Ptr("ContosoDevCenter"),
			ProjectCatalogSettings: &armdevcenter.ProjectCatalogSettingsInfo{
				CatalogItemSyncEnableStatus: to.Ptr(armdevcenter.CatalogItemSyncEnableStatusEnabled),
			},
		},
		Tags: map[string]*string{
			"CostCode": to.Ptr("12345"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdevcenter.DevCentersClientCreateOrUpdateResponse{
	// 	DevCenter: armdevcenter.DevCenter{
	// 		Name: to.Ptr("Contoso"),
	// 		Type: to.Ptr("Microsoft.DevCenter/devcenters"),
	// 		ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/Contoso"),
	// 		Location: to.Ptr("centralus"),
	// 		Properties: &armdevcenter.Properties{
	// 			DevBoxProvisioningSettings: &armdevcenter.DevBoxProvisioningSettings{
	// 				InstallAzureMonitorAgentEnableStatus: to.Ptr(armdevcenter.InstallAzureMonitorAgentEnableStatusEnabled),
	// 			},
	// 			DevCenterURI: to.Ptr("https://4c7c8922-78e9-4928-aa6f-75ba59355371-contoso.centralus.devcenter.azure.com"),
	// 			DisplayName: to.Ptr("ContosoDevCenter"),
	// 			NetworkSettings: &armdevcenter.NetworkSettings{
	// 				MicrosoftHostedNetworkEnableStatus: to.Ptr(armdevcenter.MicrosoftHostedNetworkEnableStatusEnabled),
	// 			},
	// 			ProjectCatalogSettings: &armdevcenter.ProjectCatalogSettingsInfo{
	// 				CatalogItemSyncEnableStatus: to.Ptr(armdevcenter.CatalogItemSyncEnableStatusEnabled),
	// 			},
	// 			ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateSucceeded),
	// 		},
	// 		SystemData: &armdevcenter.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-11T22:00:08.896Z"); return t}()),
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-11T22:00:10.896Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user2"),
	// 			LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 			"CostCode": to.Ptr("12345"),
	// 			"hidden-title": to.Ptr("ContosoDevCenter"),
	// 		},
	// 	},
	// }
}
