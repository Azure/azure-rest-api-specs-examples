package armnewrelicobservability_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/newrelic/armnewrelicobservability"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/newrelic/resource-manager/NewRelic.Observability/stable/2022-07-01/examples/Monitors_SwitchBilling_MaximumSet_Gen.json
func ExampleMonitorsClient_SwitchBilling_monitorsSwitchBillingMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnewrelicobservability.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMonitorsClient().SwitchBilling(ctx, "rgNewRelic", "fhcjxnxumkdlgpwanewtkdnyuz", armnewrelicobservability.SwitchBillingRequest{
		AzureResourceID: to.Ptr("enfghpfw"),
		OrganizationID:  to.Ptr("k"),
		PlanData: &armnewrelicobservability.PlanData{
			BillingCycle:  to.Ptr(armnewrelicobservability.BillingCycleYEARLY),
			EffectiveDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-05T14:11:37.786Z"); return t }()),
			PlanDetails:   to.Ptr("tbbiaga"),
			UsageType:     to.Ptr(armnewrelicobservability.UsageTypePAYG),
		},
		UserEmail: to.Ptr("ruxvg@xqkmdhrnoo.hlmbpm"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NewRelicMonitorResource = armnewrelicobservability.NewRelicMonitorResource{
	// 	Name: to.Ptr("fteaqmtwspcfgyopqzrepiqu"),
	// 	Type: to.Ptr("hdj"),
	// 	ID: to.Ptr("ilrwkvbuwu"),
	// 	SystemData: &armnewrelicobservability.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-05T14:11:37.787Z"); return t}()),
	// 		CreatedBy: to.Ptr("pcdjzdldbwsdlfi"),
	// 		CreatedByType: to.Ptr(armnewrelicobservability.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-05T14:11:37.787Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("xbsjrxmwwlmpnpvcica"),
	// 		LastModifiedByType: to.Ptr(armnewrelicobservability.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("k"),
	// 	Tags: map[string]*string{
	// 		"key6976": to.Ptr("oaxfhf"),
	// 	},
	// 	Identity: &armnewrelicobservability.ManagedServiceIdentity{
	// 		Type: to.Ptr(armnewrelicobservability.ManagedServiceIdentityTypeNone),
	// 		PrincipalID: to.Ptr("cac47a92-2856-11ed-a261-0242ac120002"),
	// 		TenantID: to.Ptr("cac47a92-2856-11ed-a261-0242ac120002"),
	// 		UserAssignedIdentities: map[string]*armnewrelicobservability.UserAssignedIdentity{
	// 			"key8903": &armnewrelicobservability.UserAssignedIdentity{
	// 				ClientID: to.Ptr("cac47a92-2856-11ed-a261-0242ac120002"),
	// 				PrincipalID: to.Ptr("cac47a92-2856-11ed-a261-0242ac120002"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armnewrelicobservability.MonitorProperties{
	// 		AccountCreationSource: to.Ptr(armnewrelicobservability.AccountCreationSourceLIFTR),
	// 		LiftrResourceCategory: to.Ptr(armnewrelicobservability.LiftrResourceCategoriesUnknown),
	// 		LiftrResourcePreference: to.Ptr[int32](12),
	// 		MarketplaceSubscriptionID: to.Ptr("jizcsbgrdjhrfqqvvruhgftqhra"),
	// 		MarketplaceSubscriptionStatus: to.Ptr(armnewrelicobservability.MarketplaceSubscriptionStatusActive),
	// 		MonitoringStatus: to.Ptr(armnewrelicobservability.MonitoringStatusEnabled),
	// 		NewRelicAccountProperties: &armnewrelicobservability.AccountPropertiesForNewRelic{
	// 			AccountInfo: &armnewrelicobservability.AccountInfo{
	// 				AccountID: to.Ptr("xhqmg"),
	// 				Region: to.Ptr("ljcf"),
	// 			},
	// 			OrganizationInfo: &armnewrelicobservability.OrganizationInfo{
	// 				OrganizationID: to.Ptr("k"),
	// 			},
	// 			SingleSignOnProperties: &armnewrelicobservability.NewRelicSingleSignOnProperties{
	// 				EnterpriseAppID: to.Ptr("kwiwfz"),
	// 				ProvisioningState: to.Ptr(armnewrelicobservability.ProvisioningStateSucceeded),
	// 				SingleSignOnState: to.Ptr(armnewrelicobservability.SingleSignOnStatesInitial),
	// 				SingleSignOnURL: to.Ptr("kvseueuljsxmfwpqctz"),
	// 			},
	// 			UserID: to.Ptr("vcscxlncofcuduadesd"),
	// 		},
	// 		OrgCreationSource: to.Ptr(armnewrelicobservability.OrgCreationSourceLIFTR),
	// 		PlanData: &armnewrelicobservability.PlanData{
	// 			BillingCycle: to.Ptr(armnewrelicobservability.BillingCycleYEARLY),
	// 			EffectiveDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-05T14:11:37.786Z"); return t}()),
	// 			PlanDetails: to.Ptr("tbbiaga"),
	// 			UsageType: to.Ptr(armnewrelicobservability.UsageTypePAYG),
	// 		},
	// 		ProvisioningState: to.Ptr(armnewrelicobservability.ProvisioningStateSucceeded),
	// 		UserInfo: &armnewrelicobservability.UserInfo{
	// 			Country: to.Ptr("hslqnwdanrconqyekwbnttaetv"),
	// 			EmailAddress: to.Ptr("%6%@4-g.N1.3F-kI1.Ue-.lJso"),
	// 			FirstName: to.Ptr("vdftzcggirefejajwahhwhyibutramdaotvnuf"),
	// 			LastName: to.Ptr("bcsztgqovdlmzfkjdrngidwzqsevagexzzilnlc"),
	// 			PhoneNumber: to.Ptr("krf"),
	// 		},
	// 	},
	// }
}
