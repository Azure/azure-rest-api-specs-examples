
import com.azure.resourcemanager.recoveryservices.models.Vault;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Vaults Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/PATCHVault.
     * json
     */
    /**
     * Sample code: Update Resource.
     * 
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void updateResource(com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        Vault resource = manager.vaults()
            .getByResourceGroupWithResponse("HelloWorld", "swaggerExample", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("PatchKey", "fakeTokenPlaceholder")).apply();
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
