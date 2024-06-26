package armrecoveryservicesdatareplication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/ProtectedItem_List.json
func ExampleProtectedItemClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProtectedItemClient().NewListPager("rgrecoveryservicesdatareplication", "4", nil)
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
		// page.ProtectedItemModelCollection = armrecoveryservicesdatareplication.ProtectedItemModelCollection{
		// 	Value: []*armrecoveryservicesdatareplication.ProtectedItemModel{
		// 		{
		// 			Name: to.Ptr("t"),
		// 			Type: to.Ptr("xlyjashandpfwivuipoplgkgsnwoh"),
		// 			ID: to.Ptr("egmhsfbgkarlobrgybkz"),
		// 			Properties: &armrecoveryservicesdatareplication.ProtectedItemModelProperties{
		// 				AllowedJobs: []*string{
		// 					to.Ptr("guryeoocjbvqvalfkrxecpxwynpxs")},
		// 					CorrelationID: to.Ptr("mvxvtcqwgp"),
		// 					CurrentJob: &armrecoveryservicesdatareplication.ProtectedItemModelPropertiesCurrentJob{
		// 						Name: to.Ptr("kqtvbrfmqaxdgpttkbmzpwafjp"),
		// 						DisplayName: to.Ptr("awutlqrisstqb"),
		// 						EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:55.456Z"); return t}()),
		// 						ID: to.Ptr("bnmbzxzyfgwh"),
		// 						ScenarioName: to.Ptr("ljbnhbdmreowdqnlcqycvaramwuii"),
		// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:55.456Z"); return t}()),
		// 						State: to.Ptr("ztlpngveoqcdejpwaiudhrioskauqv"),
		// 					},
		// 					CustomProperties: &armrecoveryservicesdatareplication.ProtectedItemModelCustomProperties{
		// 						InstanceType: to.Ptr("ProtectedItemModelCustomProperties"),
		// 					},
		// 					DraID: to.Ptr("vxrmsufvxothxauhvqdowascmy"),
		// 					FabricID: to.Ptr("ebsxoblmhlhqjzzjzdwo"),
		// 					FabricObjectID: to.Ptr("kjcizdpahzqsrwyiywbhyzdxsufj"),
		// 					FabricObjectName: to.Ptr("glrjwtvmejxuagjepcwaxhih"),
		// 					HealthErrors: []*armrecoveryservicesdatareplication.HealthErrorModel{
		// 						{
		// 							AffectedResourceCorrelationIDs: []*string{
		// 								to.Ptr("fope")},
		// 								AffectedResourceType: to.Ptr("scfniv"),
		// 								Category: to.Ptr("leigw"),
		// 								Causes: to.Ptr("xznphqrrmsdzm"),
		// 								ChildErrors: []*armrecoveryservicesdatareplication.InnerHealthErrorModel{
		// 									{
		// 										Category: to.Ptr("lcsdxrqxquke"),
		// 										Causes: to.Ptr("kefaugkpxjkpulimjthjnl"),
		// 										Code: to.Ptr("yuxxpblihirpedwkigywgwjjrlzq"),
		// 										CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:52.128Z"); return t}()),
		// 										HealthCategory: to.Ptr("mhdgfjqwbikhxmhtomkl"),
		// 										IsCustomerResolvable: to.Ptr(true),
		// 										Message: to.Ptr("sskcei"),
		// 										Recommendation: to.Ptr("kqybwaesqumywtjepi"),
		// 										Severity: to.Ptr("wqxxiuaqjyagq"),
		// 										Source: to.Ptr("wevvftugwydzzw"),
		// 										Summary: to.Ptr("djsmgrltruljo"),
		// 								}},
		// 								Code: to.Ptr("dgxkefzmeukd"),
		// 								CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:52.128Z"); return t}()),
		// 								HealthCategory: to.Ptr("itc"),
		// 								IsCustomerResolvable: to.Ptr(true),
		// 								Message: to.Ptr("lbywtdprdqdekl"),
		// 								Recommendation: to.Ptr("gmssteizlhjtclyeoo"),
		// 								Severity: to.Ptr("vvdajssdcypewdyechilxjmuijvdd"),
		// 								Source: to.Ptr("iy"),
		// 								Summary: to.Ptr("jtooblbvaxxrvcwgscbobq"),
		// 						}},
		// 						LastFailedEnableProtectionJob: &armrecoveryservicesdatareplication.ProtectedItemModelPropertiesLastFailedEnableProtectionJob{
		// 							Name: to.Ptr("hvvolptulpcxwbnjdzky"),
		// 							DisplayName: to.Ptr("zrqjbcozwiuypjjnvy"),
		// 							EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:55.456Z"); return t}()),
		// 							ID: to.Ptr("hjzgyxgdy"),
		// 							ScenarioName: to.Ptr("fhz"),
		// 							StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:55.456Z"); return t}()),
		// 							State: to.Ptr("ljsixxmmcaq"),
		// 						},
		// 						LastFailedPlannedFailoverJob: &armrecoveryservicesdatareplication.ProtectedItemModelPropertiesLastFailedPlannedFailoverJob{
		// 							Name: to.Ptr("ofblltxwhwzhyr"),
		// 							DisplayName: to.Ptr("whxsvbrzdhqsepbocfzsfx"),
		// 							EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:55.456Z"); return t}()),
		// 							ID: to.Ptr("ndjurplurnkguwfxx"),
		// 							ScenarioName: to.Ptr("ceksuyfiplxj"),
		// 							StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:55.456Z"); return t}()),
		// 							State: to.Ptr("wpur"),
		// 						},
		// 						LastSuccessfulPlannedFailoverTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:55.456Z"); return t}()),
		// 						LastSuccessfulTestFailoverTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:55.456Z"); return t}()),
		// 						LastSuccessfulUnplannedFailoverTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:55.456Z"); return t}()),
		// 						LastTestFailoverJob: &armrecoveryservicesdatareplication.ProtectedItemModelPropertiesLastTestFailoverJob{
		// 							Name: to.Ptr("cedjijdtnznsnigghrxnsaz"),
		// 							DisplayName: to.Ptr("lhkjfbonwdtxckwzfebfwdyu"),
		// 							EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:55.456Z"); return t}()),
		// 							ID: to.Ptr("cta"),
		// 							ScenarioName: to.Ptr("dfuovvz"),
		// 							StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:55.456Z"); return t}()),
		// 							State: to.Ptr("nhbzw"),
		// 						},
		// 						PolicyName: to.Ptr("tjoeiynplt"),
		// 						ProtectionState: to.Ptr(armrecoveryservicesdatareplication.ProtectionStateUnprotectedStatesBegin),
		// 						ProtectionStateDescription: to.Ptr("lp"),
		// 						ProvisioningState: to.Ptr(armrecoveryservicesdatareplication.ProvisioningStateSucceeded),
		// 						ReplicationExtensionName: to.Ptr("jwxdo"),
		// 						ReplicationHealth: to.Ptr(armrecoveryservicesdatareplication.HealthStatusNormal),
		// 						ResyncRequired: to.Ptr(true),
		// 						ResynchronizationState: to.Ptr(armrecoveryservicesdatareplication.ResynchronizationStateNone),
		// 						SourceFabricProviderID: to.Ptr("srggkxaruzlegtpdalscio"),
		// 						TargetDraID: to.Ptr("oscnhreunbyibimlpvsesu"),
		// 						TargetFabricID: to.Ptr("fb"),
		// 						TargetFabricProviderID: to.Ptr("sutiqezfbeiewwjezflvcitqj"),
		// 						TestFailoverState: to.Ptr(armrecoveryservicesdatareplication.TestFailoverStateNone),
		// 						TestFailoverStateDescription: to.Ptr("msjz"),
		// 					},
		// 					SystemData: &armrecoveryservicesdatareplication.ProtectedItemModelSystemData{
		// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:55.456Z"); return t}()),
		// 						CreatedBy: to.Ptr("ghut"),
		// 						CreatedByType: to.Ptr("tzczp"),
		// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:55.456Z"); return t}()),
		// 						LastModifiedBy: to.Ptr("epo"),
		// 						LastModifiedByType: to.Ptr("ekzmwexhjttb"),
		// 					},
		// 			}},
		// 		}
	}
}
