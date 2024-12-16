
/**
 * Samples for DiscoveredAssets List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/List_DiscoveredAssets_Subscription.json
     */
    /**
     * Sample code: List_DiscoveredAssets_Subscription.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        listDiscoveredAssetsSubscription(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.discoveredAssets().list(com.azure.core.util.Context.NONE);
    }
}
