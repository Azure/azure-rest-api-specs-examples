
/**
 * Samples for DiscoveredAssetEndpointProfiles ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/List_DiscoveredAssetEndpointProfiles_ResourceGroup.json
     */
    /**
     * Sample code: List_DiscoveredAssetEndpointProfiles_ResourceGroup.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void listDiscoveredAssetEndpointProfilesResourceGroup(
        com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.discoveredAssetEndpointProfiles().listByResourceGroup("myResourceGroup",
            com.azure.core.util.Context.NONE);
    }
}
