
import com.azure.resourcemanager.networkcloud.models.Cluster;
import com.azure.resourcemanager.networkcloud.models.RuntimeProtectionConfigurationPatch;
import com.azure.resourcemanager.networkcloud.models.RuntimeProtectionDefinitionUpdateMode;
import com.azure.resourcemanager.networkcloud.models.RuntimeProtectionEnforcementLevel;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Clusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Clusters_Patch_RuntimeProtectionConfiguration.json
     */
    /**
     * Sample code: Patch runtime protection configuration.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        patchRuntimeProtectionConfiguration(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        Cluster resource = manager.clusters()
            .getByResourceGroupWithResponse("resourceGroupName", "clusterName", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("key1", "fakeTokenPlaceholder", "key2", "fakeTokenPlaceholder"))
            .withRuntimeProtectionConfiguration(new RuntimeProtectionConfigurationPatch()
                .withDefinitionUpdateMode(RuntimeProtectionDefinitionUpdateMode.AUTOMATIC)
                .withEnforcementLevel(RuntimeProtectionEnforcementLevel.ON_DEMAND))
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
