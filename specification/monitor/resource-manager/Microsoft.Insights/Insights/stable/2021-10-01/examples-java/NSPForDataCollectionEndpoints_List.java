
/**
 * Samples for DataCollectionEndpoints ListNSP.
 */
public final class Main {
    /*
     * x-ms-original-file: 2021-10-01/NSPForDataCollectionEndpoints_List.json
     */
    /**
     * Sample code: List NSP configs by data collection endpoint.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void
        listNSPConfigsByDataCollectionEndpoint(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getDataCollectionEndpoints().listNSP("exampleRG", "someDataCollectionEndpoint",
            com.azure.core.util.Context.NONE);
    }
}
