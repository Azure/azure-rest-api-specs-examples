
/**
 * Samples for AccessBridges List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/AccessBridges_ListBySubscription.json
     */
    /**
     * Sample code: List access bridges for subscription.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listAccessBridgesForSubscription(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.accessBridges().list(null, null, com.azure.core.util.Context.NONE);
    }
}
