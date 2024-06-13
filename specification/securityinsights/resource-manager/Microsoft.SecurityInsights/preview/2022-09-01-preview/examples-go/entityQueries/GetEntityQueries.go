package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/entityQueries/GetEntityQueries.json
func ExampleEntityQueriesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEntityQueriesClient().NewListPager("myRg", "myWorkspace", &armsecurityinsights.EntityQueriesClientListOptions{Kind: to.Ptr(armsecurityinsights.Enum13Expansion)})
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
		// page.EntityQueryList = armsecurityinsights.EntityQueryList{
		// 	Value: []armsecurityinsights.EntityQueryClassification{
		// 		&armsecurityinsights.ExpansionEntityQuery{
		// 			Name: to.Ptr("37ca3555-c135-4a73-a65e-9c1d00323f5d"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/entityQueries"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entityQueries/37ca3555-c135-4a73-a65e-9c1d00323f5d"),
		// 			Kind: to.Ptr(armsecurityinsights.EntityQueryKindExpansion),
		// 			Properties: &armsecurityinsights.ExpansionEntityQueriesProperties{
		// 				DataSources: []*string{
		// 					to.Ptr("AzureActivity")},
		// 					DisplayName: to.Ptr("Least active accounts on Azure from this IP"),
		// 					InputEntityType: to.Ptr(armsecurityinsights.EntityTypeIP),
		// 					InputFields: []*string{
		// 						to.Ptr("address")},
		// 						OutputEntityTypes: []*armsecurityinsights.EntityType{
		// 							to.Ptr(armsecurityinsights.EntityTypeAccount)},
		// 							QueryTemplate: to.Ptr("let AccountActivity_byIP = (v_IP_Address:string){\r\n                            AzureActivity\r\n                            | where Caller != '' and CallerIpAddress == v_IP_Address\r\n                            | summarize Account_Aux_StartTime = min(TimeGenerated), Account_Aux_EndTime = max(TimeGenerated), Count = count() by Caller, TenantId\r\n                            | top 10 by Count asc nulls last \r\n                            | extend UPN = iff(Caller contains '@', Caller, ''), Account_AadUserId = iff(Caller !contains '@', Caller,'')\r\n                            | extend Account_Name = split(UPN,'@')[0] , Account_UPNSuffix = split(UPN,'@')[1]\r\n                            | project Account_Name, Account_UPNSuffix, Account_AadUserId, Account_AadTenantId=TenantId, Account_Aux_StartTime , Account_Aux_EndTime};\r\n                            AccountActivity_byIP('<address>')"),
		// 						},
		// 					},
		// 					&armsecurityinsights.ExpansionEntityQuery{
		// 						Name: to.Ptr("97a1d515-abf2-4231-9a35-985f9de0bb91"),
		// 						Type: to.Ptr("Microsoft.SecurityInsights/entityQueries"),
		// 						ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entityQueries/97a1d515-abf2-4231-9a35-985f9de0bb91"),
		// 						Kind: to.Ptr(armsecurityinsights.EntityQueryKindExpansion),
		// 						Properties: &armsecurityinsights.ExpansionEntityQueriesProperties{
		// 							DataSources: []*string{
		// 								to.Ptr("AzureActivity")},
		// 								DisplayName: to.Ptr("Most active accounts on Azure from this IP"),
		// 								InputEntityType: to.Ptr(armsecurityinsights.EntityTypeIP),
		// 								InputFields: []*string{
		// 									to.Ptr("address")},
		// 									OutputEntityTypes: []*armsecurityinsights.EntityType{
		// 										to.Ptr(armsecurityinsights.EntityTypeAccount)},
		// 										QueryTemplate: to.Ptr("let AccountActivity_byIP = (v_IP_Address:string){\r\n                            AzureActivity\r\n                            | where Caller != '' and CallerIpAddress == v_IP_Address\r\n                            | summarize Account_Aux_StartTime = min(TimeGenerated), Account_Aux_EndTime = max(TimeGenerated), Count = count() by Caller, TenantId\r\n                            | top 10 by Count desc nulls last \r\n                            | extend UPN = iff(Caller contains '@', Caller, ''), Account_AadUserId = iff(Caller !contains '@', Caller,'')\r\n                            | extend Account_Name = split(UPN,'@')[0] , Account_UPNSuffix = split(UPN,'@')[1]\r\n                            | project Account_Name, Account_UPNSuffix, Account_AadUserId, Account_AadTenantId=TenantId, Account_Aux_StartTime , Account_Aux_EndTime};\r\n                            AccountActivity_byIP('<address>')"),
		// 									},
		// 							}},
		// 						}
	}
}
