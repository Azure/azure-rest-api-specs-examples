package armdiscovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/discovery/armdiscovery"
)

// Generated from example definition: 2026-06-01/Projects_Update_MaximumSet_Gen.json
func ExampleProjectsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdiscovery.NewClientFactory("A54D43BD-2F5F-4BB1-95D4-9A8D23CC7DD4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProjectsClient().BeginUpdate(ctx, "rgdiscovery", "f20325b2d0d95502c7", "84859fdc69fda61de3", armdiscovery.Project{
		Properties: &armdiscovery.ProjectProperties{
			StorageContainerIDs: []*string{
				to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/resourceGroups/rgdiscovery/providers/Microsoft.Discovery/storageContainers/storageContainer12"),
			},
			Settings: &armdiscovery.ProjectSettings{
				BehaviorPreferences: to.Ptr("obmbzlehsieuwdzpaezckgxgxpttlekvszgxidurhzshsocvzfetyamavxmvgbffbooawggvvdnbanshghtgaelkmnvdwiyxralhqjqoifwivnyudifcxocydkxnkqskmqlfxgnddjcxajkckdtukmtomqutvdapvgpimpoxyhiorbfocptcdjguejlpwwvfhhcywouuueclnsijqnipnhsryhwefwaxzwbsyvhhqtmlaupajoiymmorguwwhnyppjdvksjctsovopmxrqjsvunifvqnlgvxlpzibkwgjpzmbnvwzbsegifwvtpcyqaelxbwuzfowummpaowxrosuilphuckyizxsflyipbulsaxgjksyhpzshqvuplidvmxsdebuxrcrcwxuxhkdawtrpgrtjjevfaybgclxrisajtsxdogilvfxqtavvmbpiqrgfwuluthnlpnibybgutrjdjykoazfvhfbeugrposbhbnmvymystnsvtqyytgylwdfurfvbvtimnxvsfvvgrmbzaebdjdgazmfjevfchsrmkdsejnuhjagylaxnbdniumthqqsiytyybfbxrqgdkyrjejzxdysovhersuklhtdbhbteawgyspwadijhhlinzolskxdnkqnfyppnsrwowqlrbupsaizuhyyiueffmmdxuqwxhnifmiisfflgonrdpccgzzwmitmladixlnvrbpwlnecminutgyxbvploxeiljrzmorxuvgsibedcolxbtvcspforsqpjmnyoxlriecmpxshdkmqrpvjovoocslfgmcwlpkautcupwpwxoykfgubztgmiynxjmmdcwlcjyoehsnljgrififinrhpazsqbjbmazlvspsxmqjwbowphefrfamqhsbpcsrjwcarzxuckdicnwagbsjblfqtizlcbcloxpnnonqepaasxpchihtjxjcsoqlmlyuixffjepnkhcwauazlobbgoopnoservazcndgrizxqdyfvgzebpgwwxanmjwnqhwevclvamfvzmvgfunpswapumcealprovgtqnuduyhrgwjxvmouguxxdkslfpjgtefyiinyhiryycexbwwecgylfxvcouzldvlnkcyziznifoerfxmphdamvjvjeollxzjvxngznewcwgpciyouericwmyiacfinbybqqzmszlmnkmjloswhcyhmafmisykxrebkbkhqfjlmjqbhqqsflkfgkbmzefykgqouqkbsvzgqqtvyrbhsqoiuijjzxkxrhykhgintrrrarjxeyorvezwjaqurqpoqfzqpwiegektkpzopedpzpbclcrxwtqwrnrwrnluvpmyjqoafipygvdwtqrgkhuixdofbmmnfzgjzaqvvvqyqloraumziiryixxakubzpdptvncdbgmdwkpnctclcavmpwbfopogcaknicrrvpizwkkzwuugapvxddzgxkfxxzlruxqjstflfrwtuksvagrtjpnwaluinivxszjloewyzxstheiqgdijmfhzjilxnxmijmjttmgtpiqvyjjcfmzgkhyhmloloxzpqbwbvoihxpksokhkfwtgcfajzmalohzvzddogrbcryjtljxhubwdrlpkhfmhotucplhfigfevbdzcnwgjecfqpoetivmlrcxrnoerydbxzpdefpdvflonvauvnubwefnzxciczhsolruoarxbxsajubnsrwfugauoawodribtjogctqokckqoyvukongairkeolvjxesznlbvnzdvzhuyelhldkslubefmlnobhsotgbewtxjackyedifplyflxatdnenmcghbplapmmdoqxlrljrzfbjqvuyaocwqtfqenvjtictmrverpcndfradmmkdzvzudnwkcawzlbpnvhqlldywtncrmtclxyusnspukmbvqjppzkeritwipziaypqwqpcldrrvnvreklbvctsgimbaprmpbdzstugagjmszspdkuqovryvgwduqydwzaexsoxazigkukcihhrygzslcyfgzdgtvrwtzojuecyzufwyropitlabebanyeexiqtfdemghbhshigwxkarfapwfvqajrvahfbnrdhqrzycvmisewcfiweydwaqgbsypqxkgcoyyihvwxeqhxrcqhngrixfytrorynriugviiswpdcgglihaxucyswezwpkvubuocgicgarbheijlwiqsfqpubyuiukcxjmsxyrwmhimveovtwkrrnlfdudsxyhgnuggnjxgwrbdmlltnbtaqvahwdisidxjjrepymalxfgcmrzsixudqznloxkjrslrzcjjjwtoahiopohusiqywtajlmduguuiwqperiqopdssmafsijgwambxfngwmcmxsiujjudlluldcjgvpjmyzibzvadmeuiskpkkblgroizqspxexhyuoesdvalfjwtzcjpmqyuwlhjwhfeiqdncrhpvpcaacrpuwnzmbjazdbvpwhpxrdjhymrodbvmkxjxptmfhcjdywacxlnwrquaxplsmqkhyrtihqvzmjgyhqcxlcbdyfjvtdprwpubnmataqgumwqvttkiqhxksjtkcytfustjaxosiucxadxzmoagvmsxnflzcvhjzwrohgnwmqbmlruxbvhlqlsxoxxmmqaoknawlidnbbgnqofzqiihxkijzlelbjzvrxlxzcnxyxjjabqiokgfvwhmeyikdwwyhmyxjk"),
			},
		},
		Tags: map[string]*string{
			"key2710": to.Ptr("ghqxvtoxhybuwft"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdiscovery.ProjectsClientUpdateResponse{
	// 	Project: armdiscovery.Project{
	// 		ID: to.Ptr("/subscriptions/A54D43BD-2F5F-4BB1-95D4-9A8D23CC7DD4/resourceGroups/rgdiscovery/providers/Microsoft.Discovery/workspaces/f20325b2d0d95502c7/projects/84859fdc69fda61de3"),
	// 		Name: to.Ptr("84859fdc69fda61de3"),
	// 		Tags: map[string]*string{
	// 			"key4676": to.Ptr("webgttjblqpznhjzqef"),
	// 		},
	// 		Location: to.Ptr("uksouth"),
	// 		Type: to.Ptr("Microsoft.Discovery/workspaces/projects"),
	// 		SystemData: &armdiscovery.SystemData{
	// 			CreatedBy: to.Ptr("uymdmmhvojuqtvvxokgefohqpcjw"),
	// 			CreatedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("ucuttxilomgszapozsuit"),
	// 			LastModifiedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
	// 		},
	// 		Properties: &armdiscovery.ProjectProperties{
	// 			ProvisioningState: to.Ptr(armdiscovery.ProvisioningStateSucceeded),
	// 			FoundryProjectEndpoint: to.Ptr("https://microsoft.com/a"),
	// 			StorageContainerIDs: []*string{
	// 				to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/resourceGroups/rgdiscovery/providers/Microsoft.Discovery/storageContainers/storageContainer12"),
	// 			},
	// 			Settings: &armdiscovery.ProjectSettings{
	// 				BehaviorPreferences: to.Ptr("obmbzlehsieuwdzpaezckgxgxpttlekvszgxidurhzshsocvzfetyamavxmvgbffbooawggvvdnbanshghtgaelkmnvdwiyxralhqjqoifwivnyudifcxocydkxnkqskmqlfxgnddjcxajkckdtukmtomqutvdapvgpimpoxyhiorbfocptcdjguejlpwwvfhhcywouuueclnsijqnipnhsryhwefwaxzwbsyvhhqtmlaupajoiymmorguwwhnyppjdvksjctsovopmxrqjsvunifvqnlgvxlpzibkwgjpzmbnvwzbsegifwvtpcyqaelxbwuzfowummpaowxrosuilphuckyizxsflyipbulsaxgjksyhpzshqvuplidvmxsdebuxrcrcwxuxhkdawtrpgrtjjevfaybgclxrisajtsxdogilvfxqtavvmbpiqrgfwuluthnlpnibybgutrjdjykoazfvhfbeugrposbhbnmvymystnsvtqyytgylwdfurfvbvtimnxvsfvvgrmbzaebdjdgazmfjevfchsrmkdsejnuhjagylaxnbdniumthqqsiytyybfbxrqgdkyrjejzxdysovhersuklhtdbhbteawgyspwadijhhlinzolskxdnkqnfyppnsrwowqlrbupsaizuhyyiueffmmdxuqwxhnifmiisfflgonrdpccgzzwmitmladixlnvrbpwlnecminutgyxbvploxeiljrzmorxuvgsibedcolxbtvcspforsqpjmnyoxlriecmpxshdkmqrpvjovoocslfgmcwlpkautcupwpwxoykfgubztgmiynxjmmdcwlcjyoehsnljgrififinrhpazsqbjbmazlvspsxmqjwbowphefrfamqhsbpcsrjwcarzxuckdicnwagbsjblfqtizlcbcloxpnnonqepaasxpchihtjxjcsoqlmlyuixffjepnkhcwauazlobbgoopnoservazcndgrizxqdyfvgzebpgwwxanmjwnqhwevclvamfvzmvgfunpswapumcealprovgtqnuduyhrgwjxvmouguxxdkslfpjgtefyiinyhiryycexbwwecgylfxvcouzldvlnkcyziznifoerfxmphdamvjvjeollxzjvxngznewcwgpciyouericwmyiacfinbybqqzmszlmnkmjloswhcyhmafmisykxrebkbkhqfjlmjqbhqqsflkfgkbmzefykgqouqkbsvzgqqtvyrbhsqoiuijjzxkxrhykhgintrrrarjxeyorvezwjaqurqpoqfzqpwiegektkpzopedpzpbclcrxwtqwrnrwrnluvpmyjqoafipygvdwtqrgkhuixdofbmmnfzgjzaqvvvqyqloraumziiryixxakubzpdptvncdbgmdwkpnctclcavmpwbfopogcaknicrrvpizwkkzwuugapvxddzgxkfxxzlruxqjstflfrwtuksvagrtjpnwaluinivxszjloewyzxstheiqgdijmfhzjilxnxmijmjttmgtpiqvyjjcfmzgkhyhmloloxzpqbwbvoihxpksokhkfwtgcfajzmalohzvzddogrbcryjtljxhubwdrlpkhfmhotucplhfigfevbdzcnwgjecfqpoetivmlrcxrnoerydbxzpdefpdvflonvauvnubwefnzxciczhsolruoarxbxsajubnsrwfugauoawodribtjogctqokckqoyvukongairkeolvjxesznlbvnzdvzhuyelhldkslubefmlnobhsotgbewtxjackyedifplyflxatdnenmcghbplapmmdoqxlrljrzfbjqvuyaocwqtfqenvjtictmrverpcndfradmmkdzvzudnwkcawzlbpnvhqlldywtncrmtclxyusnspukmbvqjppzkeritwipziaypqwqpcldrrvnvreklbvctsgimbaprmpbdzstugagjmszspdkuqovryvgwduqydwzaexsoxazigkukcihhrygzslcyfgzdgtvrwtzojuecyzufwyropitlabebanyeexiqtfdemghbhshigwxkarfapwfvqajrvahfbnrdhqrzycvmisewcfiweydwaqgbsypqxkgcoyyihvwxeqhxrcqhngrixfytrorynriugviiswpdcgglihaxucyswezwpkvubuocgicgarbheijlwiqsfqpubyuiukcxjmsxyrwmhimveovtwkrrnlfdudsxyhgnuggnjxgwrbdmlltnbtaqvahwdisidxjjrepymalxfgcmrzsixudqznloxkjrslrzcjjjwtoahiopohusiqywtajlmduguuiwqperiqopdssmafsijgwambxfngwmcmxsiujjudlluldcjgvpjmyzibzvadmeuiskpkkblgroizqspxexhyuoesdvalfjwtzcjpmqyuwlhjwhfeiqdncrhpvpcaacrpuwnzmbjazdbvpwhpxrdjhymrodbvmkxjxptmfhcjdywacxlnwrquaxplsmqkhyrtihqvzmjgyhqcxlcbdyfjvtdprwpubnmataqgumwqvttkiqhxksjtkcytfustjaxosiucxadxzmoagvmsxnflzcvhjzwrohgnwmqbmlruxbvhlqlsxoxxmmqaoknawlidnbbgnqofzqiihxkijzlelbjzvrxlxzcnxyxjjabqiokgfvwhmeyikdwwyhmyxjk"),
	// 			},
	// 		},
	// 	},
	// }
}
