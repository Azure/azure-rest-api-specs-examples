package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4883fa5dbf6f2c9093fac8ce334547e9dfac68fa/specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/NetworkWatcherFlowLogCreate.json
func ExampleFlowLogsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFlowLogsClient().BeginCreateOrUpdate(ctx, "rg1", "nw1", "fl", armnetwork.FlowLog{
		Location: to.Ptr("centraluseuap"),
		Identity: &armnetwork.ManagedServiceIdentity{
			Type: to.Ptr(armnetwork.ResourceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armnetwork.Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties{
				"/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {},
			},
		},
		Properties: &armnetwork.FlowLogPropertiesFormat{
			Format: &armnetwork.FlowLogFormatParameters{
				Type:    to.Ptr(armnetwork.FlowLogFormatTypeJSON),
				Version: to.Ptr[int32](1),
			},
			Enabled:                  to.Ptr(true),
			EnabledFilteringCriteria: to.Ptr("srcIP=158.255.7.8 || dstPort=56891"),
			StorageID:                to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/nwtest1mgvbfmqsigdxe"),
			TargetResourceID:         to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/desmondcentral-nsg"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FlowLog = armnetwork.FlowLog{
	// 	Name: to.Ptr("Microsoft.Networkdesmond-rgdesmondcentral-nsg"),
	// 	Type: to.Ptr("Microsoft.Network/networkWatchers/FlowLogs"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkWatchers/nw/FlowLogs/fl"),
	// 	Location: to.Ptr("centraluseuap"),
	// 	Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 	Identity: &armnetwork.ManagedServiceIdentity{
	// 		Type: to.Ptr(armnetwork.ResourceIdentityTypeUserAssigned),
	// 		UserAssignedIdentities: map[string]*armnetwork.Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties{
	// 			"/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": &armnetwork.Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties{
	// 				ClientID: to.Ptr("c16d15e1-f60a-40e4-8a05-df3d3f655c14"),
	// 				PrincipalID: to.Ptr("e3858881-e40c-43bd-9cde-88da39c05023"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armnetwork.FlowLogPropertiesFormat{
	// 		Format: &armnetwork.FlowLogFormatParameters{
	// 			Type: to.Ptr(armnetwork.FlowLogFormatTypeJSON),
	// 			Version: to.Ptr[int32](1),
	// 		},
	// 		Enabled: to.Ptr(true),
	// 		EnabledFilteringCriteria: to.Ptr("srcIP=158.255.7.8 || dstPort=56891"),
	// 		FlowAnalyticsConfiguration: &armnetwork.TrafficAnalyticsProperties{
	// 		},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		RetentionPolicy: &armnetwork.RetentionPolicyParameters{
	// 			Days: to.Ptr[int32](0),
	// 			Enabled: to.Ptr(false),
	// 		},
	// 		StorageID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/nwtest1mgvbfmqsigdxe"),
	// 		TargetResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		TargetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/desmondcentral-nsg"),
	// 	},
	// }
}
