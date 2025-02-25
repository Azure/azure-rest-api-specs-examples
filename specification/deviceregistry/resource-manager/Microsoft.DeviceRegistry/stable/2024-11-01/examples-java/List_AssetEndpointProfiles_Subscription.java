
/**
 * Samples for AssetEndpointProfiles List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/List_AssetEndpointProfiles_Subscription.json
     */
    /**
     * Sample code: List_AssetEndpointProfiles_Subscription.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        listAssetEndpointProfilesSubscription(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.assetEndpointProfiles().list(com.azure.core.util.Context.NONE);
    }
}
