package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c548b0bd279f5e233661b1c81fb5b61b19965cd/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/PolicyEvents_QuerySubscriptionLevelPolicyDefinitionScopeNextLink.json
func ExamplePolicyEventsClient_NewListQueryResultsForPolicyDefinitionPager_queryAtSubscriptionLevelPolicyDefinitionScopeWithNextLink() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPolicyEventsClient().NewListQueryResultsForPolicyDefinitionPager(armpolicyinsights.PolicyEventsResourceTypeDefault, "fffedd8f-ffff-fffd-fffd-fffed2f84852", "24813039-7534-408a-9842-eb99f45721b1", &armpolicyinsights.QueryOptions{Top: nil,
		Filter:    nil,
		OrderBy:   nil,
		Select:    nil,
		From:      nil,
		To:        nil,
		Apply:     nil,
		SkipToken: to.Ptr("WpmWfBSvPhkAK6QD"),
		Expand:    nil,
	}, nil)
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
		// page.PolicyEventsQueryResults = armpolicyinsights.PolicyEventsQueryResults{
		// 	ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyDefinitions/24813039-7534-408a-9842-eb99f45721b1/providers/Microsoft.PolicyInsights/policyEvents/$metadata#default"),
		// 	ODataCount: to.Ptr[int32](2),
		// 	Value: []*armpolicyinsights.PolicyEvent{
		// 		{
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyDefinitions/24813039-7534-408a-9842-eb99f45721b1/providers/Microsoft.PolicyInsights/policyEvents/$metadata#default/$entity"),
		// 			ComplianceState: to.Ptr("NonCompliant"),
		// 			IsCompliant: to.Ptr(false),
		// 			ManagementGroupIDs: to.Ptr("mymg,fff988bf-fff1-ffff-fffb-fffcd011db47"),
		// 			PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyAssignments/f4cc58b7db524a9799381531"),
		// 			PolicyAssignmentName: to.Ptr("f4cc58b7db524a9799381531"),
		// 			PolicyAssignmentOwner: to.Ptr("tbd"),
		// 			PolicyAssignmentParameters: to.Ptr("{\"TAGNAME_1\":{\"value\":\"NA\"}}"),
		// 			PolicyAssignmentScope: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852"),
		// 			PolicyDefinitionAction: to.Ptr("audit"),
		// 			PolicyDefinitionCategory: to.Ptr("tbd"),
		// 			PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyDefinitions/24813039-7534-408a-9842-eb99f45721b1"),
		// 			PolicyDefinitionName: to.Ptr("24813039-7534-408a-9842-eb99f45721b1"),
		// 			PolicyDefinitionReferenceID: to.Ptr("14799174781370023846"),
		// 			PolicySetDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policySetDefinitions/12b58873-e0f8-4b95-936c-86cbe7c9d697"),
		// 			PolicySetDefinitionName: to.Ptr("12b58873-e0f8-4b95-936c-86cbe7c9d697"),
		// 			PrincipalOid: to.Ptr("fff890fa-fff0-fff3-fff9-fffd7653f078"),
		// 			ResourceGroup: to.Ptr("myResourceGroup"),
		// 			ResourceID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myResourceGroup/providers/Microsoft.ClassicCompute/domainNames/myDomainName/deploymentSlots/production/state/start"),
		// 			ResourceLocation: to.Ptr("eastus"),
		// 			ResourceTags: to.Ptr("tbd"),
		// 			ResourceType: to.Ptr("/Microsoft.ClassicCompute/domainNames/deploymentSlots/state"),
		// 			SubscriptionID: to.Ptr("fffedd8f-ffff-fffd-fffd-fffed2f84852"),
		// 			TenantID: to.Ptr("fff988bf-fff1-ffff-fffb-fffcd011db47"),
		// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-08T19:58:11.590Z"); return t}()),
		// 		},
		// 		{
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyDefinitions/24813039-7534-408a-9842-eb99f45721b1/providers/Microsoft.PolicyInsights/policyEvents/$metadata#default/$entity"),
		// 			ComplianceState: to.Ptr("NonCompliant"),
		// 			IsCompliant: to.Ptr(false),
		// 			ManagementGroupIDs: to.Ptr("mymg,fff988bf-fff1-ffff-fffb-fffcd011db47"),
		// 			PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyAssignments/f4cc58b7db524a9799381531"),
		// 			PolicyAssignmentName: to.Ptr("f4cc58b7db524a9799381531"),
		// 			PolicyAssignmentOwner: to.Ptr("tbd"),
		// 			PolicyAssignmentParameters: to.Ptr("{\"TAGNAME_1\":{\"value\":\"NA\"}}"),
		// 			PolicyAssignmentScope: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852"),
		// 			PolicyDefinitionAction: to.Ptr("audit"),
		// 			PolicyDefinitionCategory: to.Ptr("tbd"),
		// 			PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyDefinitions/24813039-7534-408a-9842-eb99f45721b1"),
		// 			PolicyDefinitionName: to.Ptr("24813039-7534-408a-9842-eb99f45721b1"),
		// 			PolicyDefinitionReferenceID: to.Ptr("1679708035638239273"),
		// 			PolicySetDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policySetDefinitions/12b58873-e0f8-4b95-936c-86cbe7c9d697"),
		// 			PolicySetDefinitionName: to.Ptr("12b58873-e0f8-4b95-936c-86cbe7c9d697"),
		// 			PrincipalOid: to.Ptr("fff890fa-fff0-fff3-fff9-fffd7653f078"),
		// 			ResourceGroup: to.Ptr("myResourceGroup"),
		// 			ResourceID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myResourceGroup/providers/Microsoft.ClassicCompute/domainNames/myDomainName/deploymentSlots/production/state/start"),
		// 			ResourceLocation: to.Ptr("eastus"),
		// 			ResourceTags: to.Ptr("tbd"),
		// 			ResourceType: to.Ptr("/Microsoft.ClassicCompute/domainNames/deploymentSlots/state"),
		// 			SubscriptionID: to.Ptr("fffedd8f-ffff-fffd-fffd-fffed2f84852"),
		// 			TenantID: to.Ptr("fff988bf-fff1-ffff-fffb-fffcd011db47"),
		// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-08T19:58:11.590Z"); return t}()),
		// 	}},
		// }
	}
}
