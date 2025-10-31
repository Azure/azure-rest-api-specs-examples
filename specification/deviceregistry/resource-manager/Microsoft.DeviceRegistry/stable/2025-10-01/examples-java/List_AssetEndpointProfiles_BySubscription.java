
/**
 * Samples for AssetEndpointProfiles List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/List_AssetEndpointProfiles_BySubscription.json
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
