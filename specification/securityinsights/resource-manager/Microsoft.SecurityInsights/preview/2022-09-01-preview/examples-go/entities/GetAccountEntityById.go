package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/entities/GetAccountEntityById.json
func ExampleEntitiesClient_Get_getAnAccountEntity() {
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
	// 	                            EntityClassification: &armsecurityinsights.AccountEntity{
	// 		Name: to.Ptr("e1d3d618-e11f-478b-98e3-bb381539a8e1"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/entities"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entities/e1d3d618-e11f-478b-98e3-bb381539a8e1"),
	// 		Kind: to.Ptr(armsecurityinsights.EntityKindAccount),
	// 		Properties: &armsecurityinsights.AccountEntityProperties{
	// 			FriendlyName: to.Ptr("administrator"),
	// 			AADTenantID: to.Ptr("70fbdad0-7441-4564-b2b5-2b8862d0fee0"),
	// 			AADUserID: to.Ptr("f7033626-2572-46b1-bba0-06646f4f95b3"),
	// 			AccountName: to.Ptr("administrator"),
	// 			DNSDomain: to.Ptr("contoso.com"),
	// 			IsDomainJoined: to.Ptr(true),
	// 			NtDomain: to.Ptr("domain"),
	// 			ObjectGUID: to.Ptr("11227b78-3c6e-436e-a2a2-02fc7662eca0"),
	// 			Puid: to.Ptr("ee3cb2d8-14ba-45ef-8009-d6f1cacfa04d"),
	// 			Sid: to.Ptr("S-1-5-18"),
	// 			UpnSuffix: to.Ptr("contoso"),
	// 		},
	// 	},
	// 	                        }
}
