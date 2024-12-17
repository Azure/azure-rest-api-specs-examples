
/**
 * Samples for DiscoveredAssets GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/Get_DiscoveredAsset.json
     */
    /**
     * Sample code: Get_DiscoveredAsset.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void getDiscoveredAsset(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.discoveredAssets().getByResourceGroupWithResponse("myResourceGroup", "my-discoveredasset",
            com.azure.core.util.Context.NONE);
    }
}
