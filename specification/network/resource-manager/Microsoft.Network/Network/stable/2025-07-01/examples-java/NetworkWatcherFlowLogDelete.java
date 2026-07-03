
/**
 * Samples for FlowLogs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkWatcherFlowLogDelete.json
     */
    /**
     * Sample code: Delete flow log.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteFlowLog(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFlowLogs().delete("rg1", "nw1", "fl", com.azure.core.util.Context.NONE);
    }
}
