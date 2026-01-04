
import com.azure.resourcemanager.network.models.NextHopParameters;

/**
 * Samples for NetworkWatchers GetNextHop.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkWatcherNextHopGet.json
     */
    /**
     * Sample code: Get next hop.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getNextHop(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkWatchers().getNextHop("rg1", "nw1",
            new NextHopParameters()
                .withTargetResourceId(
                    "/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1")
                .withSourceIpAddress("10.0.0.5").withDestinationIpAddress("10.0.0.10").withTargetNicResourceId(
                    "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/nic1"),
            com.azure.core.util.Context.NONE);
    }
}
