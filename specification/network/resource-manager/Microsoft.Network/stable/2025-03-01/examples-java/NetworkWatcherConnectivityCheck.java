
import com.azure.resourcemanager.network.models.ConnectivityDestination;
import com.azure.resourcemanager.network.models.ConnectivityParameters;
import com.azure.resourcemanager.network.models.ConnectivitySource;
import com.azure.resourcemanager.network.models.IpVersion;

/**
 * Samples for NetworkWatchers CheckConnectivity.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * NetworkWatcherConnectivityCheck.json
     */
    /**
     * Sample code: Check connectivity.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void checkConnectivity(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkWatchers().checkConnectivity("rg1", "nw1",
            new ConnectivityParameters()
                .withSource(new ConnectivitySource().withResourceId(
                    "/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1"))
                .withDestination(new ConnectivityDestination().withAddress("192.168.100.4").withPort(3389))
                .withPreferredIpVersion(IpVersion.IPV4),
            com.azure.core.util.Context.NONE);
    }
}
