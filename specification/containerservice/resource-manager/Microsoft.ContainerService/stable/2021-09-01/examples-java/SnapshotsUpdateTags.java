import com.azure.core.util.Context;
import com.azure.resourcemanager.containerservice.models.TagsObject;
import java.util.HashMap;
import java.util.Map;

/** Samples for Snapshots UpdateTags. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/SnapshotsUpdateTags.json
     */
    /**
     * Sample code: Update Snapshot Tags.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateSnapshotTags(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .kubernetesClusters()
            .manager()
            .serviceClient()
            .getSnapshots()
            .updateTagsWithResponse(
                "rg1", "snapshot1", new TagsObject().withTags(mapOf("key2", "new-val2", "key3", "val3")), Context.NONE);
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
