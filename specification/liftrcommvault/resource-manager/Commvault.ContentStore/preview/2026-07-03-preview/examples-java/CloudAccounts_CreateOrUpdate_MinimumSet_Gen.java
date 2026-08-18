
import com.azure.resourcemanager.commvaultcontentstore.models.CloudAccountProperties;
import com.azure.resourcemanager.commvaultcontentstore.models.ManagedServiceIdentity;
import com.azure.resourcemanager.commvaultcontentstore.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.commvaultcontentstore.models.MarketplaceDetails;
import com.azure.resourcemanager.commvaultcontentstore.models.OfferDetails;
import com.azure.resourcemanager.commvaultcontentstore.models.UserDetails;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for CloudAccounts CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-03-preview/CloudAccounts_CreateOrUpdate_MinimumSet_Gen.json
     */
    /**
     * Sample code: CloudAccounts_CreateOrUpdate_MinimumSet - CCA create with create-only role bootstrap fields omitted.
     * 
     * @param manager Entry point to CommvaultContentStoreManager.
     */
    public static void cloudAccountsCreateOrUpdateMinimumSetCCACreateWithCreateOnlyRoleBootstrapFieldsOmitted(
        com.azure.resourcemanager.commvaultcontentstore.CommvaultContentStoreManager manager) {
        manager.cloudAccounts().define("sample-cloudAccountName").withRegion("eastus")
            .withExistingResourceGroup("rgcommvault").withTags(mapOf())
            .withProperties(new CloudAccountProperties()
                .withMarketplace(new MarketplaceDetails().withSubscriptionId("tblwyuznrazgchhfczgtlaifwamndt")
                    .withOfferDetails(new OfferDetails().withPublisherId("npghpdbgiohslbbeihxdwucejb")
                        .withOfferId("recofyvhkddgkuvducosjstenmy").withPlanId("pqoyqqavjh")
                        .withPlanName("hwcltkdvndwfmmnthzwvocujri").withTermUnit("wzrzqyfzrpqhy")
                        .withTermId("avpgkctrkwdmudsz")))
                .withUser(
                    new UserDetails().withFirstName("John").withLastName("Doe").withEmailAddress("john.doe@contoso.com")
                        .withUpn("john.doe@contoso.com").withPhoneNumber("1234567890")))
            .withIdentity(new ManagedServiceIdentity().withType(ManagedServiceIdentityType.NONE)
                .withUserAssignedIdentities(mapOf()))
            .create();
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
