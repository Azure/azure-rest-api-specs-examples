import com.azure.core.util.Context;
import com.azure.resourcemanager.network.models.TopologyParameters;

/** Samples for NetworkWatchers GetTopology. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkWatcherTopologyGet.json
     */
    /**
     * Sample code: Get Topology.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getTopology(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkWatchers()
            .getTopologyWithResponse(
                "rg1", "nw1", new TopologyParameters().withTargetResourceGroupName("rg2"), Context.NONE);
    }
}
