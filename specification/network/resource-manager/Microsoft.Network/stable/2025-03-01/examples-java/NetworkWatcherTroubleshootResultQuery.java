
import com.azure.resourcemanager.network.models.QueryTroubleshootingParameters;

/**
 * Samples for NetworkWatchers GetTroubleshootingResult.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * NetworkWatcherTroubleshootResultQuery.json
     */
    /**
     * Sample code: Get troubleshoot result.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getTroubleshootResult(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkWatchers().getTroubleshootingResult("rg1", "nw1",
            new QueryTroubleshootingParameters().withTargetResourceId(
                "/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1"),
            com.azure.core.util.Context.NONE);
    }
}
