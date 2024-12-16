
/**
 * Samples for DiscoveredAssets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/Delete_DiscoveredAsset.json
     */
    /**
     * Sample code: Delete_DiscoveredAsset.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void deleteDiscoveredAsset(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.discoveredAssets().delete("myResourceGroup", "my-discoveredasset", com.azure.core.util.Context.NONE);
    }
}
