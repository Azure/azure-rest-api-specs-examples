package armmanagedapplications_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/solutions/armmanagedapplications/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5fb045bd44f143bae17da2e01552ae531f77d0ba/specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/createOrUpdateJitRequest.json
func ExampleJitRequestsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedapplications.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewJitRequestsClient().BeginCreateOrUpdate(ctx, "rg", "myJitRequest", armmanagedapplications.JitRequestDefinition{
		Properties: &armmanagedapplications.JitRequestProperties{
			ApplicationResourceID: to.Ptr("/subscriptions/00c76877-e316-48a7-af60-4a09fec9d43f/resourceGroups/52F30DB2/providers/Microsoft.Solutions/applications/7E193158"),
			JitAuthorizationPolicies: []*armmanagedapplications.JitAuthorizationPolicies{
				{
					PrincipalID:      to.Ptr("1db8e132e2934dbcb8e1178a61319491"),
					RoleDefinitionID: to.Ptr("ecd05a23-931a-4c38-a52b-ac7c4c583334"),
				}},
			JitSchedulingPolicy: &armmanagedapplications.JitSchedulingPolicy{
				Type:      to.Ptr(armmanagedapplications.JitSchedulingTypeOnce),
				Duration:  to.Ptr("PT8H"),
				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-22T05:48:30.666Z"); return t }()),
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
	// res.JitRequestDefinition = armmanagedapplications.JitRequestDefinition{
	// 	Name: to.Ptr("myJitRequest"),
	// 	Type: to.Ptr("Microsoft.Solutions/jitRequests"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/jitRequests/myJitRequest"),
	// 	Properties: &armmanagedapplications.JitRequestProperties{
	// 		ApplicationResourceID: to.Ptr("/subscriptions/00c76877-e316-48a7-af60-4a09fec9d43f/resourceGroups/52F30DB2/providers/Microsoft.Solutions/applications/7E193158"),
	// 		CreatedBy: &armmanagedapplications.ApplicationClientDetails{
	// 			ApplicationID: to.Ptr("33a83f1f-c363-4ae7-9e0a-a0c08466354d"),
	// 			Oid: to.Ptr(""),
	// 		},
	// 		JitAuthorizationPolicies: []*armmanagedapplications.JitAuthorizationPolicies{
	// 			{
	// 				PrincipalID: to.Ptr("1db8e132e2934dbcb8e1178a61319491"),
	// 				RoleDefinitionID: to.Ptr("ecd05a23-931a-4c38-a52b-ac7c4c583334"),
	// 		}},
	// 		JitRequestState: to.Ptr(armmanagedapplications.JitRequestStatePending),
	// 		JitSchedulingPolicy: &armmanagedapplications.JitSchedulingPolicy{
	// 			Type: to.Ptr(armmanagedapplications.JitSchedulingTypeOnce),
	// 			Duration: to.Ptr("PT8H"),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-22T05:48:30.666Z"); return t}()),
	// 		},
	// 		ProvisioningState: to.Ptr(armmanagedapplications.ProvisioningStateSucceeded),
	// 		UpdatedBy: &armmanagedapplications.ApplicationClientDetails{
	// 			ApplicationID: to.Ptr("33a83f1f-c363-4ae7-9e0a-a0c08466354d"),
	// 			Oid: to.Ptr(""),
	// 		},
	// 	},
	// }
}
