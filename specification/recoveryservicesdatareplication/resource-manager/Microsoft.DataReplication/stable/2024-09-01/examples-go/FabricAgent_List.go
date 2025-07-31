package armrecoveryservicesdatareplication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
)

// Generated from example definition: 2024-09-01/FabricAgent_List.json
func ExampleFabricAgentClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("930CEC23-4430-4513-B855-DBA237E2F3BF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFabricAgentClient().NewListPager("rgswagger_2024-09-01", "wPR", nil)
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
		// page = armrecoveryservicesdatareplication.FabricAgentClientListResponse{
		// 	FabricAgentModelListResult: armrecoveryservicesdatareplication.FabricAgentModelListResult{
		// 		Value: []*armrecoveryservicesdatareplication.FabricAgentModel{
		// 			{
		// 				Properties: &armrecoveryservicesdatareplication.FabricAgentModelProperties{
		// 					ProvisioningState: to.Ptr(armrecoveryservicesdatareplication.ProvisioningStateCanceled),
		// 					HealthErrors: []*armrecoveryservicesdatareplication.HealthErrorModel{
		// 						{
		// 							AffectedResourceType: to.Ptr("scfniv"),
		// 							AffectedResourceCorrelationIDs: []*string{
		// 								to.Ptr("fope"),
		// 							},
		// 							ChildErrors: []*armrecoveryservicesdatareplication.InnerHealthErrorModel{
		// 								{
		// 									Code: to.Ptr("yuxxpblihirpedwkigywgwjjrlzq"),
		// 									HealthCategory: to.Ptr("mhdgfjqwbikhxmhtomkl"),
		// 									Category: to.Ptr("lcsdxrqxquke"),
		// 									Severity: to.Ptr("wqxxiuaqjyagq"),
		// 									Source: to.Ptr("wevvftugwydzzw"),
		// 									CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:52.128Z"); return t}()),
		// 									IsCustomerResolvable: to.Ptr(true),
		// 									Summary: to.Ptr("djsmgrltruljo"),
		// 									Message: to.Ptr("sskcei"),
		// 									Causes: to.Ptr("kefaugkpxjkpulimjthjnl"),
		// 									Recommendation: to.Ptr("kqybwaesqumywtjepi"),
		// 								},
		// 							},
		// 							Code: to.Ptr("dgxkefzmeukd"),
		// 							HealthCategory: to.Ptr("itc"),
		// 							Category: to.Ptr("leigw"),
		// 							Severity: to.Ptr("vvdajssdcypewdyechilxjmuijvdd"),
		// 							Source: to.Ptr("iy"),
		// 							CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:52.128Z"); return t}()),
		// 							IsCustomerResolvable: to.Ptr(true),
		// 							Summary: to.Ptr("jtooblbvaxxrvcwgscbobq"),
		// 							Message: to.Ptr("lbywtdprdqdekl"),
		// 							Causes: to.Ptr("xznphqrrmsdzm"),
		// 							Recommendation: to.Ptr("gmssteizlhjtclyeoo"),
		// 						},
		// 					},
		// 					MachineID: to.Ptr("envzcoijbqhtrpncbjbhk"),
		// 					MachineName: to.Ptr("y"),
		// 					AuthenticationIdentity: &armrecoveryservicesdatareplication.IdentityModel{
		// 						TenantID: to.Ptr("joclkkdovixwapephhxaqtefubhhmq"),
		// 						ApplicationID: to.Ptr("cwktzrwajuvfyyymfstpey"),
		// 						ObjectID: to.Ptr("khsiaqfbpuhp"),
		// 						Audience: to.Ptr("dkjobanyqgzenivyxhvavottpc"),
		// 						AADAuthority: to.Ptr("bubwwbowfhdmujrt"),
		// 					},
		// 					ResourceAccessIdentity: &armrecoveryservicesdatareplication.IdentityModel{
		// 						TenantID: to.Ptr("joclkkdovixwapephhxaqtefubhhmq"),
		// 						ApplicationID: to.Ptr("cwktzrwajuvfyyymfstpey"),
		// 						ObjectID: to.Ptr("khsiaqfbpuhp"),
		// 						Audience: to.Ptr("dkjobanyqgzenivyxhvavottpc"),
		// 						AADAuthority: to.Ptr("bubwwbowfhdmujrt"),
		// 					},
		// 					CustomProperties: &armrecoveryservicesdatareplication.VMwareFabricAgentModelCustomProperties{
		// 						InstanceType: to.Ptr("FabricAgentModelCustomProperties"),
		// 					},
		// 				},
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataReplication/replicationFabrics/fabric1/fabricAgents/agent1"),
		// 				Name: to.Ptr("rhojydcwjgvgexpdwswjib"),
		// 				Type: to.Ptr("toipsryjyqchikyakeiuatshiu"),
		// 				SystemData: &armrecoveryservicesdatareplication.SystemData{
		// 					CreatedBy: to.Ptr("yhdmbqrsgimuucexvpas"),
		// 					CreatedByType: to.Ptr(armrecoveryservicesdatareplication.CreatedByType("brnojz")),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:53.716Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("balecqnwu"),
		// 					LastModifiedByType: to.Ptr(armrecoveryservicesdatareplication.CreatedByType("ukvqlncihf")),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:53.716Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
