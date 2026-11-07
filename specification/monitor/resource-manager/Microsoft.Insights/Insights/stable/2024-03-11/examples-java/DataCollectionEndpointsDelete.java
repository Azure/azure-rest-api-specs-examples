
/**
 * Samples for DataCollectionEndpoints Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-11/DataCollectionEndpointsDelete.json
     */
    /**
     * Sample code: Delete data collection endpoint.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void deleteDataCollectionEndpoint(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getDataCollectionEndpoints().deleteWithResponse("myResourceGroup",
            "myDataCollectionEndpoint", com.azure.core.util.Context.NONE);
    }
}
