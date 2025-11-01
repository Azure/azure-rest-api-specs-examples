
/**
 * Samples for Assets List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/List_Assets_BySubscription.json
     */
    /**
     * Sample code: List_Assets_BySubscription.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        listAssetsBySubscription(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.assets().list(com.azure.core.util.Context.NONE);
    }
}
