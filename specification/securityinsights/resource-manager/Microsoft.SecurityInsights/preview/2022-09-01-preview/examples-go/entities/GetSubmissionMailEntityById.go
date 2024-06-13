package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/entities/GetSubmissionMailEntityById.json
func ExampleEntitiesClient_Get_getASubmissionMailEntity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEntitiesClient().Get(ctx, "myRg", "myWorkspace", "e1d3d618-e11f-478b-98e3-bb381539a8e1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.EntitiesClientGetResponse{
	// 	                            EntityClassification: &armsecurityinsights.SubmissionMailEntity{
	// 		Name: to.Ptr("e1d3d618-e11f-478b-98e3-bb381539a8e1"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/entities"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entities/e1d3d618-e11f-478b-98e3-bb381539a8e1"),
	// 		Kind: to.Ptr(armsecurityinsights.EntityKindSubmissionMail),
	// 		Properties: &armsecurityinsights.SubmissionMailEntityProperties{
	// 			FriendlyName: to.Ptr("recipient"),
	// 			Recipient: to.Ptr("recipient"),
	// 			ReportType: to.Ptr("report type"),
	// 			Sender: to.Ptr("sender"),
	// 			SenderIP: to.Ptr("1.4.35.34"),
	// 			Subject: to.Ptr("subject"),
	// 			SubmissionID: to.Ptr("5bb3d8fe-54bc-499c-bc21-86fe8df2a184"),
	// 			Submitter: to.Ptr("submitter"),
	// 		},
	// 	},
	// 	                        }
}
