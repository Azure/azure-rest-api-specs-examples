
/**
 * Samples for RackSkus List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/RackSkus_ListBySubscription.json
     */
    /**
     * Sample code: List rack SKUs for subscription.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void listRackSKUsForSubscription(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.rackSkus().list(com.azure.core.util.Context.NONE);
    }
}
