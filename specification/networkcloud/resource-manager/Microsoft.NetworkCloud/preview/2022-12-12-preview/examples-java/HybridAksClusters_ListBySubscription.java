/** Samples for HybridAksClusters List. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/HybridAksClusters_ListBySubscription.json
     */
    /**
     * Sample code: List Hybrid AKS provisioned clusters data for subscription.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void listHybridAKSProvisionedClustersDataForSubscription(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.hybridAksClusters().list(com.azure.core.util.Context.NONE);
    }
}
