package armdynatrace_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dynatrace/armdynatrace/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/Monitors_CreateOrUpdate_MaximumSet_Gen.json
func ExampleMonitorsClient_BeginCreateOrUpdate_monitorsCreateOrUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdynatrace.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMonitorsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myMonitor", armdynatrace.MonitorResource{
		Location: to.Ptr("West US 2"),
		Tags: map[string]*string{
			"Environment": to.Ptr("Dev"),
		},
		Identity: &armdynatrace.IdentityProperties{
			Type: to.Ptr(armdynatrace.ManagedIdentityTypeSystemAssigned),
		},
		Properties: &armdynatrace.MonitorProperties{
			DynatraceEnvironmentProperties: &armdynatrace.EnvironmentProperties{
				AccountInfo:            &armdynatrace.AccountInfo{},
				EnvironmentInfo:        &armdynatrace.EnvironmentInfo{},
				SingleSignOnProperties: &armdynatrace.SingleSignOnProperties{},
			},
			LiftrResourceCategory:         to.Ptr(armdynatrace.LiftrResourceCategoriesUnknown),
			MarketplaceSubscriptionStatus: to.Ptr(armdynatrace.MarketplaceSubscriptionStatusActive),
			MonitoringStatus:              to.Ptr(armdynatrace.MonitoringStatusEnabled),
			PlanData: &armdynatrace.PlanData{
				BillingCycle:  to.Ptr("Monthly"),
				EffectiveDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-30T13:14:33.000Z"); return t }()),
				PlanDetails:   to.Ptr("dynatraceapitestplan"),
				UsageType:     to.Ptr("Committed"),
			},
			ProvisioningState: to.Ptr(armdynatrace.ProvisioningStateAccepted),
			UserInfo: &armdynatrace.UserInfo{
				Country:      to.Ptr("westus2"),
				EmailAddress: to.Ptr("alice@microsoft.com"),
				FirstName:    to.Ptr("Alice"),
				LastName:     to.Ptr("Bobab"),
				PhoneNumber:  to.Ptr("123456"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MonitorResource = armdynatrace.MonitorResource{
	// 	Name: to.Ptr("myMonitor"),
	// 	Type: to.Ptr("Dynatrace.Observability/monitors"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/monitors/myMonitor"),
	// 	Location: to.Ptr("West US 2"),
	// 	Tags: map[string]*string{
	// 		"Environment": to.Ptr("Dev"),
	// 	},
	// 	Identity: &armdynatrace.IdentityProperties{
	// 		Type: to.Ptr(armdynatrace.ManagedIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("4534676867978"),
	// 		TenantID: to.Ptr("23456789001"),
	// 	},
	// 	Properties: &armdynatrace.MonitorProperties{
	// 		DynatraceEnvironmentProperties: &armdynatrace.EnvironmentProperties{
	// 			AccountInfo: &armdynatrace.AccountInfo{
	// 				AccountID: to.Ptr("1234567890"),
	// 				RegionID: to.Ptr("wus2"),
	// 			},
	// 			EnvironmentInfo: &armdynatrace.EnvironmentInfo{
	// 				EnvironmentID: to.Ptr("a23xcv456"),
	// 				IngestionKey: to.Ptr("1234567890"),
	// 				LogsIngestionEndpoint: to.Ptr("https://dynatrace.com"),
	// 			},
	// 			SingleSignOnProperties: &armdynatrace.SingleSignOnProperties{
	// 				AADDomains: []*string{
	// 					to.Ptr("mpliftrdt20210811outlook.onmicrosoft.com")},
	// 					EnterpriseAppID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					SingleSignOnState: to.Ptr(armdynatrace.SingleSignOnStatesEnable),
	// 					SingleSignOnURL: to.Ptr("https://www.dynatrace.io/IAmSomeHash"),
	// 				},
	// 				UserID: to.Ptr("alice123"),
	// 			},
	// 			LiftrResourceCategory: to.Ptr(armdynatrace.LiftrResourceCategoriesUnknown),
	// 			LiftrResourcePreference: to.Ptr[int32](28),
	// 			MarketplaceSubscriptionStatus: to.Ptr(armdynatrace.MarketplaceSubscriptionStatusActive),
	// 			MonitoringStatus: to.Ptr(armdynatrace.MonitoringStatusEnabled),
	// 			PlanData: &armdynatrace.PlanData{
	// 				BillingCycle: to.Ptr("Monthly"),
	// 				EffectiveDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-30T13:14:33.000Z"); return t}()),
	// 				PlanDetails: to.Ptr("dynatraceapitestplan"),
	// 				UsageType: to.Ptr("Committed"),
	// 			},
	// 			ProvisioningState: to.Ptr(armdynatrace.ProvisioningStateSucceeded),
	// 			UserInfo: &armdynatrace.UserInfo{
	// 				Country: to.Ptr("westus2"),
	// 				EmailAddress: to.Ptr("alice@microsoft.com"),
	// 				FirstName: to.Ptr("Alice"),
	// 				LastName: to.Ptr("Bobab"),
	// 				PhoneNumber: to.Ptr("123456"),
	// 			},
	// 		},
	// 		SystemData: &armdynatrace.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-16T04:25:21.040Z"); return t}()),
	// 			CreatedBy: to.Ptr("alice@microsoft.com"),
	// 			CreatedByType: to.Ptr(armdynatrace.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-16T04:25:21.040Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("alice@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armdynatrace.CreatedByTypeUser),
	// 		},
	// 	}
}
