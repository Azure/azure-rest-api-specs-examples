
import com.azure.resourcemanager.servicefabricmanagedclusters.models.ApplicationResource;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Applications Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/ApplicationPatchOperation_example.json
     */
    /**
     * Sample code: Patch an application.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void patchAnApplication(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        ApplicationResource resource = manager.applications()
            .getWithResponse("resRg", "myCluster", "myApp", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("a", "b")).apply();
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
