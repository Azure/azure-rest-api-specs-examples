package armnewrelicobservability_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/newrelic/armnewrelicobservability/v2"
)

// Generated from example definition: 2025-05-01-preview/Monitors_ListBySubscription_MaximumSet_Gen.json
func ExampleMonitorsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnewrelicobservability.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMonitorsClient().NewListBySubscriptionPager(nil)
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
		// page = armnewrelicobservability.MonitorsClientListBySubscriptionResponse{
		// 	NewRelicMonitorResourceListResult: armnewrelicobservability.NewRelicMonitorResourceListResult{
		// 		NextLink: to.Ptr("http://testlink.com"),
		// 		Value: []*armnewrelicobservability.NewRelicMonitorResource{
		// 			{
		// 				Name: to.Ptr("fteaqmtwspcfgyopqzrepiqu"),
		// 				Type: to.Ptr("hdj"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/NewRelic.Observability/monitors/myMonitor"),
		// 				Identity: &armnewrelicobservability.ManagedServiceIdentity{
		// 					Type: to.Ptr(armnewrelicobservability.ManagedServiceIdentityTypeNone),
		// 					PrincipalID: to.Ptr("cac47a92-2856-11ed-a261-0242ac120002"),
		// 					TenantID: to.Ptr("cac47a92-2856-11ed-a261-0242ac120002"),
		// 					UserAssignedIdentities: map[string]*armnewrelicobservability.UserAssignedIdentity{
		// 						"key8903": &armnewrelicobservability.UserAssignedIdentity{
		// 							ClientID: to.Ptr("cac47a92-2856-11ed-a261-0242ac120002"),
		// 							PrincipalID: to.Ptr("cac47a92-2856-11ed-a261-0242ac120002"),
		// 						},
		// 					},
		// 				},
		// 				Location: to.Ptr("k"),
		// 				Properties: &armnewrelicobservability.MonitorProperties{
		// 					AccountCreationSource: to.Ptr(armnewrelicobservability.AccountCreationSourceLIFTR),
		// 					LiftrResourceCategory: to.Ptr(armnewrelicobservability.LiftrResourceCategoriesUnknown),
		// 					LiftrResourcePreference: to.Ptr[int32](12),
		// 					MarketplaceSubscriptionID: to.Ptr("jizcsbgrdjhrfqqvvruhgftqhra"),
		// 					MarketplaceSubscriptionStatus: to.Ptr(armnewrelicobservability.MarketplaceSubscriptionStatusActive),
		// 					MonitoringStatus: to.Ptr(armnewrelicobservability.MonitoringStatusEnabled),
		// 					NewRelicAccountProperties: &armnewrelicobservability.AccountPropertiesForNewRelic{
		// 						AccountInfo: &armnewrelicobservability.AccountInfo{
		// 							AccountID: to.Ptr("xhqmg"),
		// 							Region: to.Ptr("ljcf"),
		// 						},
		// 						OrganizationInfo: &armnewrelicobservability.OrganizationInfo{
		// 							OrganizationID: to.Ptr("k"),
		// 						},
		// 						SingleSignOnProperties: &armnewrelicobservability.NewRelicSingleSignOnProperties{
		// 							EnterpriseAppID: to.Ptr("kwiwfz"),
		// 							ProvisioningState: to.Ptr(armnewrelicobservability.ProvisioningStateAccepted),
		// 							SingleSignOnState: to.Ptr(armnewrelicobservability.SingleSignOnStatesInitial),
		// 							SingleSignOnURL: to.Ptr("kvseueuljsxmfwpqctz"),
		// 						},
		// 						UserID: to.Ptr("vcscxlncofcuduadesd"),
		// 					},
		// 					OrgCreationSource: to.Ptr(armnewrelicobservability.OrgCreationSourceLIFTR),
		// 					PlanData: &armnewrelicobservability.PlanData{
		// 						BillingCycle: to.Ptr("Yearly"),
		// 						EffectiveDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-05T14:11:37.786Z"); return t}()),
		// 						PlanDetails: to.Ptr("tbbiaga"),
		// 						UsageType: to.Ptr(armnewrelicobservability.UsageTypePAYG),
		// 					},
		// 					ProvisioningState: to.Ptr(armnewrelicobservability.ProvisioningStateAccepted),
		// 					UserInfo: &armnewrelicobservability.UserInfo{
		// 						Country: to.Ptr("hslqnwdanrconqyekwbnttaetv"),
		// 						EmailAddress: to.Ptr("%6%@4-g.N1.3F-kI1.Ue-.lJso"),
		// 						FirstName: to.Ptr("vdftzcggirefejajwahhwhyibutramdaotvnuf"),
		// 						LastName: to.Ptr("bcsztgqovdlmzfkjdrngidwzqsevagexzzilnlc"),
		// 						PhoneNumber: to.Ptr("krf"),
		// 					},
		// 				},
		// 				SystemData: &armnewrelicobservability.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-05T14:11:37.787Z"); return t}()),
		// 					CreatedBy: to.Ptr("pcdjzdldbwsdlfi"),
		// 					CreatedByType: to.Ptr(armnewrelicobservability.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-05T14:11:37.787Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("xbsjrxmwwlmpnpvcica"),
		// 					LastModifiedByType: to.Ptr(armnewrelicobservability.CreatedByTypeUser),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key6976": to.Ptr("oaxfhf"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
