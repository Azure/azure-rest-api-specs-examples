
import com.azure.resourcemanager.mongodbatlas.models.ManagedServiceIdentity;
import com.azure.resourcemanager.mongodbatlas.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.mongodbatlas.models.MarketplaceDetails;
import com.azure.resourcemanager.mongodbatlas.models.OfferDetails;
import com.azure.resourcemanager.mongodbatlas.models.OrganizationProperties;
import com.azure.resourcemanager.mongodbatlas.models.PartnerProperties;
import com.azure.resourcemanager.mongodbatlas.models.UserDetails;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Organizations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Organizations_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_CreateOrUpdate_MaximumSet.
     * 
     * @param manager Entry point to MongoDBAtlasManager.
     */
    public static void
        organizationsCreateOrUpdateMaximumSet(com.azure.resourcemanager.mongodbatlas.MongoDBAtlasManager manager) {
        manager.organizations().define("U.1-:7").withRegion("wobqn").withExistingResourceGroup("rgopenapi")
            .withTags(mapOf())
            .withProperties(new OrganizationProperties()
                .withMarketplace(new MarketplaceDetails().withSubscriptionId("o")
                    .withOfferDetails(new OfferDetails().withPublisherId("rxglearenxsgpwzlsxmiicynks")
                        .withOfferId("ohnquleylybvjrtnpjupvwlk").withPlanId("obhxnhvrtbcnoovgofbs")
                        .withPlanName("lkwdzpfhvjezjusrqzyftcikxdt").withTermUnit("omkxrnburbnruglwqgjlahvjmbfcse")
                        .withTermId("bqmmltwmtpdcdeszbka")))
                .withUser(new UserDetails().withFirstName("aslybvdwwddqxwazxvxhjrs")
                    .withLastName("cnuitqoqpcyvmuqowgnxpwxjcveyr").withEmailAddress(".K_@e7N-g1.xjqnbPs")
                    .withUpn("howdzmfy").withPhoneNumber("ilypntsrbmbbbexbasuu").withCompanyName("oxdcwwl"))
                .withPartnerProperties(new PartnerProperties().withOrganizationId("lyombjlhvwxithkiy")
                    .withRedirectUrl("cbxwtehraetlluocdihfgchvjzockn").withOrganizationName("U.1-:7")))
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
