package armkoreaiartemis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/koreaiartemis/armkoreaiartemis"
)

// Generated from example definition: 2026-07-01-preview/Workspaces_GetDetails.json
func ExampleWorkspacesClient_GetDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkoreaiartemis.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspacesClient().GetDetails(ctx, "contoso-rg", "contoso-org", "contoso-prod", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armkoreaiartemis.WorkspacesClientGetDetailsResponse{
	// 	WorkspaceDetail: armkoreaiartemis.WorkspaceDetail{
	// 		ID: to.Ptr("ws_01JK2MXYZ9ABCDEF000002"),
	// 		AzureResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-rg/providers/Kore.AgentPlatform/organizations/contoso-org/workspaces/contoso-prod"),
	// 		OrganizationID: to.Ptr("org_01JK2MXYZ9ABCDEF123456"),
	// 		PartnerProperties: &armkoreaiartemis.WorkspaceDetailPartnerProperties{
	// 			Name: to.Ptr("Contoso Production"),
	// 			OwnerEmail: to.Ptr("owner@contoso.com"),
	// 			RetentionDays: to.Ptr[int32](90),
	// 			Default: to.Ptr(false),
	// 			MarketplaceDetails: []*armkoreaiartemis.WorkspaceMarketplaceDetails{
	// 				{
	// 					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-rg/providers/Microsoft.SaaS/resources/kore-ai-saas"),
	// 					MpID: to.Ptr("a1b2c3d4-e5f6-7890-abcd-ef1234567890"),
	// 					PlanID: to.Ptr("business-monthly"),
	// 					PlanName: to.Ptr("AI for Service"),
	// 					OfferID: to.Ptr("ai-for-service"),
	// 					PublisherID: to.Ptr("kore-ai"),
	// 					TermID: to.Ptr("1y"),
	// 					Status: to.Ptr(armkoreaiartemis.WorkspaceMarketplaceStatusSubscribed),
	// 				},
	// 			},
	// 		},
	// 		Status: to.Ptr(armkoreaiartemis.WorkspaceLifecycleStatusActive),
	// 		Projects: []*armkoreaiartemis.WorkspaceProject{
	// 			{
	// 				ID: to.Ptr("019e0218-aa9a-77c7-a9cc-812f3ec2fb72"),
	// 				Name: to.Ptr("AML False Positive Agent"),
	// 				Description: to.Ptr("AML false-positive review decision-support"),
	// 				CreatedBy: to.Ptr("jane.doe@contoso.com"),
	// 				Status: to.Ptr(armkoreaiartemis.WorkspaceProjectStatusDeployed),
	// 				AgentCount: to.Ptr[int32](9),
	// 				CreatedAt: to.Ptr(time.Date(2026, time.May, 7, 11, 0, 31, 514000000, time.UTC)),
	// 				UpdatedAt: to.Ptr(time.Date(2026, time.May, 7, 11, 0, 31, 798000000, time.UTC)),
	// 			},
	// 		},
	// 		CreatedAt: to.Ptr(time.Date(2026, time.January, 5, 9, 0, 0, 0, time.UTC)),
	// 		UpdatedAt: to.Ptr(time.Date(2026, time.April, 20, 14, 0, 0, 0, time.UTC)),
	// 	},
	// }
}
