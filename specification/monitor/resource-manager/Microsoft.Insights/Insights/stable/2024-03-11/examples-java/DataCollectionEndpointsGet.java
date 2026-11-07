
/**
 * Samples for DataCollectionEndpoints GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-11/DataCollectionEndpointsGet.json
     */
    /**
     * Sample code: Get data collection endpoint.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getDataCollectionEndpoint(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getDataCollectionEndpoints().getByResourceGroupWithResponse("myResourceGroup",
            "myDataCollectionEndpoint", com.azure.core.util.Context.NONE);
    }
}
