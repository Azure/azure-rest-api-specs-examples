
import com.azure.resourcemanager.astro.models.LiftrBaseDataPartnerOrganizationPropertiesUpdate;
import com.azure.resourcemanager.astro.models.LiftrBaseSingleSignOnProperties;
import com.azure.resourcemanager.astro.models.LiftrBaseUserDetailsUpdate;
import com.azure.resourcemanager.astro.models.ManagedServiceIdentity;
import com.azure.resourcemanager.astro.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.astro.models.OrganizationResource;
import com.azure.resourcemanager.astro.models.OrganizationResourceUpdateProperties;
import com.azure.resourcemanager.astro.models.SingleSignOnStates;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Organizations Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/liftrastronomer/resource-manager/Astronomer.Astro/stable/2023-08-01/examples/
     * Organizations_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_Update.
     * 
     * @param manager Entry point to AstroManager.
     */
    public static void organizationsUpdate(com.azure.resourcemanager.astro.AstroManager manager) {
        OrganizationResource resource = manager.organizations()
            .getByResourceGroupWithResponse("rgastronomer", "6.", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("key1474", "fakeTokenPlaceholder"))
            .withIdentity(new ManagedServiceIdentity().withType(ManagedServiceIdentityType.NONE)
                .withUserAssignedIdentities(mapOf()))
            .withProperties(new OrganizationResourceUpdateProperties()
                .withUser(new LiftrBaseUserDetailsUpdate().withFirstName("qeuofehzypzljgcuysugefbgxde")
                    .withLastName("g").withEmailAddress(".K_@e7N-g1.xjqnbPs").withUpn("uwtprzdfpsqmktx")
                    .withPhoneNumber("aqpyxznvqpgkzohevynofrjdfgoo"))
                .withPartnerOrganizationProperties(new LiftrBaseDataPartnerOrganizationPropertiesUpdate()
                    .withOrganizationId("lrtmbkvyvvoszhjevohkmyjhfyty").withWorkspaceId("xsepuskdhejaadusyxq")
                    .withOrganizationName("U2P_").withWorkspaceName("L.-y_--:")
                    .withSingleSignOnProperties(new LiftrBaseSingleSignOnProperties()
                        .withSingleSignOnState(SingleSignOnStates.INITIAL).withEnterpriseAppId("mklfypyujwumgwdzae")
                        .withSingleSignOnUrl("ymmtzkyghvinvhgnqlzwrr").withAadDomains(Arrays.asList("kfbleh")))))
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
