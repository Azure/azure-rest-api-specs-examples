package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/entities/GetMailMessageEntityById.json
func ExampleEntitiesClient_Get_getAMailMessageEntity() {
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
	// 	                            EntityClassification: &armsecurityinsights.MailMessageEntity{
	// 		Name: to.Ptr("e1d3d618-e11f-478b-98e3-bb381539a8e1"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/entities"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entities/e1d3d618-e11f-478b-98e3-bb381539a8e1"),
	// 		Kind: to.Ptr(armsecurityinsights.EntityKindMailMessage),
	// 		Properties: &armsecurityinsights.MailMessageEntityProperties{
	// 			FriendlyName: to.Ptr("cmd.exe"),
	// 			DeliveryAction: to.Ptr(armsecurityinsights.DeliveryActionBlocked),
	// 			FileEntityIDs: []*string{
	// 				to.Ptr("ccfce855-e02f-491b-a1cc-5bafb371ad0c")},
	// 				InternetMessageID: to.Ptr("message id"),
	// 				P1Sender: to.Ptr("email@fake.com"),
	// 				P1SenderDisplayName: to.Ptr("p1 sender display name"),
	// 				P1SenderDomain: to.Ptr("p1 sender domain"),
	// 				P2Sender: to.Ptr("the sender"),
	// 				P2SenderDisplayName: to.Ptr("p2 sender display name"),
	// 				P2SenderDomain: to.Ptr("p2 Sender Domain"),
	// 				Recipient: to.Ptr("recipient"),
	// 				SenderIP: to.Ptr("1.23.34.43"),
	// 				Subject: to.Ptr("subject"),
	// 				ThreatDetectionMethods: []*string{
	// 					to.Ptr("thrreat1"),
	// 					to.Ptr("thread2")},
	// 					Threats: []*string{
	// 						to.Ptr("thrreat1"),
	// 						to.Ptr("thread2")},
	// 						Urls: []*string{
	// 							to.Ptr("http://moqbrarcwmnk.banxhdcojlg.biz")},
	// 							Language: to.Ptr("language"),
	// 						},
	// 					},
	// 					                        }
}
