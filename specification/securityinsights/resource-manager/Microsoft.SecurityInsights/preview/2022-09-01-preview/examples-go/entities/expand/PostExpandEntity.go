package armsecurityinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/entities/expand/PostExpandEntity.json
func ExampleEntitiesClient_Expand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEntitiesClient().Expand(ctx, "myRg", "myWorkspace", "e1d3d618-e11f-478b-98e3-bb381539a8e1", armsecurityinsights.EntityExpandParameters{
		EndTime:     to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-26T00:00:00.000Z"); return t }()),
		ExpansionID: to.Ptr("a77992f3-25e9-4d01-99a4-5ff606cc410a"),
		StartTime:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-25T00:00:00.000Z"); return t }()),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EntityExpandResponse = armsecurityinsights.EntityExpandResponse{
	// 	MetaData: &armsecurityinsights.ExpansionResultsMetadata{
	// 		Aggregations: []*armsecurityinsights.ExpansionResultAggregation{
	// 			{
	// 				Count: to.Ptr[int32](1),
	// 				EntityKind: to.Ptr(armsecurityinsights.EntityKindAccount),
	// 		}},
	// 	},
	// 	Value: &armsecurityinsights.EntityExpandResponseValue{
	// 		Edges: []*armsecurityinsights.EntityEdges{
	// 			{
	// 				AdditionalData: map[string]any{
	// 					"EpochTimestamp": "1608289949",
	// 					"FirstSeen": "2021-09-01T11:12:29.597Z",
	// 					"Source": "Heartbeat",
	// 				},
	// 				TargetEntityID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entities/c1d60d86-5988-11eb-ae93-0242ac130002"),
	// 		}},
	// 		Entities: []armsecurityinsights.EntityClassification{
	// 			&armsecurityinsights.IPEntity{
	// 				Name: to.Ptr("e1d3d618-e11f-478b-98e3-bb381539a8e1"),
	// 				Type: to.Ptr("Microsoft.SecurityInsights/entities"),
	// 				ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entities/e1d3d618-e11f-478b-98e3-bb381539a8e1"),
	// 				Kind: to.Ptr(armsecurityinsights.EntityKindIP),
	// 				Properties: &armsecurityinsights.IPEntityProperties{
	// 					FriendlyName: to.Ptr("13.89.108.248"),
	// 					Address: to.Ptr("13.89.108.248"),
	// 				},
	// 		}},
	// 	},
	// }
}
