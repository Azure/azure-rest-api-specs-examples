
/**
 * Samples for DataCollectionEndpoints GetNSP.
 */
public final class Main {
    /*
     * x-ms-original-file: 2021-10-01/NSPForDataCollectionEndpoints_Get.json
     */
    /**
     * Sample code: Get NSP configs by name for data collection endpoint.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void
        getNSPConfigsByNameForDataCollectionEndpoint(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getDataCollectionEndpoints().getNSPWithResponse("exampleRG",
            "someDataCollectionEndpoint", "somePerimeterConfiguration", com.azure.core.util.Context.NONE);
    }
}
