
/**
 * Samples for DataCollectionEndpoints ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-11/DataCollectionEndpointsListByResourceGroup.json
     */
    /**
     * Sample code: List data collection endpoints by resource group.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void
        listDataCollectionEndpointsByResourceGroup(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getDataCollectionEndpoints().listByResourceGroup("myResourceGroup",
            com.azure.core.util.Context.NONE);
    }
}
