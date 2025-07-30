package armrecoveryservicesdatareplication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
)

// Generated from example definition: 2024-09-01/ProtectedItem_Update.json
func ExampleProtectedItemClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("930CEC23-4430-4513-B855-DBA237E2F3BF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProtectedItemClient().BeginUpdate(ctx, "rgswagger_2024-09-01", "4", "d", armrecoveryservicesdatareplication.ProtectedItemModelUpdate{
		Properties: &armrecoveryservicesdatareplication.ProtectedItemModelPropertiesUpdate{
			CustomProperties: &armrecoveryservicesdatareplication.HyperVToAzStackHCIProtectedItemModelCustomPropertiesUpdate{
				InstanceType: to.Ptr("ProtectedItemModelCustomPropertiesUpdate"),
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
	// res = armrecoveryservicesdatareplication.ProtectedItemClientUpdateResponse{
	// 	ProtectedItemModel: &armrecoveryservicesdatareplication.ProtectedItemModel{
	// 		Properties: &armrecoveryservicesdatareplication.ProtectedItemModelProperties{
	// 			PolicyName: to.Ptr("tjoeiynplt"),
	// 			ReplicationExtensionName: to.Ptr("jwxdo"),
	// 			CustomProperties: &armrecoveryservicesdatareplication.HyperVToAzStackHCIProtectedItemModelCustomProperties{
	// 				InstanceType: to.Ptr("ProtectedItemModelCustomProperties"),
	// 			},
	// 			CorrelationID: to.Ptr("mvxvtcqwgp"),
	// 			ProvisioningState: to.Ptr(armrecoveryservicesdatareplication.ProvisioningStateCanceled),
	// 			ProtectionState: to.Ptr(armrecoveryservicesdatareplication.ProtectionStateUnprotectedStatesBegin),
	// 			ProtectionStateDescription: to.Ptr("lp"),
	// 			TestFailoverState: to.Ptr(armrecoveryservicesdatareplication.TestFailoverStateNone),
	// 			TestFailoverStateDescription: to.Ptr("msjz"),
	// 			ResynchronizationState: to.Ptr(armrecoveryservicesdatareplication.ResynchronizationStateNone),
	// 			FabricObjectID: to.Ptr("kjcizdpahzqsrwyiywbhyzdxsufj"),
	// 			FabricObjectName: to.Ptr("glrjwtvmejxuagjepcwaxhih"),
	// 			SourceFabricProviderID: to.Ptr("srggkxaruzlegtpdalscio"),
	// 			TargetFabricProviderID: to.Ptr("sutiqezfbeiewwjezflvcitqj"),
	// 			FabricID: to.Ptr("ebsxoblmhlhqjzzjzdwo"),
	// 			TargetFabricID: to.Ptr("fb"),
	// 			FabricAgentID: to.Ptr("c"),
	// 			TargetFabricAgentID: to.Ptr("tsekhcwafyaoujmijzawvexruifs"),
	// 			ResyncRequired: to.Ptr(true),
	// 			LastSuccessfulPlannedFailoverTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:55.456Z"); return t}()),
	// 			LastSuccessfulUnplannedFailoverTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:55.456Z"); return t}()),
	// 			LastSuccessfulTestFailoverTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:55.456Z"); return t}()),
	// 			CurrentJob: &armrecoveryservicesdatareplication.ProtectedItemJobProperties{
	// 				ScenarioName: to.Ptr("ljbnhbdmreowdqnlcqycvaramwuii"),
	// 				ID: to.Ptr("bnmbzxzyfgwh"),
	// 				Name: to.Ptr("kqtvbrfmqaxdgpttkbmzpwafjp"),
	// 				DisplayName: to.Ptr("awutlqrisstqb"),
	// 				State: to.Ptr("ztlpngveoqcdejpwaiudhrioskauqv"),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:55.456Z"); return t}()),
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:55.456Z"); return t}()),
	// 			},
	// 			AllowedJobs: []*string{
	// 				to.Ptr("guryeoocjbvqvalfkrxecpxwynpxs"),
	// 			},
	// 			LastFailedEnableProtectionJob: &armrecoveryservicesdatareplication.ProtectedItemJobProperties{
	// 				ScenarioName: to.Ptr("fhz"),
	// 				ID: to.Ptr("hjzgyxgdy"),
	// 				Name: to.Ptr("hvvolptulpcxwbnjdzky"),
	// 				DisplayName: to.Ptr("zrqjbcozwiuypjjnvy"),
	// 				State: to.Ptr("ljsixxmmcaq"),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:55.456Z"); return t}()),
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:55.456Z"); return t}()),
	// 			},
	// 			LastFailedPlannedFailoverJob: &armrecoveryservicesdatareplication.ProtectedItemJobProperties{
	// 				ScenarioName: to.Ptr("ceksuyfiplxj"),
	// 				ID: to.Ptr("ndjurplurnkguwfxx"),
	// 				Name: to.Ptr("ofblltxwhwzhyr"),
	// 				DisplayName: to.Ptr("whxsvbrzdhqsepbocfzsfx"),
	// 				State: to.Ptr("wpur"),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:55.456Z"); return t}()),
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:55.456Z"); return t}()),
	// 			},
	// 			LastTestFailoverJob: &armrecoveryservicesdatareplication.ProtectedItemJobProperties{
	// 				ScenarioName: to.Ptr("dfuovvz"),
	// 				ID: to.Ptr("cta"),
	// 				Name: to.Ptr("cedjijdtnznsnigghrxnsaz"),
	// 				DisplayName: to.Ptr("lhkjfbonwdtxckwzfebfwdyu"),
	// 				State: to.Ptr("nhbzw"),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:55.456Z"); return t}()),
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:55.456Z"); return t}()),
	// 			},
	// 			ReplicationHealth: to.Ptr(armrecoveryservicesdatareplication.HealthStatusNormal),
	// 			HealthErrors: []*armrecoveryservicesdatareplication.HealthErrorModel{
	// 				{
	// 					AffectedResourceType: to.Ptr("scfniv"),
	// 					AffectedResourceCorrelationIDs: []*string{
	// 						to.Ptr("fope"),
	// 					},
	// 					ChildErrors: []*armrecoveryservicesdatareplication.InnerHealthErrorModel{
	// 						{
	// 							Code: to.Ptr("yuxxpblihirpedwkigywgwjjrlzq"),
	// 							HealthCategory: to.Ptr("mhdgfjqwbikhxmhtomkl"),
	// 							Category: to.Ptr("lcsdxrqxquke"),
	// 							Severity: to.Ptr("wqxxiuaqjyagq"),
	// 							Source: to.Ptr("wevvftugwydzzw"),
	// 							CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:52.128Z"); return t}()),
	// 							IsCustomerResolvable: to.Ptr(true),
	// 							Summary: to.Ptr("djsmgrltruljo"),
	// 							Message: to.Ptr("sskcei"),
	// 							Causes: to.Ptr("kefaugkpxjkpulimjthjnl"),
	// 							Recommendation: to.Ptr("kqybwaesqumywtjepi"),
	// 						},
	// 					},
	// 					Code: to.Ptr("dgxkefzmeukd"),
	// 					HealthCategory: to.Ptr("itc"),
	// 					Category: to.Ptr("leigw"),
	// 					Severity: to.Ptr("vvdajssdcypewdyechilxjmuijvdd"),
	// 					Source: to.Ptr("iy"),
	// 					CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:52.128Z"); return t}()),
	// 					IsCustomerResolvable: to.Ptr(true),
	// 					Summary: to.Ptr("jtooblbvaxxrvcwgscbobq"),
	// 					Message: to.Ptr("lbywtdprdqdekl"),
	// 					Causes: to.Ptr("xznphqrrmsdzm"),
	// 					Recommendation: to.Ptr("gmssteizlhjtclyeoo"),
	// 				},
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataReplication/replicationVaults/vault1/protectedItems/item1"),
	// 		Name: to.Ptr("t"),
	// 		Type: to.Ptr("xlyjashandpfwivuipoplgkgsnwoh"),
	// 		SystemData: &armrecoveryservicesdatareplication.SystemData{
	// 			CreatedBy: to.Ptr("ewufpudzcjrljhmmzhfnxoqdqwnya"),
	// 			CreatedByType: to.Ptr(armrecoveryservicesdatareplication.CreatedByType("zioqm")),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:53.022Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("rx"),
	// 			LastModifiedByType: to.Ptr(armrecoveryservicesdatareplication.CreatedByType("tqbvuqoakaaqij")),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:53.022Z"); return t}()),
	// 		},
	// 	},
	// }
}
