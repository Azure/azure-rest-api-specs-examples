import com.azure.resourcemanager.devcenter.models.ManagedServiceIdentity;
import com.azure.resourcemanager.devcenter.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.devcenter.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/** Samples for DevCenters CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/DevCenters_CreateWithUserIdentity.json
     */
    /**
     * Sample code: DevCenters_CreateWithUserIdentity.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void devCentersCreateWithUserIdentity(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager
            .devCenters()
            .define("Contoso")
            .withRegion("centralus")
            .withExistingResourceGroup("rg1")
            .withTags(mapOf("CostCode", "12345"))
            .withIdentity(
                new ManagedServiceIdentity()
                    .withType(ManagedServiceIdentityType.USER_ASSIGNED)
                    .withUserAssignedIdentities(
                        mapOf(
                            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/identityGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testidentity1",
                            new UserAssignedIdentity())))
            .create();
    }

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
