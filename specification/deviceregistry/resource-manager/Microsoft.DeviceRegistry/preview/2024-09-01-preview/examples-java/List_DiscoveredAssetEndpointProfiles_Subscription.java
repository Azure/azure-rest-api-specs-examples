
/**
 * Samples for DiscoveredAssetEndpointProfiles List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/List_DiscoveredAssetEndpointProfiles_Subscription.json
     */
    /**
     * Sample code: List_DiscoveredAssetEndpointProfiles_Subscription.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void listDiscoveredAssetEndpointProfilesSubscription(
        com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.discoveredAssetEndpointProfiles().list(com.azure.core.util.Context.NONE);
    }
}
