
/**
 * Samples for L2Networks List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/L2Networks_ListBySubscription.json
     */
    /**
     * Sample code: List L2 networks for subscription.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listL2NetworksForSubscription(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.l2Networks().list(null, null, com.azure.core.util.Context.NONE);
    }
}
