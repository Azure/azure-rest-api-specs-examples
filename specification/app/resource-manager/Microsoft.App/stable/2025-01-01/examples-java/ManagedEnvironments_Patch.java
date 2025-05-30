
import com.azure.resourcemanager.appcontainers.models.ManagedEnvironment;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ManagedEnvironments Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/ManagedEnvironments_Patch.json
     */
    /**
     * Sample code: Patch Managed Environment.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        patchManagedEnvironment(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        ManagedEnvironment resource = manager.managedEnvironments()
            .getByResourceGroupWithResponse("examplerg", "testcontainerenv", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("tag1", "value1", "tag2", "value2")).apply();
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
