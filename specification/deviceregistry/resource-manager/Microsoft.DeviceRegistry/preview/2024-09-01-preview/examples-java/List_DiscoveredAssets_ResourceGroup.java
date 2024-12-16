
/**
 * Samples for DiscoveredAssets ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/List_DiscoveredAssets_ResourceGroup.json
     */
    /**
     * Sample code: List_DiscoveredAssets_ResourceGroup.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        listDiscoveredAssetsResourceGroup(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.discoveredAssets().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
