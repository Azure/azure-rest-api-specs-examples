
import com.azure.resourcemanager.networkcloud.models.ClusterManager;
import com.azure.resourcemanager.networkcloud.models.ManagedServiceIdentity;
import com.azure.resourcemanager.networkcloud.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.networkcloud.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ClusterManagers Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/
     * ClusterManagers_Patch.json
     */
    /**
     * Sample code: Patch cluster manager.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void patchClusterManager(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        ClusterManager resource = manager.clusterManagers()
            .getByResourceGroupWithResponse("resourceGroupName", "clusterManagerName", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("key1", "fakeTokenPlaceholder", "key2", "fakeTokenPlaceholder"))
            .withIdentity(new ManagedServiceIdentity().withType(ManagedServiceIdentityType.USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userIdentity1",
                    null,
                    "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userIdentity2",
                    new UserAssignedIdentity())))
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
