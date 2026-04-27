
/**
 * Samples for AssetEndpointProfiles List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/List_AssetEndpointProfiles_BySubscription.json
     */
    /**
     * Sample code: List_AssetEndpointProfiles_BySubscription.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void listAssetEndpointProfilesBySubscription(
        com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.assetEndpointProfiles().list(com.azure.core.util.Context.NONE);
    }
}
