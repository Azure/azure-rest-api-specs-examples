package armoracledatabase_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase/v2"
)

// Generated from example definition: 2025-09-01/AutonomousDatabases_Update_MaximumSet_Gen.json
func ExampleAutonomousDatabasesClient_BeginUpdate_patchAutonomousDatabaseGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAutonomousDatabasesClient().BeginUpdate(ctx, "rgopenapi", "databasedb1", armoracledatabase.AutonomousDatabaseUpdate{
		Tags: map[string]*string{
			"key9827": to.Ptr("bygpoqozrwfyiootncgcqq"),
		},
		Properties: &armoracledatabase.AutonomousDatabaseUpdateProperties{
			AdminPassword:                     to.Ptr("<a-password-goes-here>"),
			AutonomousMaintenanceScheduleType: to.Ptr(armoracledatabase.AutonomousMaintenanceScheduleTypeEarly),
			ComputeCount:                      to.Ptr[float32](56.1),
			CPUCoreCount:                      to.Ptr[int32](45),
			CustomerContacts: []*armoracledatabase.CustomerContact{
				{
					Email: to.Ptr("dummyemail@microsoft.com"),
				},
			},
			DataStorageSizeInTbs:           to.Ptr[int32](133),
			DataStorageSizeInGbs:           to.Ptr[int32](175271),
			DisplayName:                    to.Ptr("lrdrjpyyvufnxdzpwvlkmfukpstrjftdxcejcxtnqhxqbhvtzeiokllnspotsqeggddxkjjtf"),
			IsAutoScalingEnabled:           to.Ptr(true),
			IsAutoScalingForStorageEnabled: to.Ptr(true),
			PeerDbID:                       to.Ptr("qmpfwtvpfvbgmulethqznsyyjlpxmyfqfanrymzqsgraavtmlqqbexpzguyqybngoupbshlzpxv"),
			IsLocalDataGuardEnabled:        to.Ptr(true),
			IsMtlsConnectionRequired:       to.Ptr(true),
			LicenseModel:                   to.Ptr(armoracledatabase.LicenseModelLicenseIncluded),
			ScheduledOperationsList: []*armoracledatabase.ScheduledOperationsTypeUpdate{
				{
					DayOfWeek: &armoracledatabase.DayOfWeekUpdate{
						Name: to.Ptr(armoracledatabase.DayOfWeekNameMonday),
					},
					ScheduledStartTime: to.Ptr("lwwvkazgmfremfwhckfb"),
					ScheduledStopTime:  to.Ptr("hjwagzxijpiaogulmnmbuqakpqxhkjvaypjqnvbvtjddc"),
				},
			},
			DatabaseEdition: to.Ptr(armoracledatabase.DatabaseEditionTypeStandardEdition),
			LongTermBackupSchedule: &armoracledatabase.LongTermBackUpScheduleDetails{
				RepeatCadence:         to.Ptr(armoracledatabase.RepeatCadenceTypeOneTime),
				TimeOfBackup:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T04:32:58.715Z"); return t }()),
				RetentionPeriodInDays: to.Ptr[int32](188),
				IsDisabled:            to.Ptr(true),
			},
			LocalAdgAutoFailoverMaxDataLossLimit: to.Ptr[int32](212),
			OpenMode:                             to.Ptr(armoracledatabase.OpenModeTypeReadOnly),
			PermissionLevel:                      to.Ptr(armoracledatabase.PermissionLevelTypeRestricted),
			Role:                                 to.Ptr(armoracledatabase.RoleTypePrimary),
			BackupRetentionPeriodInDays:          to.Ptr[int32](12),
			WhitelistedIPs: []*string{
				to.Ptr("kfierlppwurtqrhfxwgfgrnqtmvraignzwsddwmpdijeveuevuoejfmbjvpnlrmmdflilxcwkkzvdofctsdjfxrrrwctihhnchtrouauesqbmlcqhzwnppnhrtitecenlfyshassvajukbwxudhlwixkvkgsessvshcwmleoqujeemwenhwlsccbcjnnviugzgylsxkssalqoicatcvkahogdvweymhdxboyqwhaxuzlmrdbvgbnnetobkbwygcsflzanwknlybvvzgjzjirpfrksbxwgllgfxwdflcisvxpkjecpgdaxccqkzxofedkrawvhzeabmunpykwd"),
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
	// res = armoracledatabase.AutonomousDatabasesClientUpdateResponse{
	// 	AutonomousDatabase: &armoracledatabase.AutonomousDatabase{
	// 		Properties: &armoracledatabase.AutonomousDatabaseProperties{
	// 			DataBaseType: to.Ptr(armoracledatabase.DataBaseTypeRegular),
	// 			DisplayName: to.Ptr("example_autonomous_databasedb1"),
	// 			ComputeModel: to.Ptr(armoracledatabase.ComputeModelECPU),
	// 			ComputeCount: to.Ptr[float32](2),
	// 			DataStorageSizeInTbs: to.Ptr[int32](1),
	// 			DbVersion: to.Ptr("18.4.0.0"),
	// 			CharacterSet: to.Ptr("AL32UTF8"),
	// 			NcharacterSet: to.Ptr("AL16UTF16"),
	// 			SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
	// 			VnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1"),
	// 			AutonomousMaintenanceScheduleType: to.Ptr(armoracledatabase.AutonomousMaintenanceScheduleTypeRegular),
	// 			CPUCoreCount: to.Ptr[int32](1),
	// 			CustomerContacts: []*armoracledatabase.CustomerContact{
	// 				{
	// 					Email: to.Ptr("dummyemail@microsoft.com"),
	// 				},
	// 			},
	// 			DataStorageSizeInGbs: to.Ptr[int32](1024),
	// 			DbWorkload: to.Ptr(armoracledatabase.WorkloadTypeOLTP),
	// 			IsAutoScalingEnabled: to.Ptr(true),
	// 			IsAutoScalingForStorageEnabled: to.Ptr(true),
	// 			PeerDbIDs: []*string{
	// 				to.Ptr("gpubz"),
	// 			},
	// 			IsLocalDataGuardEnabled: to.Ptr(true),
	// 			IsRemoteDataGuardEnabled: to.Ptr(true),
	// 			LocalDisasterRecoveryType: to.Ptr(armoracledatabase.DisasterRecoveryTypeAdg),
	// 			TimeDisasterRecoveryRoleChanged: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T04:32:58.722Z"); return t}()),
	// 			RemoteDisasterRecoveryConfiguration: &armoracledatabase.DisasterRecoveryConfigurationDetails{
	// 				DisasterRecoveryType: to.Ptr(armoracledatabase.DisasterRecoveryTypeAdg),
	// 				TimeSnapshotStandbyEnabledTill: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T04:32:58.715Z"); return t}()),
	// 				IsSnapshotStandby: to.Ptr(true),
	// 				IsReplicateAutomaticBackups: to.Ptr(true),
	// 			},
	// 			LocalStandbyDb: &armoracledatabase.AutonomousDatabaseStandbySummary{
	// 				LagTimeInSeconds: to.Ptr[int32](13),
	// 				LifecycleState: to.Ptr(armoracledatabase.AutonomousDatabaseLifecycleStateProvisioning),
	// 				LifecycleDetails: to.Ptr("zoiyaaibuuhm"),
	// 				TimeDataGuardRoleChanged: to.Ptr("inggk"),
	// 				TimeDisasterRecoveryRoleChanged: to.Ptr("q"),
	// 			},
	// 			FailedDataRecoveryInSeconds: to.Ptr[int32](27),
	// 			IsMtlsConnectionRequired: to.Ptr(true),
	// 			LicenseModel: to.Ptr(armoracledatabase.LicenseModelBringYourOwnLicense),
	// 			LifecycleDetails: to.Ptr("success"),
	// 			ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
	// 			LifecycleState: to.Ptr(armoracledatabase.AutonomousDatabaseLifecycleState("Succeeded")),
	// 			ScheduledOperationsList: []*armoracledatabase.ScheduledOperationsType{
	// 				{
	// 					DayOfWeek: &armoracledatabase.DayOfWeek{
	// 						Name: to.Ptr(armoracledatabase.DayOfWeekNameMonday),
	// 					},
	// 					ScheduledStartTime: to.Ptr("lwwvkazgmfremfwhckfb"),
	// 					ScheduledStopTime: to.Ptr("hjwagzxijpiaogulmnmbuqakpqxhkjvaypjqnvbvtjddc"),
	// 				},
	// 			},
	// 			PrivateEndpointIP: to.Ptr("rdlbhw"),
	// 			PrivateEndpointLabel: to.Ptr("worwqllbglhyakksevparfuaivc"),
	// 			OciURL: to.Ptr("https://fake"),
	// 			TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-09T20:44:09.466Z"); return t}()),
	// 			TimeMaintenanceBegin: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T04:32:58.722Z"); return t}()),
	// 			TimeMaintenanceEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T04:32:58.722Z"); return t}()),
	// 			ActualUsedDataStorageSizeInTbs: to.Ptr[float64](8),
	// 			AllocatedStorageSizeInTbs: to.Ptr[float64](20),
	// 			ApexDetails: &armoracledatabase.ApexDetailsType{
	// 				ApexVersion: to.Ptr("scvpjwygbzqzevlztyfvqiaom"),
	// 				OrdsVersion: to.Ptr("djlwvhpipimxaguklshpppjrzasbk"),
	// 			},
	// 			AvailableUpgradeVersions: []*string{
	// 				to.Ptr("dgvzwolnmulrlwzzazgqm"),
	// 			},
	// 			ConnectionStrings: &armoracledatabase.ConnectionStringType{
	// 				AllConnectionStrings: &armoracledatabase.AllConnectionStringType{
	// 					High: to.Ptr("exdinleextbcjinutlkvnqyxhvandtihncykjzrhfdhfrxdarny"),
	// 					Low: to.Ptr("vurudxqtummqqbnidwthmsqgujufjxwfnejdpuxkgyoxlgqhcgsfjcpzaqeioslpehjfashipdsjhkrpdobstvgxsqrgvcrpbiiabhvymdsylqsjedrimqhtttmszlaqyukopuufbtkbtwgdydrvnvkcdqmphwzpcjxlgefzrdajyczzjdpuzvhqvupbvrpvqhzoaalg"),
	// 					Medium: to.Ptr("ishtubsuzgwtkfdqfgyxjlhehiokdvjfhwqhvmgtuksboshulroytcnubtrxxjbgoutftpzeavbldsoqjwmamgfwevhppyyeckythirzvaqujrjaiqnpyvycakhwgtuftmuxavdgdvbqxgsdqwbnqzmrzymwiydhxekenbweaghgvyveuysxqmovwodzwwfrxhtlcegekjk"),
	// 				},
	// 				Dedicated: to.Ptr("okoggzbidoaknwikuqpvepxvvcseukouprpfrldmakztkleeizbjf"),
	// 				High: to.Ptr("pggylyrivfn"),
	// 				Low: to.Ptr("zrjsbtdbfluaipbzgcvvhyuvqoczjneihaiftkfmuvvyoldslgvvpwthieyrcoxvh"),
	// 				Medium: to.Ptr("ebjnwenxvyeinsabrppychqbcawfxgplfacbsizltwfhpdafbkawopppqsxemlnmrbiqlstjupgdmpfcyyxgofmitbdiarrpprhntntqqjklseigycfcpmmlqiznxzliserjppmgfjatnmtbdxqtlbmrmpfbpoxmyffkkoptpayigeeefmqczroouqjxchswffywpsmyqohbyaclhsrwgqyuuyynvxyyzkche"),
	// 				Profiles: []*armoracledatabase.ProfileType{
	// 					{
	// 						ConsumerGroup: to.Ptr(armoracledatabase.ConsumerGroupHigh),
	// 						DisplayName: to.Ptr("mqqdgidxuovxhcwrkanybxzplautekarsxbcbzlkikpmmvjrdrrkncbamdtcuksplamigrdkydjbzeurbmjgehgppovxqhuzasduwptrlyaurzszzqpztckhpdniepaglzeublbwffxebfespqyfpljlutregvlzzjo"),
	// 						HostFormat: to.Ptr(armoracledatabase.HostFormatTypeFqdn),
	// 						IsRegional: to.Ptr(true),
	// 						Protocol: to.Ptr(armoracledatabase.ProtocolTypeTCP),
	// 						SessionMode: to.Ptr(armoracledatabase.SessionModeTypeDirect),
	// 						SyntaxFormat: to.Ptr(armoracledatabase.SyntaxFormatTypeLong),
	// 						TLSAuthentication: to.Ptr(armoracledatabase.TLSAuthenticationTypeServer),
	// 						Value: to.Ptr("bdrnwqpzbbzdipqqhnroxiuewqg"),
	// 					},
	// 				},
	// 			},
	// 			ConnectionUrls: &armoracledatabase.ConnectionURLType{
	// 				ApexURL: to.Ptr("epnebmudvzijxrfgabsdjewqfotqjmnxvokfhlyklhvtrjpprnqujthmceuhpfuumcbfxktppfguqduzkukxqkofoyyycljjtruyjtoiesxlrwwzonozaxetzrkpmzwasyvryvkryawxxf"),
	// 				DatabaseTransformsURL: to.Ptr("hujiemysucgdgtasazsdtwnxmtjppugunrqnzfzneatukuyzvkfseusjaxrourznsrwxjbvzfansdcyfxnvcyghl"),
	// 				GraphStudioURL: to.Ptr("bucnwmwixwemqqtoozfclfzqenskkyssvcatwbptsezpzdwgnaexgxutsvaibnkawyohqklnktzlhdbhbstm"),
	// 				MachineLearningNotebookURL: to.Ptr("vfhnqsrabxcrjnpqaqkgnpwhxffsqkrgcijdkkvnaoangzkcbgwklufujhmlgydxueybugxzgokxbbappdslttpdthhbmxrgcicqzyjyahjeiqopuglgbjfbhufuvsogquelagbjtyotwhmecwupooitcaftldxjycgfnlilrnicqjxnsucieftadjbvptzltmgqkxhttfkkbutaxvtfzbvbbxbmpxeeyfethpofnmbbqbtlqvnfgelvtjizckgixpptkilcvrntknusvppgnobokjpepynndswcqsnewhfnlxgmownfwfnokhbqulzyuessvxxtcdcnmumbbpjchmjbvjecbbinjolmuoaixzunawlxnoqbpzkczdsubpqpdltnfydwevearrdirzaszsudcxaspozeop"),
	// 				MongoDbURL: to.Ptr("dzmsqtcgsrdgwjlnrfmzcqcrkdqwmjrccxsszwdgpcygywnuurklwthgonxcnwaqcgzoexnaanwzsqwemcijuzxqbrkpvydizjraicgnspizwwnwureyey"),
	// 				OrdsURL: to.Ptr("lmqdgziantbczaneiqxopnaexcroelkbcgggjipzqfhoduwqodoyeghzjyuyhesewopbujxnoiziidhslxdawrfayjvxzjwfobtjrepldlmwhauiurzhbpyxsbueugddmdfindxsdjddqamwbptozzmobugnpezxyxdopripljdwnogjgrjwjfbugdkqrjwzjvavsdmgvspnzcgcjzxauqxqljmfgrtyohfppbmprmexiirvlmymxzyciaraihbwihvahyaciv"),
	// 				SQLDevWebURL: to.Ptr("pktklsvlakmblcakipxy"),
	// 			},
	// 			DataSafeStatus: to.Ptr(armoracledatabase.DataSafeStatusTypeRegistering),
	// 			DatabaseEdition: to.Ptr(armoracledatabase.DatabaseEditionTypeEnterpriseEdition),
	// 			AutonomousDatabaseID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/autonomousDatabases/databasedb1"),
	// 			InMemoryAreaInGbs: to.Ptr[int32](29),
	// 			NextLongTermBackupTimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T04:32:58.722Z"); return t}()),
	// 			LongTermBackupSchedule: &armoracledatabase.LongTermBackUpScheduleDetails{
	// 				RepeatCadence: to.Ptr(armoracledatabase.RepeatCadenceTypeOneTime),
	// 				TimeOfBackup: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T04:32:58.715Z"); return t}()),
	// 				RetentionPeriodInDays: to.Ptr[int32](188),
	// 				IsDisabled: to.Ptr(true),
	// 			},
	// 			IsPreview: to.Ptr(true),
	// 			LocalAdgAutoFailoverMaxDataLossLimit: to.Ptr[int32](1759),
	// 			MemoryPerOracleComputeUnitInGbs: to.Ptr[int32](17),
	// 			OpenMode: to.Ptr(armoracledatabase.OpenModeTypeReadOnly),
	// 			OperationsInsightsStatus: to.Ptr(armoracledatabase.OperationsInsightsStatusTypeEnabling),
	// 			PermissionLevel: to.Ptr(armoracledatabase.PermissionLevelTypeRestricted),
	// 			PrivateEndpoint: to.Ptr("zsknxwclpifygotuivnzkyy"),
	// 			ProvisionableCpus: []*int32{
	// 				to.Ptr[int32](11),
	// 			},
	// 			Role: to.Ptr(armoracledatabase.RoleTypePrimary),
	// 			ServiceConsoleURL: to.Ptr("vyqingzejychumpyufkkiskxhxtfdqhnkugercwlzgakttvdgkrxckvqkxdxlrqpgqosgobgjzyfdewpvhqthefbhtbkxeocreutrfbwmxnfircxsyevouxqyprlxqnmwvepamsylfibuppyslcxobxyynunowjqzupsmkeuuvkctrywubqafaorernjupfslwdlemro"),
	// 			SQLWebDeveloperURL: to.Ptr("ljdauohhvgrrnnzuuauxcuehjmugbhyqszoqcrwiwqnbuozzccqpizilcijojlnyhljxcorelkqldumfnzuljhpwqcaymbdyhvokpbqqkwrarqfrobbzsfbacdtqprrwnuqwrphqlqzilyoyehjfclemrfgylxgawydpflybcbdptzvlukuqhvtgrtaahpgkwymrfellwpbaklxllpbhtkyvlkqnumnsicpaknacvnlqqiyvwwxusjrzqthptckfuyjpferpwszwargimrpqjbigszibkukbdqzctyvktqzrgmhooimbuiblruulhrumetcllofxpqtjlongqtxtnpgfxojlnkkzwiqxrksqeepzqcqudimxtsthqljievdfhbzraluasdwaenmnipdgsnkxeqrlpygcggsfgsjqdxaxqagvoamckaosdzzltrjqlxyxbiaifrtjcnboziaweiyqenlafpchpjxxolvmsseubbaedachwgzectncatkeihzvdncmmhhwgravfcezhcjbhlvopzbnmtaytzadgdszsxpvygjewkyksmdyzepmdjqbkrakuaolveenpngakvazigbxoihfepvebfzpdfdmqrpyqgsdaraifdtnyopibqjavwhffvpowizaskrjosdpivqsgalwypptcwrfscpaqarlnjjzedjhutykdlgxnmmdmrbovscpvbwoeimqsfmiiyxewjldifllcqbmcpoixhufxbptbjkzsvmxurdzjksqirvvhejsjahvtohzsnxkxnmdgkgprwygcnpkxpgrnurnujwdtgkquzkcaanpaamhbbloywlzftxdrxzuhxdxudixgzuyckwkudtzcwdcnxqwdsojyefojusbfjftrmshziassecmgluwvcmwkcpurxajmwvbfgalhhnbebovzulefrdmhepjzqicqnilorzrtbqhemnxgkifxjjeksokqnezkekqiotwthuocqehxfwazqqonwpisrmapjjbwrrrorzepqcgjhfurxbmsfhvnnlzymavjdmfwhehhvqjdrccgabtbqkceqervrztnmfhjgjcvdeccksykspvtvgldhwuwtgobygzhdomdzthrbbhoqnbxfaamrcznmqegfbswuhypgsivoylcvtcyccvsbsxzgxqeqdslytyezyzwtokvivbvwrohcaselsenmzfevvrjkpeyymjiaegybeumjiggaszdwcwqzispeobhysxameswplkdcwhgdmjplexxkbxdmardfisbaplwtsqpmtachrrtsakzhpzwtwutnmcsuehstovmdxntvpzrzwmkms"),
	// 			SupportedRegionsToCloneTo: []*string{
	// 				to.Ptr("germanywestcentral"),
	// 			},
	// 			TimeDataGuardRoleChanged: to.Ptr("exsyrzikvlzvulmkpjxxetftlanim"),
	// 			TimeDeletionOfFreeAutonomousDatabase: to.Ptr("jtigszni"),
	// 			TimeLocalDataGuardEnabled: to.Ptr("rxmqyrqrdpfvleeer"),
	// 			TimeOfLastFailover: to.Ptr("vpcasufezytnkepvpibgqckn"),
	// 			TimeOfLastRefresh: to.Ptr("tqnkkzisphpiqapurwmpdlb"),
	// 			TimeOfLastRefreshPoint: to.Ptr("jfzhrbodayonnuwacgvwhovjgno"),
	// 			TimeOfLastSwitchover: to.Ptr("uvpwuwthnnbzmdmteqzboaah"),
	// 			TimeReclamationOfFreeAutonomousDatabase: to.Ptr("kekqxlbxjsehiretiq"),
	// 			UsedDataStorageSizeInGbs: to.Ptr[int32](13),
	// 			UsedDataStorageSizeInTbs: to.Ptr[int32](16),
	// 			Ocid: to.Ptr("ocid1..aaaaa"),
	// 			BackupRetentionPeriodInDays: to.Ptr[int32](1),
	// 			WhitelistedIPs: []*string{
	// 				to.Ptr("kfierlppwurtqrhfxwgfgrnqtmvraignzwsddwmpdijeveuevuoejfmbjvpnlrmmdflilxcwkkzvdofctsdjfxrrrwctihhnchtrouauesqbmlcqhzwnppnhrtitecenlfyshassvajukbwxudhlwixkvkgsessvshcwmleoqujeemwenhwlsccbcjnnviugzgylsxkssalqoicatcvkahogdvweymhdxboyqwhaxuzlmrdbvgbnnetobkbwygcsflzanwknlybvvzgjzjirpfrksbxwgllgfxwdflcisvxpkjecpgdaxccqkzxofedkrawvhzeabmunpykwd"),
	// 				to.Ptr("1.1.1.0/24"),
	// 				to.Ptr("1.1.2.25"),
	// 			},
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		Tags: map[string]*string{
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/autonomousDatabases/databasedb1"),
	// 		Name: to.Ptr("databasedb1"),
	// 		Type: to.Ptr("Oracle.Database/autonomousDatabases"),
	// 		SystemData: &armoracledatabase.SystemData{
	// 			CreatedBy: to.Ptr("sqehacivpuim"),
	// 			CreatedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T04:32:58.716Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("axrqfdkqylvjv"),
	// 			LastModifiedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T04:32:58.716Z"); return t}()),
	// 		},
	// 	},
	// }
}
