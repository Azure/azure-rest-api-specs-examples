
/**
 * Samples for Volumes List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Volumes_ListBySubscription.json
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
