
/**
 * Samples for PublicCloudConnectors List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/PublicCloudConnectors_ListBySubscription.json
     */
    /**
     * Sample code: PublicCloudConnectors_ListBySubscription.
     * 
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void publicCloudConnectorsListBySubscription(
        com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager.publicCloudConnectors().list(com.azure.core.util.Context.NONE);
    }
}
