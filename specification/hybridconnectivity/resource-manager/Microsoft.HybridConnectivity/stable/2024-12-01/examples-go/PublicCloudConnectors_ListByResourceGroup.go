package armhybridconnectivity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity"
)

// Generated from example definition: 2024-12-01/PublicCloudConnectors_ListByResourceGroup.json
func ExamplePublicCloudConnectorsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPublicCloudConnectorsClient("5ACC4579-DB34-4C2F-8F8C-25061168F342").NewListByResourceGroupPager("rgpublicCloud", nil)
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
		// page = armhybridconnectivity.PublicCloudConnectorsClientListByResourceGroupResponse{
		// 	PublicCloudConnectorListResult: armhybridconnectivity.PublicCloudConnectorListResult{
		// 		Value: []*armhybridconnectivity.PublicCloudConnector{
		// 			{
		// 				Properties: &armhybridconnectivity.PublicCloudConnectorProperties{
		// 					AwsCloudProfile: &armhybridconnectivity.AwsCloudProfile{
		// 						ExcludedAccounts: []*string{
		// 							to.Ptr("rwgqpukglvbqmogqcliqolucp"),
		// 						},
		// 						AccountID: to.Ptr("troiiavknxmcpczvxwjhrdue"),
		// 						IsOrganizationalAccount: to.Ptr(true),
		// 					},
		// 					HostType: to.Ptr(armhybridconnectivity.HostTypeAWS),
		// 					ProvisioningState: to.Ptr(armhybridconnectivity.ResourceProvisioningStateSucceeded),
		// 					ConnectorPrimaryIdentifier: to.Ptr("20a4e2be-8158-4b9e-b512-7a1af6f827de"),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 				Location: to.Ptr("jpiglusfxynfcewcjwvvnn"),
		// 				ID: to.Ptr("/subscriptions/5ACC4579-DB34-4C2F-8F8C-25061168F342/providers/Microsoft.HybridConnectivity/PublicCloudConnectors/esixipkbydb"),
		// 				Name: to.Ptr("esixipkbydb"),
		// 				Type: to.Ptr("eelsjvqvkdxdncptsobrswhulnm"),
		// 				SystemData: &armhybridconnectivity.SystemData{
		// 					CreatedBy: to.Ptr("rpxzkcrobprrdvuoqxz"),
		// 					CreatedByType: to.Ptr(armhybridconnectivity.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-18T22:52:07.890Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("jidegyskxi"),
		// 					LastModifiedByType: to.Ptr(armhybridconnectivity.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-18T22:52:07.890Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
