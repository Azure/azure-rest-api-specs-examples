
import com.azure.resourcemanager.compute.models.DiskAccessUpdate;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DiskAccesses Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/diskAccessExamples/DiskAccess_Update.json
     */
    /**
     * Sample code: update a disk access resource.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void updateADiskAccessResource(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDiskAccesses().update("myResourceGroup", "myDiskAccess",
            new DiskAccessUpdate().withTags(mapOf("department", "Development", "project", "PrivateEndpoints")),
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
