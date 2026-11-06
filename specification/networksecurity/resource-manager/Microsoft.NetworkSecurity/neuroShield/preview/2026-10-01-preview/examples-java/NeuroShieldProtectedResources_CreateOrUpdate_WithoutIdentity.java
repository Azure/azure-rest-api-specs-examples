
import com.azure.resourcemanager.networksecurity.models.NeuroShieldProtectedResourceProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for NeuroShieldProtectedResources CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/NeuroShieldProtectedResources_CreateOrUpdate_WithoutIdentity.json
     */
    /**
     * Sample code: Create or update protected resource without an identity.
     * 
     * @param manager Entry point to NetworkSecurityManager.
     */
    public static void createOrUpdateProtectedResourceWithoutAnIdentity(
        com.azure.resourcemanager.networksecurity.NetworkSecurityManager manager) {
        manager.neuroShieldProtectedResources().define("vm-01")
            .withExistingNeuroShield("myResourceGroup", "myNeuroShield")
            .withProperties(new NeuroShieldProtectedResourceProperties().withTargetResourceId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myComputeRg/providers/Microsoft.Compute/virtualMachines/vm-01"))
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
