
import com.azure.resourcemanager.mongodbatlas.models.ManagedServiceIdentity;
import com.azure.resourcemanager.mongodbatlas.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.mongodbatlas.models.OrganizationProperties;
import com.azure.resourcemanager.mongodbatlas.models.OrganizationResource;
import com.azure.resourcemanager.mongodbatlas.models.PartnerProperties;
import com.azure.resourcemanager.mongodbatlas.models.UserDetails;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Organizations Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-18-preview/Organizations_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_Update_MaximumSet.
     * 
     * @param manager Entry point to MongoDBAtlasManager.
     */
    public static void
        organizationsUpdateMaximumSet(com.azure.resourcemanager.mongodbatlas.MongoDBAtlasManager manager) {
        OrganizationResource resource = manager.organizations()
            .getByResourceGroupWithResponse("rgopenapi", "U.1-:7", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf())
            .withProperties(new OrganizationProperties()
                .withUser(new UserDetails().withFirstName("btyhwmlbzzihjfimviefebg").withLastName("xx")
                    .withEmailAddress(".K_@e7N-g1.xjqnbPs").withUpn("mxtbogd").withPhoneNumber("isvc")
                    .withCompanyName("oztteysco"))
                .withPartnerProperties(new PartnerProperties().withOrganizationId("vugtqrobendjkinziswxlqueouo")
                    .withRedirectUrl("cbxwtehraetlluocdihfgchvjzockn").withOrganizationName("U.1-:7")))
            .withIdentity(new ManagedServiceIdentity().withType(ManagedServiceIdentityType.NONE)
                .withUserAssignedIdentities(mapOf()))
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
