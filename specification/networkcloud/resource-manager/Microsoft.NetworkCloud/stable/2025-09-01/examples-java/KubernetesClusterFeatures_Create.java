
import com.azure.resourcemanager.networkcloud.models.StringKeyValuePair;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for KubernetesClusterFeatures CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * KubernetesClusterFeatures_Create.json
     */
    /**
     * Sample code: Create or update Kubernetes cluster feature.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        createOrUpdateKubernetesClusterFeature(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.kubernetesClusterFeatures().define("featureName").withRegion("location")
            .withExistingKubernetesCluster("resourceGroupName", "kubernetesClusterName")
            .withTags(mapOf("key1", "fakeTokenPlaceholder", "key2", "fakeTokenPlaceholder"))
            .withOptions(
                Arrays.asList(new StringKeyValuePair().withKey("fakeTokenPlaceholder").withValue("featureOptionValue")))
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
