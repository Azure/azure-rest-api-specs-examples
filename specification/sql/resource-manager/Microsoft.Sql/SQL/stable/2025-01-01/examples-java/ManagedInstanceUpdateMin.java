
import com.azure.resourcemanager.sql.models.ManagedInstanceUpdate;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ManagedInstances Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceUpdateMin.json
     */
    /**
     * Sample code: Update managed instance with minimal properties.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        updateManagedInstanceWithMinimalProperties(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstances().update("testrg", "testinstance",
            new ManagedInstanceUpdate().withTags(mapOf("tagKey1", "fakeTokenPlaceholder")),
            com.azure.core.util.Context.NONE);
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
