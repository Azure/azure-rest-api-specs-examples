
import com.azure.resourcemanager.network.models.FlowLogStatusParameters;

/**
 * Samples for NetworkWatchers GetFlowLogStatus.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkWatcherFlowLogStatusQuery.json
     */
    /**
     * Sample code: Get flow log status.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getFlowLogStatus(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkWatchers().getFlowLogStatus("rg1", "nw1",
            new FlowLogStatusParameters().withTargetResourceId(
                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg1"),
            com.azure.core.util.Context.NONE);
    }
}
