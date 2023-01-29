import com.azure.resourcemanager.compute.models.SnapshotUpdate;
import java.util.HashMap;
import java.util.Map;

/** Samples for Snapshots Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-07-02/examples/snapshotExamples/Snapshot_Update.json
     */
    /**
     * Sample code: Update a snapshot.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateASnapshot(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getSnapshots()
            .update(
                "myResourceGroup",
                "mySnapshot",
                new SnapshotUpdate()
                    .withTags(mapOf("department", "Development", "project", "UpdateSnapshots"))
                    .withDiskSizeGB(20),
                com.azure.core.util.Context.NONE);
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
