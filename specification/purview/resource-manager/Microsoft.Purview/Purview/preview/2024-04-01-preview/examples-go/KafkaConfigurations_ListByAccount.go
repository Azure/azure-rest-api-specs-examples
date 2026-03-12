package armpurview_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purview/armpurview/v2"
)

// Generated from example definition: 2024-04-01-preview/KafkaConfigurations_ListByAccount.json
func ExampleKafkaConfigurationsClient_NewListByAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurview.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewKafkaConfigurationsClient().NewListByAccountPager("rgpurview", "account1", &armpurview.KafkaConfigurationsClientListByAccountOptions{
		SkipToken: to.Ptr("token")})
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
		// page = armpurview.KafkaConfigurationsClientListByAccountResponse{
		// 	KafkaConfigurationList: armpurview.KafkaConfigurationList{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rgpurview/providers/Microsoft.Purview/accounts/account1/kafkaConfigurations?api-version=2024-04-01-preview&$skipToken=token2"),
		// 		Value: []*armpurview.KafkaConfiguration{
		// 			{
		// 				Name: to.Ptr("kafkaconfigName"),
		// 				Type: to.Ptr("Microsoft.Purview/accounts/kafkaconfiguration"),
		// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/SampleResourceGroup/providers/Microsoft.Purview/accounts/account1/kafkaConfigurations/default"),
		// 				Properties: &armpurview.KafkaConfigurationProperties{
		// 					ConsumerGroup: to.Ptr("consumerGroup"),
		// 					Credentials: &armpurview.Credentials{
		// 						Type: to.Ptr(armpurview.KafkaConfigurationIdentityTypeUserAssigned),
		// 						IdentityID: to.Ptr("/subscriptions/47e8596d-ee73-4eb2-b6b4-cc13c2b87ssd/resourceGroups/testRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testId"),
		// 					},
		// 					EventHubPartitionID: to.Ptr("partitionId"),
		// 					EventHubResourceID: to.Ptr("/subscriptions/225be6fe-ec1c-4d51-a368-f69348d2e6c5/resourceGroups/testRG/providers/Microsoft.EventHub/namespaces/eventHubNameSpaceName"),
		// 					EventHubType: to.Ptr(armpurview.EventHubTypeNotification),
		// 					EventStreamingState: to.Ptr(armpurview.EventStreamingStateDisabled),
		// 					EventStreamingType: to.Ptr(armpurview.EventStreamingTypeNone),
		// 				},
		// 				SystemData: &armpurview.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-17T04:23:24.157Z"); return t}()),
		// 					CreatedBy: to.Ptr("user"),
		// 					CreatedByType: to.Ptr(armpurview.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-17T04:23:24.157Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user"),
		// 					LastModifiedByType: to.Ptr(armpurview.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
