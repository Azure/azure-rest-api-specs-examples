
/**
 * Samples for Volumes List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2025-07-01-preview/examples/
     * Volumes_ListBySubscription.json
     */
    /**
     * Sample code: List volume for subscription.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void listVolumeForSubscription(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.volumes().list(null, null, com.azure.core.util.Context.NONE);
    }
}
