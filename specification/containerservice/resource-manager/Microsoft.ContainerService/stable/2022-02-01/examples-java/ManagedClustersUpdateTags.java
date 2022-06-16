import com.azure.core.util.Context;
import com.azure.resourcemanager.containerservice.models.TagsObject;
import java.util.HashMap;
import java.util.Map;

/** Samples for ManagedClusters UpdateTags. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-02-01/examples/ManagedClustersUpdateTags.json
     */
    /**
     * Sample code: Update Managed Cluster Tags.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateManagedClusterTags(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .kubernetesClusters()
            .manager()
            .serviceClient()
            .getManagedClusters()
            .updateTags(
                "rg1", "clustername1", new TagsObject().withTags(mapOf("archv3", "", "tier", "testing")), Context.NONE);
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
