
/**
 * Samples for DiscoveredAssetEndpointProfiles Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/Delete_DiscoveredAssetEndpointProfile.json
     */
    /**
     * Sample code: Delete_DiscoveredAssetEndpointProfile.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        deleteDiscoveredAssetEndpointProfile(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.discoveredAssetEndpointProfiles().delete("myResourceGroup", "my-discoveredassetendpointprofile",
            com.azure.core.util.Context.NONE);
    }
}
