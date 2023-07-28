/** Samples for ConnectionMonitors Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/examples/NetworkWatcherConnectionMonitorStart.json
     */
    /**
     * Sample code: Start connection monitor.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void startConnectionMonitor(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getConnectionMonitors()
            .start("rg1", "nw1", "cm1", com.azure.core.util.Context.NONE);
    }
}
