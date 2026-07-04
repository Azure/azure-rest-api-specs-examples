
/**
 * Samples for FlowLogs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkWatcherFlowLogList.json
     */
    /**
     * Sample code: List connection monitors.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listConnectionMonitors(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFlowLogs().list("rg1", "nw1", com.azure.core.util.Context.NONE);
    }
}
