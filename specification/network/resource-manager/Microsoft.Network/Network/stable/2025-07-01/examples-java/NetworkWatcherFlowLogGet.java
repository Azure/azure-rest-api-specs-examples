
/**
 * Samples for FlowLogs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkWatcherFlowLogGet.json
     */
    /**
     * Sample code: Get flow log.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getFlowLog(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFlowLogs().getWithResponse("rg1", "nw1", "flowLog1",
            com.azure.core.util.Context.NONE);
    }
}
