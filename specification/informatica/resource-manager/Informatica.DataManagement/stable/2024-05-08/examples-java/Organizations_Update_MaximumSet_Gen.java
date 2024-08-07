
import com.azure.resourcemanager.informaticadatamanagement.models.CompanyDetailsUpdate;
import com.azure.resourcemanager.informaticadatamanagement.models.InformaticaOrganizationResource;
import com.azure.resourcemanager.informaticadatamanagement.models.MarketplaceDetailsUpdate;
import com.azure.resourcemanager.informaticadatamanagement.models.OfferDetailsUpdate;
import com.azure.resourcemanager.informaticadatamanagement.models.OrganizationPropertiesCustomUpdate;
import com.azure.resourcemanager.informaticadatamanagement.models.UserDetailsUpdate;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Organizations Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/
     * Organizations_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_Update.
     * 
     * @param manager Entry point to InformaticaDataManagementManager.
     */
    public static void organizationsUpdate(
        com.azure.resourcemanager.informaticadatamanagement.InformaticaDataManagementManager manager) {
        InformaticaOrganizationResource resource = manager.organizations()
            .getByResourceGroupWithResponse("rgopenapi", "_-", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("key1918", "fakeTokenPlaceholder")).withProperties(
            new OrganizationPropertiesCustomUpdate().withMarketplaceDetails(new MarketplaceDetailsUpdate()
                .withMarketplaceSubscriptionId("szhyxzgjtssjmlguivepc")
                .withOfferDetails(new OfferDetailsUpdate().withPublisherId(
                    "ktzfghsyjqbsswhltoaemgotmnorhdogvkaxplutbjjqzuepxizliynyakersobagvpwvpzwjtjjxigsqgcyqaahaxdijghnexliofhfjlqzjmmbvrhcvjxdodnexxizbgfhjopbwzjojxsluasnwwsgcajefglbcvzpaeblanhmurcculndtfwnfjyxol")
                    .withOfferId("idaxbflabvjsippplyenvrpgeydsjxcmyubgukffkcdvlvrtwpdhnvdblxjsldiuswrchsibk")
                    .withPlanId("giihvvnwdwzkfqrhkpqzbgfotzyixnsvmxzauseebillhslauglzfxzvzvts")
                    .withPlanName(
                        "tfqjenotaewzdeerliteqxdawuqxhwdzbtiiimsaedrlsnbdoonnloakjtvnwhhrcyxxsgoachguthqvlahpjyofpoqpfacfmiaauawazkmxkjgvktbptojknzojtjrfzvbbjjkvstabqyaczxinijhoxrjukftsagpwgsvpmczopztmplipyufhuaumfx")
                    .withTermUnit("nykqoplazujcwmfldntifjqrnx")
                    .withTermId("eolmwogtgpdncqoigqcdomupwummaicwvdxgbskpdsmjizdfbdgbxbuekcpwmenqzbhqxpdnjtup")))
                .withUserDetails(new UserDetailsUpdate().withFirstName("qguqrmanyupoi").withLastName("ugzg")
                    .withEmailAddress("7_-46@13D--3.m-4x-.11.c-9-.DHLYFc").withUpn("viwjrkn").withPhoneNumber("uxa"))
                .withCompanyDetails(new CompanyDetailsUpdate().withCompanyName("xkrvbozrjcvappqeeyt")
                    .withOfficeAddress("sfcx").withCountry("rvlzppgvopcw").withDomain("dponvwnrdrnzahcurqssesukbsokdd")
                    .withBusiness("mwqblnruflwpolgbxpqbqneve").withNumberOfEmployees(22))
                .withExistingResourceId(
                    "/subscriptions/subid/resourceGroups/rg1/providers/Informatica.DataManagement/organizations/org1/serverlessRuntimes/serverlessRuntimeName"))
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
