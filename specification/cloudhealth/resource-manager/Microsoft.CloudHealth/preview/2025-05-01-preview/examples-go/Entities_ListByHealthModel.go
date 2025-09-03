package armcloudhealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cloudhealth/armcloudhealth"
)

// Generated from example definition: 2025-05-01-preview/Entities_ListByHealthModel.json
func ExampleEntitiesClient_NewListByHealthModelPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcloudhealth.NewClientFactory("4980D7D5-4E07-47AD-AD34-E76C6BC9F061", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEntitiesClient().NewListByHealthModelPager("rgopenapi", "gPWT6GP85xRV248L7LhNRTD--2Yc73wu-5Qk-0tS", nil)
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
		// page = armcloudhealth.EntitiesClientListByHealthModelResponse{
		// 	EntityListResult: armcloudhealth.EntityListResult{
		// 		Value: []*armcloudhealth.Entity{
		// 			{
		// 				Properties: &armcloudhealth.EntityProperties{
		// 					ProvisioningState: to.Ptr(armcloudhealth.HealthModelProvisioningStateSucceeded),
		// 					DisplayName: to.Ptr("olsaxmichzpzqzpdworuxbvucaazeoxvnaujlvbpjijarbnfdmskksgdtyfdlfuftoecrgmmenvvgfkzlfcogkyhk"),
		// 					CanvasPosition: &armcloudhealth.EntityCoordinates{
		// 						X: to.Ptr[float32](14),
		// 						Y: to.Ptr[float32](13),
		// 					},
		// 					Icon: &armcloudhealth.IconDefinition{
		// 						IconName: to.Ptr("SystemComponent"),
		// 					},
		// 					HealthObjective: to.Ptr[float32](62),
		// 					Impact: to.Ptr(armcloudhealth.EntityImpactStandard),
		// 					Labels: map[string]*string{
		// 						"key1376": to.Ptr("ixfvzsfnpvkkbrcedaligqtopbadmojqgnzglztqytwickyeurumvqdqetmmikaqwuodokzjagoddnlolnputcdpyguuagucpbafkdpekxlxlmlrknzzjjxmbkysveyfmhzkgxverwcwdaolfqranhuearqchyxrtdlzabumuajmuxysgpaqwmwlrmqegyowtcnighwuudbgkzgzqfptsvdlvzgmnvnraufeocfkevwngzulwzjazttrwqcaakwwcehrdyhseimrwoqdkdxtcjadqcdrtdwwnieddexogctivtosotbddmjnjtexxrhngtnombrlqtasnncqmmmtnivnlkejghrvbfmullwnpjfhpejlzhcytgrrrybydizeefuxtuvjhlksxzyouukfhpsfpoqgcqoohvbqdvotzubjvqptqtapcahdhwhkxnyenlpsiepigjwukxwvpdgturhosxlfgidortepltfkkukrkuuaafhdjjwiozztomxjlluegvxrfsjiktoakpsjpqrxfsiajmzgexrfkergxbyoahyqhenydbvbtyqjwquruumwammhatjegpuokgwlildvghtlnpbsafhqqfltgxqxxusuybxamqqbhljfoaiaxkeqqsthzjdzisirutglrksotlabiitdzjudxtbbtetjxwizrvvzfcwuxiisalsbtvrpsfolasbegwivrkrjldijywozlldwxhedsnvvfecjuajhnbqmatkvsbuatpwdjfuhzdykrjprpvakktnwehehgbiewdxyxfyrruwjkndywcwwkyeupwjxceaqousxtufkhyqjnuluqkxxqupafnyrmhaxtnqzbpiavuwjagkdpqehfutmnjhmdruoatlvhkvvzaylaytvwuiirfurxeyegmiultvffswlzzhmicvxjozvngnyerepqskhjpaaicvwmlqhgbddjnudpppkeuuydjzzxhoxyefeszxpiwuexvyatxutnavzrmjmbnmalskaxnnrkcukdnkndoijtbirepqlcrmwygambodzwcppjpzerwyowpumxkypyrjzunzhutwjqpdbpwanunjnnnxmyqmlwzcmpvievaefcebzkooipxomlpviertmwoeekznzacypgptjuoegmikqwpnjlhwoqcgbggxtbneavfuixusnwdjdzxbbanoymoqnbldycwlzkeffsdbqxurzehzkmlowtoyqegkacazobopayenowenkpmyydloxnzkgjpvsirrklujljafltsvayxdvnfasmywffeifmrbqghaoprihgikeuuxyzoddhuqzsrladhpfexafdkdorlkfaukpailrlpyncyovlgdueygtzlsuykbuhbbcdfkotzqcelkrlxwuswzittcantblfpeikpwnppmnfzvsgarmqfpwmpqpfcicfnkxhepddwpqnvfqhedkecrpsutpfikwlmicfjcffommczdqvizrlofzqibjutnfczjecgfsyhhmylvdfhnsrovfivhgdmbbzwdqscjhvamvnpneryiafcykojgzvcmptphjbwsakbsmdrgvscrduqqgdztswobpctbnhdtjbrldpefdmdtemihpfbxllcsrydxbfwwolekwwxbyxexdlsepdzjpaxwpmqlsbrmyanzazdffepfwnltmopdqifsonrqcbkjphiydjhmcqfnkdbsrmvgknrvjywpmaehfshklwlmqajhkjlysxhqqeogtrbeqmopbfeqgrdwoihoebidkkurwygxhbhyivbubdtexrixwyqrsxwqjkdjmfdvxxdsfrnlumgwsywfzcnakabctdmcvqtyiijfwtvlkrmcvuncwcesmwvipmxoxeqaelzfgwznowvnwarfffevwdgmjvdkmkcrmtoxkboczjchbqnrdrodbetttulazfxwqxtrfcgjiusoubmqklesrqydiumytmqaqknvtxsvutmtxmbpccferzekdqzeqirnnavzmktdewfwihxwluqgtkynnqanexybirsssrfsjpzoseujluxtobzwuwqbllertseuwtifdihejvvtopoopdlogqvokhinvmryxlgjjcrtbgynxoztxpsyjgqxlgsvrvidpobbohjeniozridnjxbqewtgpjtkhdnqlceovpuxsrjcackwgtsllqazbxnhowajrtynuvldfcnedngnrwiwbjfogfzwlqkhbzklvziysuiyqbezkqpoanudvtabhlpgxljwailxvjarcumvxuwzxfevllhhvjyxqweesdgworgnneveovfgprphymmgjjoxjndsefmzrglkyootgjyarycquagpfkhiifqdmrwvwfyhtxtzhziefmysgdupawzaohqoecrebadanvnacsoeszhggciahbmpbsmbpjfzcmqcoquatvooeifsvdmfobivkzgvgbnadusjqcgvhqcxwrprtpulluwzutqivwhzncrpflgfikjiwubkndelhiprzzwrqunqwmpkrbhrcwutrwpeybrcplgqzxpohnasthxsdhjhxfqzntsiderxmcirhaoagswhtnjvjhtfbvujrxihbnqubzogedbhmnlmuylleruqpcpwhaevtkxuimmdvmqjhvdpkxkpsrwxbjfvoerlizufmcsjybvisohwhftdtslijozojfvrxswbxtxmksagrfupnrzuvepklqeoqtbksyhvavqrfmfioogifjlaacqnfmsjrnmhssaxnulrqaefikbkhsnfaelmiabbdhpsauikymiaynbxywybqgzhegxhrelpadodltzwgfirqertmoauuglcrpjxlznalzlqdisvtphfqefmgegxotsetvylexpbjlyxcznssqdbkshwocmq"),
		// 					},
		// 					HealthState: to.Ptr(armcloudhealth.HealthStateDegraded),
		// 					Signals: &armcloudhealth.SignalGroup{
		// 						AzureResource: &armcloudhealth.AzureResourceSignalGroup{
		// 							SignalAssignments: []*armcloudhealth.SignalAssignment{
		// 								{
		// 									SignalDefinitions: []*string{
		// 										to.Ptr("B3P1X3e-FZtZ-4Ak-2VLHGQ-4m4-05DE-XNW5zW3P-46XY-DC3SSX"),
		// 									},
		// 								},
		// 								{
		// 									SignalDefinitions: []*string{
		// 										to.Ptr("B3P1X3e-FZtZ-4Ak-2VLHGQ-4m4-05DE-XNW5zW3P-46XY-DC3SSX"),
		// 									},
		// 								},
		// 							},
		// 							AuthenticationSetting: to.Ptr("B3P1X3e-FZtZ-4Ak-2VLHGQ-4m4-05DE-XNW5zW3P-46XY-DC3SSX"),
		// 							AzureResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Compute/virtualMachines/myvm"),
		// 						},
		// 						AzureLogAnalytics: &armcloudhealth.LogAnalyticsSignalGroup{
		// 							SignalAssignments: []*armcloudhealth.SignalAssignment{
		// 								{
		// 									SignalDefinitions: []*string{
		// 										to.Ptr("B3P1X3e-FZtZ-4Ak-2VLHGQ-4m4-05DE-XNW5zW3P-46XY-DC3SSX"),
		// 									},
		// 								},
		// 								{
		// 									SignalDefinitions: []*string{
		// 										to.Ptr("B3P1X3e-FZtZ-4Ak-2VLHGQ-4m4-05DE-XNW5zW3P-46XY-DC3SSX"),
		// 									},
		// 								},
		// 							},
		// 							AuthenticationSetting: to.Ptr("B3P1X3e-FZtZ-4Ak-2VLHGQ-4m4-05DE-XNW5zW3P-46XY-DC3SSX"),
		// 							LogAnalyticsWorkspaceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.OperationalInsights/workspaces/myworkspace"),
		// 						},
		// 						AzureMonitorWorkspace: &armcloudhealth.AzureMonitorWorkspaceSignalGroup{
		// 							SignalAssignments: []*armcloudhealth.SignalAssignment{
		// 								{
		// 									SignalDefinitions: []*string{
		// 										to.Ptr("B3P1X3e-FZtZ-4Ak-2VLHGQ-4m4-05DE-XNW5zW3P-46XY-DC3SSX"),
		// 									},
		// 								},
		// 								{
		// 									SignalDefinitions: []*string{
		// 										to.Ptr("B3P1X3e-FZtZ-4Ak-2VLHGQ-4m4-05DE-XNW5zW3P-46XY-DC3SSX"),
		// 									},
		// 								},
		// 							},
		// 							AuthenticationSetting: to.Ptr("B3P1X3e-FZtZ-4Ak-2VLHGQ-4m4-05DE-XNW5zW3P-46XY-DC3SSX"),
		// 							AzureMonitorWorkspaceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.OperationalInsights/workspaces/myworkspace"),
		// 						},
		// 						Dependencies: &armcloudhealth.DependenciesSignalGroup{
		// 							AggregationType: to.Ptr(armcloudhealth.DependenciesAggregationTypeThresholds),
		// 							DegradedThreshold: to.Ptr("3"),
		// 							UnhealthyThreshold: to.Ptr("50%"),
		// 						},
		// 					},
		// 					DiscoveredBy: to.Ptr("discoveryRule1"),
		// 					DeletionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-18T14:04:14.531Z"); return t}()),
		// 					Alerts: &armcloudhealth.EntityAlerts{
		// 						Unhealthy: &armcloudhealth.AlertConfiguration{
		// 							Severity: to.Ptr(armcloudhealth.AlertSeveritySev1),
		// 							Description: to.Ptr("Alert description"),
		// 							ActionGroupIDs: []*string{
		// 								to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Insights/actionGroups/myactiongroup"),
		// 							},
		// 						},
		// 						Degraded: &armcloudhealth.AlertConfiguration{
		// 							Severity: to.Ptr(armcloudhealth.AlertSeveritySev4),
		// 							Description: to.Ptr("Alert description"),
		// 							ActionGroupIDs: []*string{
		// 								to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Insights/actionGroups/myactiongroup"),
		// 							},
		// 						},
		// 					},
		// 				},
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/Microsoft.CloudHealth/healthmodels/5D334Xv2hy4-Kj48w7b3JO0--51G625B6-m/entities/entity1"),
		// 				Name: to.Ptr("entity1"),
		// 				Type: to.Ptr("Microsoft.CloudHealth/healthmodels/entities"),
		// 				SystemData: &armcloudhealth.SystemData{
		// 					CreatedBy: to.Ptr("cbhzxxlvkmufetjjjwtk"),
		// 					CreatedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-18T14:04:09.327Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("arz"),
		// 					LastModifiedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-18T14:04:09.328Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/ai"),
		// 	},
		// }
	}
}
