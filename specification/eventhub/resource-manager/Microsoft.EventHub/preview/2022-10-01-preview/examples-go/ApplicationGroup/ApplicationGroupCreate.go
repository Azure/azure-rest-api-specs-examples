package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-10-01-preview/examples/ApplicationGroup/ApplicationGroupCreate.json
func ExampleApplicationGroupClient_CreateOrUpdateApplicationGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationGroupClient().CreateOrUpdateApplicationGroup(ctx, "contosotest", "contoso-ua-test-eh-system-1", "appGroup1", armeventhub.ApplicationGroup{
		Properties: &armeventhub.ApplicationGroupProperties{
			ClientAppGroupIdentifier: to.Ptr("SASKeyName=KeyName"),
			IsEnabled:                to.Ptr(true),
			Policies: []armeventhub.ApplicationGroupPolicyClassification{
				&armeventhub.ThrottlingPolicy{
					Name:               to.Ptr("ThrottlingPolicy1"),
					Type:               to.Ptr(armeventhub.ApplicationGroupPolicyTypeThrottlingPolicy),
					MetricID:           to.Ptr(armeventhub.MetricIDIncomingMessages),
					RateLimitThreshold: to.Ptr[int64](7912),
				},
				&armeventhub.ThrottlingPolicy{
					Name:               to.Ptr("ThrottlingPolicy2"),
					Type:               to.Ptr(armeventhub.ApplicationGroupPolicyTypeThrottlingPolicy),
					MetricID:           to.Ptr(armeventhub.MetricIDIncomingBytes),
					RateLimitThreshold: to.Ptr[int64](3951729),
				},
				&armeventhub.ThrottlingPolicy{
					Name:               to.Ptr("ThrottlingPolicy3"),
					Type:               to.Ptr(armeventhub.ApplicationGroupPolicyTypeThrottlingPolicy),
					MetricID:           to.Ptr(armeventhub.MetricIDOutgoingBytes),
					RateLimitThreshold: to.Ptr[int64](245175),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ApplicationGroup = armeventhub.ApplicationGroup{
	// 	Name: to.Ptr("appGroup1"),
	// 	Type: to.Ptr("Microsoft.EventHub/Namespaces/ApplicationGroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosotest/providers/Microsoft.EventHub/namespaces/contoso-ua-test-eh-system-1/applicationgroups/appGroup1"),
	// 	Location: to.Ptr("EAST US 2 EUAP"),
	// 	Properties: &armeventhub.ApplicationGroupProperties{
	// 		ClientAppGroupIdentifier: to.Ptr("SASKeyName=KeyName"),
	// 		IsEnabled: to.Ptr(true),
	// 		Policies: []armeventhub.ApplicationGroupPolicyClassification{
	// 			&armeventhub.ThrottlingPolicy{
	// 				Name: to.Ptr("ThrottlingPolicy1"),
	// 				Type: to.Ptr(armeventhub.ApplicationGroupPolicyTypeThrottlingPolicy),
	// 				MetricID: to.Ptr(armeventhub.MetricIDIncomingMessages),
	// 				RateLimitThreshold: to.Ptr[int64](7912),
	// 			},
	// 			&armeventhub.ThrottlingPolicy{
	// 				Name: to.Ptr("ThrottlingPolicy2"),
	// 				Type: to.Ptr(armeventhub.ApplicationGroupPolicyTypeThrottlingPolicy),
	// 				MetricID: to.Ptr(armeventhub.MetricIDIncomingBytes),
	// 				RateLimitThreshold: to.Ptr[int64](3951729),
	// 			},
	// 			&armeventhub.ThrottlingPolicy{
	// 				Name: to.Ptr("ThrottlingPolicy3"),
	// 				Type: to.Ptr(armeventhub.ApplicationGroupPolicyTypeThrottlingPolicy),
	// 				MetricID: to.Ptr(armeventhub.MetricIDOutgoingBytes),
	// 				RateLimitThreshold: to.Ptr[int64](245175),
	// 		}},
	// 	},
	// }
}
