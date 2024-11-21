
import com.azure.resourcemanager.networkcloud.models.Cluster;
import com.azure.resourcemanager.networkcloud.models.ClusterSecretArchive;
import com.azure.resourcemanager.networkcloud.models.ClusterSecretArchiveEnabled;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Clusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2024-06-01-preview/examples/
     * Clusters_Patch_SecretArchive.json
     */
    /**
     * Sample code: Patch secret archive.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void patchSecretArchive(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        Cluster resource = manager.clusters()
            .getByResourceGroupWithResponse("resourceGroupName", "clusterName", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("key1", "fakeTokenPlaceholder", "key2", "fakeTokenPlaceholder"))
            .withSecretArchive(new ClusterSecretArchive().withKeyVaultId("fakeTokenPlaceholder")
                .withUseKeyVault(ClusterSecretArchiveEnabled.TRUE))
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
