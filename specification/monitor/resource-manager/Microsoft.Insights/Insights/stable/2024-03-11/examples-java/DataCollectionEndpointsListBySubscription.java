
/**
 * Samples for DataCollectionEndpoints List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-11/DataCollectionEndpointsListBySubscription.json
     */
    /**
     * Sample code: List data collection endpoints by subscription.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void
        listDataCollectionEndpointsBySubscription(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getDataCollectionEndpoints().list(com.azure.core.util.Context.NONE);
    }
}
