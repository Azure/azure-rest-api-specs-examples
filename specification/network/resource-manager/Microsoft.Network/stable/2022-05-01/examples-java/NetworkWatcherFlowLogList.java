import com.azure.core.util.Context;

/** Samples for FlowLogs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NetworkWatcherFlowLogList.json
     */
    /**
     * Sample code: List connection monitors.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listConnectionMonitors(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getFlowLogs().list("rg1", "nw1", Context.NONE);
    }
}
