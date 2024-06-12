package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/entities/GetEntities.json
func ExampleEntitiesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEntitiesClient().NewListPager("myRg", "myWorkspace", nil)
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
		// page.EntityList = armsecurityinsights.EntityList{
		// 	Value: []armsecurityinsights.EntityClassification{
		// 		&armsecurityinsights.AccountEntity{
		// 			Name: to.Ptr("e1d3d618-e11f-478b-98e3-bb381539a8e1"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/entities"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entities/e1d3d618-e11f-478b-98e3-bb381539a8e1"),
		// 			Kind: to.Ptr(armsecurityinsights.EntityKindAccount),
		// 			Properties: &armsecurityinsights.AccountEntityProperties{
		// 				FriendlyName: to.Ptr("administrator"),
		// 				AADTenantID: to.Ptr("70fbdad0-7441-4564-b2b5-2b8862d0fee0"),
		// 				AADUserID: to.Ptr("f7033626-2572-46b1-bba0-06646f4f95b3"),
		// 				AccountName: to.Ptr("administrator"),
		// 				IsDomainJoined: to.Ptr(true),
		// 				NtDomain: to.Ptr("domain"),
		// 				ObjectGUID: to.Ptr("11227b78-3c6e-436e-a2a2-02fc7662eca0"),
		// 				Puid: to.Ptr("ee3cb2d8-14ba-45ef-8009-d6f1cacfa04d"),
		// 				Sid: to.Ptr("S-1-5-18"),
		// 				UpnSuffix: to.Ptr("contoso"),
		// 			},
		// 		},
		// 		&armsecurityinsights.HostEntity{
		// 			Name: to.Ptr("fed9fe89-dce8-40f2-bf44-70f23fe93b3c"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/entities"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entities/fed9fe89-dce8-40f2-bf44-70f23fe93b3c"),
		// 			Kind: to.Ptr(armsecurityinsights.EntityKindHost),
		// 			Properties: &armsecurityinsights.HostEntityProperties{
		// 				FriendlyName: to.Ptr("vm1"),
		// 				AzureID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachines/vm1"),
		// 				DNSDomain: to.Ptr("contoso"),
		// 				HostName: to.Ptr("vm1"),
		// 				IsDomainJoined: to.Ptr(true),
		// 				NetBiosName: to.Ptr("contoso"),
		// 				NtDomain: to.Ptr("domain"),
		// 				OmsAgentID: to.Ptr("70fbdad0-7441-4564-b2b5-2b8862d0fee0"),
		// 				OSFamily: to.Ptr(armsecurityinsights.OSFamilyWindows),
		// 				OSVersion: to.Ptr("1.0"),
		// 			},
		// 		},
		// 		&armsecurityinsights.FileEntity{
		// 			Name: to.Ptr("af378b21-b4aa-4fe7-bc70-13f8621a322f"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/entities"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entities/af378b21-b4aa-4fe7-bc70-13f8621a322f"),
		// 			Kind: to.Ptr(armsecurityinsights.EntityKindFile),
		// 			Properties: &armsecurityinsights.FileEntityProperties{
		// 				FriendlyName: to.Ptr("cmd.exe"),
		// 				Directory: to.Ptr("C:\\Windows\\System32"),
		// 				FileName: to.Ptr("cmd.exe"),
		// 			},
		// 	}},
		// }
	}
}
