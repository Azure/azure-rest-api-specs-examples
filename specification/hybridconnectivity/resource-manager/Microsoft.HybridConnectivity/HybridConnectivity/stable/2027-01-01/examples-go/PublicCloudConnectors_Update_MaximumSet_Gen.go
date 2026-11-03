package armhybridconnectivity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity"
)

// Generated from example definition: 2027-01-01/PublicCloudConnectors_Update_MaximumSet_Gen.json
func ExamplePublicCloudConnectorsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPublicCloudConnectorsClient("5ACC4579-DB34-4C2F-8F8C-25061168F342").Update(ctx, "rghybridconnectivity", "abc", armhybridconnectivity.PublicCloudConnectorUpdate{
		Tags: map[string]*string{},
		Properties: &armhybridconnectivity.PublicCloudConnectorPropertiesUpdate{
			AwsCloudProfile: &armhybridconnectivity.AwsCloudProfileUpdate{
				ExcludedAccounts: []*string{
					to.Ptr("zrbtd"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhybridconnectivity.PublicCloudConnectorsClientUpdateResponse{
	// 	PublicCloudConnector: armhybridconnectivity.PublicCloudConnector{
	// 		Properties: &armhybridconnectivity.PublicCloudConnectorProperties{
	// 			AwsCloudProfile: &armhybridconnectivity.AwsCloudProfile{
	// 				AccountID: to.Ptr("snbnuxckevyqpm"),
	// 				ExcludedAccounts: []*string{
	// 					to.Ptr("zrbtd"),
	// 				},
	// 				IsOrganizationalAccount: to.Ptr(true),
	// 			},
	// 			HostType: to.Ptr(armhybridconnectivity.HostTypeAWS),
	// 			GCPCloudProfile: &armhybridconnectivity.GCPCloudProfile{
	// 				ProjectProperties: &armhybridconnectivity.GCPProjectProperties{
	// 					ProjectNumber: to.Ptr("mjubieitixhpm"),
	// 					ProjectID: to.Ptr("sample-project-id"),
	// 				},
	// 				OrganizationProperties: &armhybridconnectivity.GCPOrganizationProperties{
	// 					OrganizationID: to.Ptr("vqlzghfdinlamurmg"),
	// 					ManagementProjectNumber: to.Ptr("mjubieitixhpm"),
	// 					ManagementProjectID: to.Ptr("sample-management-project-id"),
	// 					ExcludedProjectNumbers: []*string{
	// 						to.Ptr("sepdnfxmhcrubtklwllxfbhju"),
	// 					},
	// 					ExcludedFolderIDs: []*string{
	// 						to.Ptr("xxl"),
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armhybridconnectivity.ResourceProvisioningStateSucceeded),
	// 			ConnectorPrimaryIdentifier: to.Ptr("20a4e2be-8158-4b9e-b512-7a1af6f827de"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("jpiglusfxynfcewcjwvvnn"),
	// 		ID: to.Ptr("/subscriptions/5ACC4579-DB34-4C2F-8F8C-25061168F342/providers/Microsoft.HybridConnectivity/PublicCloudConnectors/esixipkbydb"),
	// 		Name: to.Ptr("esixipkbydb"),
	// 		Type: to.Ptr("eelsjvqvkdxdncptsobrswhulnm"),
	// 		Kind: to.Ptr(armhybridconnectivity.HostType("sample-kind")),
	// 		SystemData: &armhybridconnectivity.SystemData{
	// 			CreatedBy: to.Ptr("rpxzkcrobprrdvuoqxz"),
	// 			CreatedByType: to.Ptr(armhybridconnectivity.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2024, time.January, 18, 22, 52, 7, 890000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("jidegyskxi"),
	// 			LastModifiedByType: to.Ptr(armhybridconnectivity.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2024, time.January, 18, 22, 52, 7, 890000000, time.UTC)),
	// 		},
	// 	},
	// }
}
