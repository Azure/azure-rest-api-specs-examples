
/**
 * Samples for Assets List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/List_Assets_Subscription.json
     */
    /**
     * Sample code: List_Assets_Subscription.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void listAssetsSubscription(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.assets().list(com.azure.core.util.Context.NONE);
    }
}
