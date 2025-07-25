
import com.azure.resourcemanager.hardwaresecuritymodules.models.CloudHsmCluster;
import com.azure.resourcemanager.hardwaresecuritymodules.models.ManagedServiceIdentity;
import com.azure.resourcemanager.hardwaresecuritymodules.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.hardwaresecuritymodules.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for CloudHsmClusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-31/CloudHsmCluster_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: CloudHsmCluster_Update_MaximumSet_Gen.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void cloudHsmClusterUpdateMaximumSetGen(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        CloudHsmCluster resource = manager.cloudHsmClusters()
            .getByResourceGroupWithResponse("rgcloudhsm", "chsm1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("Dept", "hsm", "Environment", "dogfood", "Slice", "A"))
            .withIdentity(new ManagedServiceIdentity().withType(ManagedServiceIdentityType.USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resources/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity-1",
                    new UserAssignedIdentity())))
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
