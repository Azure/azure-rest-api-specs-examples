package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase/v2"
)

// Generated from example definition: 2025-09-01/AutonomousDatabases_ListByResourceGroup_MaximumSet_Gen.json
func ExampleAutonomousDatabasesClient_NewListByResourceGroupPager_listAutonomousDatabaseByResourceGroupGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAutonomousDatabasesClient().NewListByResourceGroupPager("rgopenapi", nil)
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
		// page = armoracledatabase.AutonomousDatabasesClientListByResourceGroupResponse{
		// 	AutonomousDatabaseListResult: armoracledatabase.AutonomousDatabaseListResult{
		// 		Value: []*armoracledatabase.AutonomousDatabase{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/autonomousDatabases/databasedb1"),
		// 				Type: to.Ptr("Oracle.Database/autonomousDatabases"),
		// 				Location: to.Ptr("eastus"),
		// 				Tags: map[string]*string{
		// 				},
		// 				Properties: &armoracledatabase.AutonomousDatabaseProperties{
		// 					AutonomousDatabaseID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/autonomousDatabases/databasedb1"),
		// 					ActualUsedDataStorageSizeInTbs: to.Ptr[float64](3),
		// 					AutonomousMaintenanceScheduleType: to.Ptr(armoracledatabase.AutonomousMaintenanceScheduleTypeRegular),
		// 					CharacterSet: to.Ptr("AL32UTF8"),
		// 					NcharacterSet: to.Ptr("AL16UTF16"),
		// 					ComputeCount: to.Ptr[float32](2),
		// 					ComputeModel: to.Ptr(armoracledatabase.ComputeModelECPU),
		// 					CPUCoreCount: to.Ptr[int32](1),
		// 					CustomerContacts: []*armoracledatabase.CustomerContact{
		// 						{
		// 							Email: to.Ptr("agyiqecugrloatgwpvmilmvutcnyjpxzhbilhhqfvqqblfgursqelzjjnwnmpfstitmcgkovzxnstiqqwjnhwwaufbnkebpqxlqwmfnmtlkgkoxcnjwcnfqbdtokhjalagxphkuiwxtxrzuipokiuczmuwoqoebkjvhytlhtxzshwsdoywluoggznuyuozqibiwdgwqbgnyogysdjpvlowmvuvq"),
		// 						},
		// 					},
		// 					DataStorageSizeInGbs: to.Ptr[int32](1024),
		// 					DataStorageSizeInTbs: to.Ptr[int32](1),
		// 					DatabaseEdition: to.Ptr(armoracledatabase.DatabaseEditionTypeEnterpriseEdition),
		// 					DataBaseType: to.Ptr(armoracledatabase.DataBaseTypeRegular),
		// 					DbVersion: to.Ptr("18.4.0.0"),
		// 					DisplayName: to.Ptr("example_autonomous_databasedb1"),
		// 					IsAutoScalingEnabled: to.Ptr(true),
		// 					IsAutoScalingForStorageEnabled: to.Ptr(true),
		// 					IsLocalDataGuardEnabled: to.Ptr(true),
		// 					IsMtlsConnectionRequired: to.Ptr(true),
		// 					LicenseModel: to.Ptr(armoracledatabase.LicenseModelBringYourOwnLicense),
		// 					LifecycleState: to.Ptr(armoracledatabase.AutonomousDatabaseLifecycleState("Succeeded")),
		// 					LifecycleDetails: to.Ptr("success"),
		// 					PrivateEndpointIP: to.Ptr("nnoxtfettnhwhhjdtylhkhwj"),
		// 					PrivateEndpointLabel: to.Ptr("j"),
		// 					SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
		// 					VnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1"),
		// 					ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
		// 					OciURL: to.Ptr("https://fake"),
		// 					TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-09T20:44:09.466Z"); return t}()),
		// 					TimeMaintenanceBegin: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T04:32:58.715Z"); return t}()),
		// 					TimeMaintenanceEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T04:32:58.715Z"); return t}()),
		// 					UsedDataStorageSizeInGbs: to.Ptr[int32](12),
		// 					UsedDataStorageSizeInTbs: to.Ptr[int32](15),
		// 					Ocid: to.Ptr("ocid1..aaaaa"),
		// 					WhitelistedIPs: []*string{
		// 						to.Ptr("1.1.1.1"),
		// 						to.Ptr("1.1.1.0/24"),
		// 						to.Ptr("1.1.2.25"),
		// 					},
		// 					DbWorkload: to.Ptr(armoracledatabase.WorkloadTypeOLTP),
		// 					PeerDbIDs: []*string{
		// 						to.Ptr("jflsgwiukjeriuohebig"),
		// 					},
		// 					IsRemoteDataGuardEnabled: to.Ptr(true),
		// 					LocalDisasterRecoveryType: to.Ptr(armoracledatabase.DisasterRecoveryTypeAdg),
		// 					TimeDisasterRecoveryRoleChanged: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T04:32:58.712Z"); return t}()),
		// 					RemoteDisasterRecoveryConfiguration: &armoracledatabase.DisasterRecoveryConfigurationDetails{
		// 						DisasterRecoveryType: to.Ptr(armoracledatabase.DisasterRecoveryTypeAdg),
		// 						TimeSnapshotStandbyEnabledTill: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T04:32:58.715Z"); return t}()),
		// 						IsSnapshotStandby: to.Ptr(true),
		// 						IsReplicateAutomaticBackups: to.Ptr(true),
		// 					},
		// 					LocalStandbyDb: &armoracledatabase.AutonomousDatabaseStandbySummary{
		// 						LagTimeInSeconds: to.Ptr[int32](13),
		// 						LifecycleState: to.Ptr(armoracledatabase.AutonomousDatabaseLifecycleStateProvisioning),
		// 						LifecycleDetails: to.Ptr("zoiyaaibuuhm"),
		// 						TimeDataGuardRoleChanged: to.Ptr("inggk"),
		// 						TimeDisasterRecoveryRoleChanged: to.Ptr("q"),
		// 					},
		// 					FailedDataRecoveryInSeconds: to.Ptr[int32](11),
		// 					ScheduledOperationsList: []*armoracledatabase.ScheduledOperationsType{
		// 						{
		// 							DayOfWeek: &armoracledatabase.DayOfWeek{
		// 								Name: to.Ptr(armoracledatabase.DayOfWeekNameMonday),
		// 							},
		// 							ScheduledStartTime: to.Ptr("lwwvkazgmfremfwhckfb"),
		// 							ScheduledStopTime: to.Ptr("hjwagzxijpiaogulmnmbuqakpqxhkjvaypjqnvbvtjddc"),
		// 						},
		// 					},
		// 					AllocatedStorageSizeInTbs: to.Ptr[float64](1),
		// 					ApexDetails: &armoracledatabase.ApexDetailsType{
		// 						ApexVersion: to.Ptr("scvpjwygbzqzevlztyfvqiaom"),
		// 						OrdsVersion: to.Ptr("djlwvhpipimxaguklshpppjrzasbk"),
		// 					},
		// 					AvailableUpgradeVersions: []*string{
		// 						to.Ptr("yegwifyowsnpndi"),
		// 					},
		// 					ConnectionStrings: &armoracledatabase.ConnectionStringType{
		// 						AllConnectionStrings: &armoracledatabase.AllConnectionStringType{
		// 							High: to.Ptr("exdinleextbcjinutlkvnqyxhvandtihncykjzrhfdhfrxdarny"),
		// 							Low: to.Ptr("vurudxqtummqqbnidwthmsqgujufjxwfnejdpuxkgyoxlgqhcgsfjcpzaqeioslpehjfashipdsjhkrpdobstvgxsqrgvcrpbiiabhvymdsylqsjedrimqhtttmszlaqyukopuufbtkbtwgdydrvnvkcdqmphwzpcjxlgefzrdajyczzjdpuzvhqvupbvrpvqhzoaalg"),
		// 							Medium: to.Ptr("ishtubsuzgwtkfdqfgyxjlhehiokdvjfhwqhvmgtuksboshulroytcnubtrxxjbgoutftpzeavbldsoqjwmamgfwevhppyyeckythirzvaqujrjaiqnpyvycakhwgtuftmuxavdgdvbqxgsdqwbnqzmrzymwiydhxekenbweaghgvyveuysxqmovwodzwwfrxhtlcegekjk"),
		// 						},
		// 						Dedicated: to.Ptr("okoggzbidoaknwikuqpvepxvvcseukouprpfrldmakztkleeizbjf"),
		// 						High: to.Ptr("pggylyrivfn"),
		// 						Low: to.Ptr("zrjsbtdbfluaipbzgcvvhyuvqoczjneihaiftkfmuvvyoldslgvvpwthieyrcoxvh"),
		// 						Medium: to.Ptr("ebjnwenxvyeinsabrppychqbcawfxgplfacbsizltwfhpdafbkawopppqsxemlnmrbiqlstjupgdmpfcyyxgofmitbdiarrpprhntntqqjklseigycfcpmmlqiznxzliserjppmgfjatnmtbdxqtlbmrmpfbpoxmyffkkoptpayigeeefmqczroouqjxchswffywpsmyqohbyaclhsrwgqyuuyynvxyyzkche"),
		// 						Profiles: []*armoracledatabase.ProfileType{
		// 							{
		// 								ConsumerGroup: to.Ptr(armoracledatabase.ConsumerGroupHigh),
		// 								DisplayName: to.Ptr("mqqdgidxuovxhcwrkanybxzplautekarsxbcbzlkikpmmvjrdrrkncbamdtcuksplamigrdkydjbzeurbmjgehgppovxqhuzasduwptrlyaurzszzqpztckhpdniepaglzeublbwffxebfespqyfpljlutregvlzzjo"),
		// 								HostFormat: to.Ptr(armoracledatabase.HostFormatTypeFqdn),
		// 								IsRegional: to.Ptr(true),
		// 								Protocol: to.Ptr(armoracledatabase.ProtocolTypeTCP),
		// 								SessionMode: to.Ptr(armoracledatabase.SessionModeTypeDirect),
		// 								SyntaxFormat: to.Ptr(armoracledatabase.SyntaxFormatTypeLong),
		// 								TLSAuthentication: to.Ptr(armoracledatabase.TLSAuthenticationTypeServer),
		// 								Value: to.Ptr("bdrnwqpzbbzdipqqhnroxiuewqg"),
		// 							},
		// 						},
		// 					},
		// 					ConnectionUrls: &armoracledatabase.ConnectionURLType{
		// 						ApexURL: to.Ptr("epnebmudvzijxrfgabsdjewqfotqjmnxvokfhlyklhvtrjpprnqujthmceuhpfuumcbfxktppfguqduzkukxqkofoyyycljjtruyjtoiesxlrwwzonozaxetzrkpmzwasyvryvkryawxxf"),
		// 						DatabaseTransformsURL: to.Ptr("hujiemysucgdgtasazsdtwnxmtjppugunrqnzfzneatukuyzvkfseusjaxrourznsrwxjbvzfansdcyfxnvcyghl"),
		// 						GraphStudioURL: to.Ptr("bucnwmwixwemqqtoozfclfzqenskkyssvcatwbptsezpzdwgnaexgxutsvaibnkawyohqklnktzlhdbhbstm"),
		// 						MachineLearningNotebookURL: to.Ptr("vfhnqsrabxcrjnpqaqkgnpwhxffsqkrgcijdkkvnaoangzkcbgwklufujhmlgydxueybugxzgokxbbappdslttpdthhbmxrgcicqzyjyahjeiqopuglgbjfbhufuvsogquelagbjtyotwhmecwupooitcaftldxjycgfnlilrnicqjxnsucieftadjbvptzltmgqkxhttfkkbutaxvtfzbvbbxbmpxeeyfethpofnmbbqbtlqvnfgelvtjizckgixpptkilcvrntknusvppgnobokjpepynndswcqsnewhfnlxgmownfwfnokhbqulzyuessvxxtcdcnmumbbpjchmjbvjecbbinjolmuoaixzunawlxnoqbpzkczdsubpqpdltnfydwevearrdirzaszsudcxaspozeop"),
		// 						MongoDbURL: to.Ptr("dzmsqtcgsrdgwjlnrfmzcqcrkdqwmjrccxsszwdgpcygywnuurklwthgonxcnwaqcgzoexnaanwzsqwemcijuzxqbrkpvydizjraicgnspizwwnwureyey"),
		// 						OrdsURL: to.Ptr("lmqdgziantbczaneiqxopnaexcroelkbcgggjipzqfhoduwqodoyeghzjyuyhesewopbujxnoiziidhslxdawrfayjvxzjwfobtjrepldlmwhauiurzhbpyxsbueugddmdfindxsdjddqamwbptozzmobugnpezxyxdopripljdwnogjgrjwjfbugdkqrjwzjvavsdmgvspnzcgcjzxauqxqljmfgrtyohfppbmprmexiirvlmymxzyciaraihbwihvahyaciv"),
		// 						SQLDevWebURL: to.Ptr("pktklsvlakmblcakipxy"),
		// 					},
		// 					DataSafeStatus: to.Ptr(armoracledatabase.DataSafeStatusTypeRegistering),
		// 					InMemoryAreaInGbs: to.Ptr[int32](23),
		// 					NextLongTermBackupTimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T04:32:58.715Z"); return t}()),
		// 					LongTermBackupSchedule: &armoracledatabase.LongTermBackUpScheduleDetails{
		// 						RepeatCadence: to.Ptr(armoracledatabase.RepeatCadenceTypeOneTime),
		// 						TimeOfBackup: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T04:32:58.715Z"); return t}()),
		// 						RetentionPeriodInDays: to.Ptr[int32](188),
		// 						IsDisabled: to.Ptr(true),
		// 					},
		// 					IsPreview: to.Ptr(true),
		// 					LocalAdgAutoFailoverMaxDataLossLimit: to.Ptr[int32](2449),
		// 					MemoryPerOracleComputeUnitInGbs: to.Ptr[int32](11),
		// 					OpenMode: to.Ptr(armoracledatabase.OpenModeTypeReadOnly),
		// 					OperationsInsightsStatus: to.Ptr(armoracledatabase.OperationsInsightsStatusTypeEnabling),
		// 					PermissionLevel: to.Ptr(armoracledatabase.PermissionLevelTypeRestricted),
		// 					PrivateEndpoint: to.Ptr("rlhggatujqzocvibe"),
		// 					ProvisionableCpus: []*int32{
		// 						to.Ptr[int32](21),
		// 					},
		// 					Role: to.Ptr(armoracledatabase.RoleTypePrimary),
		// 					ServiceConsoleURL: to.Ptr("threiehxxauchaokxangmgukvgbefpwaoneyhdhkmizdccxkkyedotccorydefqkbqhpdifguchutkzvrbmqnzdjpnnkgjfnwxgxbvvhzkfqmavkrkvixvcjinvrlyvebcnljyerhwgisubnoldediofuctetnrbf"),
		// 					SQLWebDeveloperURL: to.Ptr("xmpfnyypagnejpkyklestynmjsdqqwcvfjkmguyplihtznhosjvlbcxgeeonyzkovfrprcdqurmvydhkvbqhrnmiroimtcbybfkkicnettcfrgmgkgbjxmwfnylsevpkdyjrvwjldufrdtmjwiksmmvuzlyxiuatcaseobsicuodznlrvmspfwypjuifgzpofynrbzrlscpqbsnrsnopdnbmupbsfzthjsgzxqnbmvyswxwnfckitncvurypweezyzgoxtntrwpejwjabrybrknoqabeywjotpkqrfghqtatmnxbsastycrhjzevpcjynumkguifkoxdawodjgjsfydhmvkagjeumihyhzcbqcvtcwymbhrvdhpumcxlmdantnthjkjcfjkmszsolqtozlguvvnprkjfmnmqyxkuzzkbngtlusasqyzdszamzczzjnbtqqnqtdrjraekwvzcuzouxdhvnlgdvtsfztyabrbrsjdwxdksgmfoeeoqgazdw"),
		// 					SupportedRegionsToCloneTo: []*string{
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 						to.Ptr("bgllibqtbp"),
		// 					},
		// 					TimeDataGuardRoleChanged: to.Ptr("wkompilkf"),
		// 					TimeDeletionOfFreeAutonomousDatabase: to.Ptr("ybhhebydubpqmcmfxip"),
		// 					TimeLocalDataGuardEnabled: to.Ptr("xrevlcxoiocqbvoiguqg"),
		// 					TimeOfLastFailover: to.Ptr("xgvir"),
		// 					TimeOfLastRefresh: to.Ptr("jpszxyfjatlkawkovzcrykgcfkn"),
		// 					TimeOfLastRefreshPoint: to.Ptr("nlukypraetzzkzsbfxbfwddqzl"),
		// 					TimeOfLastSwitchover: to.Ptr("fufkopddvnvcqmplnaffcke"),
		// 					TimeReclamationOfFreeAutonomousDatabase: to.Ptr("kyltyw"),
		// 					BackupRetentionPeriodInDays: to.Ptr[int32](24),
		// 				},
		// 				Name: to.Ptr("xyrqhmcykunr"),
		// 				SystemData: &armoracledatabase.SystemData{
		// 					CreatedBy: to.Ptr("sqehacivpuim"),
		// 					CreatedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T04:32:58.716Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("axrqfdkqylvjv"),
		// 					LastModifiedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T04:32:58.716Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
