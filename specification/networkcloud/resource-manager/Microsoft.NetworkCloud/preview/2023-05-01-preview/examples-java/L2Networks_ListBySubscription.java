/** Samples for L2Networks List. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/L2Networks_ListBySubscription.json
     */
    /**
     * Sample code: List L2 networks for subscription.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void listL2NetworksForSubscription(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.l2Networks().list(com.azure.core.util.Context.NONE);
    }
}
