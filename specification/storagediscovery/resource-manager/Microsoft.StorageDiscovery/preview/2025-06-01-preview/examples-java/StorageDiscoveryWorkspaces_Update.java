
import com.azure.resourcemanager.storagediscovery.models.StorageDiscoveryResourceType;
import com.azure.resourcemanager.storagediscovery.models.StorageDiscoveryScope;
import com.azure.resourcemanager.storagediscovery.models.StorageDiscoverySku;
import com.azure.resourcemanager.storagediscovery.models.StorageDiscoveryWorkspace;
import com.azure.resourcemanager.storagediscovery.models.StorageDiscoveryWorkspacePropertiesUpdate;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for StorageDiscoveryWorkspaces Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/StorageDiscoveryWorkspaces_Update.json
     */
    /**
     * Sample code: Update a StorageDiscoveryWorkspace.
     * 
     * @param manager Entry point to StorageDiscoveryManager.
     */
    public static void
        updateAStorageDiscoveryWorkspace(com.azure.resourcemanager.storagediscovery.StorageDiscoveryManager manager) {
        StorageDiscoveryWorkspace resource = manager.storageDiscoveryWorkspaces()
            .getByResourceGroupWithResponse("sample-rg", "Sample-Storage-Workspace", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withProperties(new StorageDiscoveryWorkspacePropertiesUpdate()
            .withSku(StorageDiscoverySku.fromString("Premium"))
            .withDescription("Updated Sample Storage Discovery Workspace")
            .withWorkspaceRoots(Arrays.asList("/subscriptions/b79cb3ba-745e-5d9a-8903-4a02327a7e09"))
            .withScopes(Arrays.asList(new StorageDiscoveryScope().withDisplayName("Updated-Sample-Collection")
                .withResourceTypes(Arrays.asList(StorageDiscoveryResourceType.fromString(
                    "/subscriptions/b79cb3ba-745e-5d9a-8903-4a02327a7e09/resourceGroups/sample-rg/providers/Microsoft.Storage/storageAccounts/updated-sample-storageAccount")))
                .withTagKeysOnly(Arrays.asList("updated-filtertag1", "updated-filtertag2"))
                .withTags(mapOf("updated-filtertag3", "updated-value3", "updated-filtertag4", "updated-value4")))))
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
