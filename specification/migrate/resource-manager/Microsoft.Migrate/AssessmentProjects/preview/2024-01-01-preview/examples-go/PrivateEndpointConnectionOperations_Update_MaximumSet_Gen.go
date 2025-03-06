package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/PrivateEndpointConnectionOperations_Update_MaximumSet_Gen.json
func ExamplePrivateEndpointConnectionOperationsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionOperationsClient().BeginUpdate(ctx, "sakanwar", "sakanwar1204project", "sakanwar1204project1634pe.bf42f8a1-09f5-4ee4-aea6-a019cc60f9d7", armmigrationassessment.PrivateEndpointConnection{
		Properties: &armmigrationassessment.PrivateEndpointConnectionProperties{
			PrivateEndpoint: &armmigrationassessment.PrivateEndpoint{},
			PrivateLinkServiceConnectionState: &armmigrationassessment.PrivateLinkServiceConnectionState{
				ActionsRequired: to.Ptr(""),
				Status:          to.Ptr(armmigrationassessment.PrivateEndpointServiceConnectionStatusApproved),
			},
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
	// res.PrivateEndpointConnection = armmigrationassessment.PrivateEndpointConnection{
	// 	Name: to.Ptr("sakanwar1204project1634pe"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentprojects/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/31be0ff4-c932-4cb3-8efc-efa411d79280/resourceGroups/PrivLink-SelfHost/providers/Microsoft.Migrate/assessmentprojects/PiyushApp15328project/privateEndpointConnections/sakanwar1204project1634pe"),
	// 	SystemData: &armmigrationassessment.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-15T07:22:18.589Z"); return t}()),
	// 		CreatedBy: to.Ptr("sakanwar"),
	// 		CreatedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-15T07:22:18.589Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("sakanwar"),
	// 		LastModifiedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmigrationassessment.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armmigrationassessment.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/31be0ff4-c932-4cb3-8efc-efa411d79280/resourceGroups/PrivLink-SelfHost/providers/Microsoft.Network/privateEndpoints/sakanwar1204project1634pe"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armmigrationassessment.PrivateLinkServiceConnectionState{
	// 			ActionsRequired: to.Ptr(""),
	// 			Status: to.Ptr(armmigrationassessment.PrivateEndpointServiceConnectionStatusApproved),
	// 		},
	// 		ProvisioningState: to.Ptr(armmigrationassessment.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 	},
	// }
}
