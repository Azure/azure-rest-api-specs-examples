
import com.azure.resourcemanager.discovery.models.Project;
import com.azure.resourcemanager.discovery.models.ProjectProperties;
import com.azure.resourcemanager.discovery.models.ProjectSettings;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Projects Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/Projects_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: Projects_Update_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void projectsUpdateMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        Project resource = manager.projects().getWithResponse("rgdiscovery", "f20325b2d0d95502c7", "84859fdc69fda61de3",
            com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("key2710", "fakeTokenPlaceholder")).withProperties(new ProjectProperties()
            .withStorageContainerIds(Arrays.asList(
                "/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/resourceGroups/rgdiscovery/providers/Microsoft.Discovery/storageContainers/storageContainer12"))
            .withSettings(new ProjectSettings().withBehaviorPreferences(
                "obmbzlehsieuwdzpaezckgxgxpttlekvszgxidurhzshsocvzfetyamavxmvgbffbooawggvvdnbanshghtgaelkmnvdwiyxralhqjqoifwivnyudifcxocydkxnkqskmqlfxgnddjcxajkckdtukmtomqutvdapvgpimpoxyhiorbfocptcdjguejlpwwvfhhcywouuueclnsijqnipnhsryhwefwaxzwbsyvhhqtmlaupajoiymmorguwwhnyppjdvksjctsovopmxrqjsvunifvqnlgvxlpzibkwgjpzmbnvwzbsegifwvtpcyqaelxbwuzfowummpaowxrosuilphuckyizxsflyipbulsaxgjksyhpzshqvuplidvmxsdebuxrcrcwxuxhkdawtrpgrtjjevfaybgclxrisajtsxdogilvfxqtavvmbpiqrgfwuluthnlpnibybgutrjdjykoazfvhfbeugrposbhbnmvymystnsvtqyytgylwdfurfvbvtimnxvsfvvgrmbzaebdjdgazmfjevfchsrmkdsejnuhjagylaxnbdniumthqqsiytyybfbxrqgdkyrjejzxdysovhersuklhtdbhbteawgyspwadijhhlinzolskxdnkqnfyppnsrwowqlrbupsaizuhyyiueffmmdxuqwxhnifmiisfflgonrdpccgzzwmitmladixlnvrbpwlnecminutgyxbvploxeiljrzmorxuvgsibedcolxbtvcspforsqpjmnyoxlriecmpxshdkmqrpvjovoocslfgmcwlpkautcupwpwxoykfgubztgmiynxjmmdcwlcjyoehsnljgrififinrhpazsqbjbmazlvspsxmqjwbowphefrfamqhsbpcsrjwcarzxuckdicnwagbsjblfqtizlcbcloxpnnonqepaasxpchihtjxjcsoqlmlyuixffjepnkhcwauazlobbgoopnoservazcndgrizxqdyfvgzebpgwwxanmjwnqhwevclvamfvzmvgfunpswapumcealprovgtqnuduyhrgwjxvmouguxxdkslfpjgtefyiinyhiryycexbwwecgylfxvcouzldvlnkcyziznifoerfxmphdamvjvjeollxzjvxngznewcwgpciyouericwmyiacfinbybqqzmszlmnkmjloswhcyhmafmisykxrebkbkhqfjlmjqbhqqsflkfgkbmzefykgqouqkbsvzgqqtvyrbhsqoiuijjzxkxrhykhgintrrrarjxeyorvezwjaqurqpoqfzqpwiegektkpzopedpzpbclcrxwtqwrnrwrnluvpmyjqoafipygvdwtqrgkhuixdofbmmnfzgjzaqvvvqyqloraumziiryixxakubzpdptvncdbgmdwkpnctclcavmpwbfopogcaknicrrvpizwkkzwuugapvxddzgxkfxxzlruxqjstflfrwtuksvagrtjpnwaluinivxszjloewyzxstheiqgdijmfhzjilxnxmijmjttmgtpiqvyjjcfmzgkhyhmloloxzpqbwbvoihxpksokhkfwtgcfajzmalohzvzddogrbcryjtljxhubwdrlpkhfmhotucplhfigfevbdzcnwgjecfqpoetivmlrcxrnoerydbxzpdefpdvflonvauvnubwefnzxciczhsolruoarxbxsajubnsrwfugauoawodribtjogctqokckqoyvukongairkeolvjxesznlbvnzdvzhuyelhldkslubefmlnobhsotgbewtxjackyedifplyflxatdnenmcghbplapmmdoqxlrljrzfbjqvuyaocwqtfqenvjtictmrverpcndfradmmkdzvzudnwkcawzlbpnvhqlldywtncrmtclxyusnspukmbvqjppzkeritwipziaypqwqpcldrrvnvreklbvctsgimbaprmpbdzstugagjmszspdkuqovryvgwduqydwzaexsoxazigkukcihhrygzslcyfgzdgtvrwtzojuecyzufwyropitlabebanyeexiqtfdemghbhshigwxkarfapwfvqajrvahfbnrdhqrzycvmisewcfiweydwaqgbsypqxkgcoyyihvwxeqhxrcqhngrixfytrorynriugviiswpdcgglihaxucyswezwpkvubuocgicgarbheijlwiqsfqpubyuiukcxjmsxyrwmhimveovtwkrrnlfdudsxyhgnuggnjxgwrbdmlltnbtaqvahwdisidxjjrepymalxfgcmrzsixudqznloxkjrslrzcjjjwtoahiopohusiqywtajlmduguuiwqperiqopdssmafsijgwambxfngwmcmxsiujjudlluldcjgvpjmyzibzvadmeuiskpkkblgroizqspxexhyuoesdvalfjwtzcjpmqyuwlhjwhfeiqdncrhpvpcaacrpuwnzmbjazdbvpwhpxrdjhymrodbvmkxjxptmfhcjdywacxlnwrquaxplsmqkhyrtihqvzmjgyhqcxlcbdyfjvtdprwpubnmataqgumwqvttkiqhxksjtkcytfustjaxosiucxadxzmoagvmsxnflzcvhjzwrohgnwmqbmlruxbvhlqlsxoxxmmqaoknawlidnbbgnqofzqiihxkijzlelbjzvrxlxzcnxyxjjabqiokgfvwhmeyikdwwyhmyxjk")))
            .apply();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
