
import com.azure.resourcemanager.network.fluent.models.ConnectionMonitorInner;
import com.azure.resourcemanager.network.models.ConnectionMonitorEndpoint;
import com.azure.resourcemanager.network.models.ConnectionMonitorTcpConfiguration;
import com.azure.resourcemanager.network.models.ConnectionMonitorTestConfiguration;
import com.azure.resourcemanager.network.models.ConnectionMonitorTestConfigurationProtocol;
import com.azure.resourcemanager.network.models.ConnectionMonitorTestGroup;
import java.util.Arrays;

/**
 * Samples for ConnectionMonitors CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkWatcherConnectionMonitorCreate.json
     */
    /**
     * Sample code: Create connection monitor V1.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createConnectionMonitorV1(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getConnectionMonitors().createOrUpdate("rg1", "nw1", "cm1",
            new ConnectionMonitorInner().withLocation("eastus")
                .withEndpoints(Arrays.asList(
                    new ConnectionMonitorEndpoint().withName("source").withResourceId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/ct1"),
                    new ConnectionMonitorEndpoint().withName("destination").withAddress("bing.com")))
                .withTestConfigurations(Arrays.asList(new ConnectionMonitorTestConfiguration().withName("tcp")
                    .withTestFrequencySec(60).withProtocol(ConnectionMonitorTestConfigurationProtocol.TCP)
                    .withTcpConfiguration(new ConnectionMonitorTcpConfiguration().withPort(80))))
                .withTestGroups(Arrays
                    .asList(new ConnectionMonitorTestGroup().withName("tg").withTestConfigurations(Arrays.asList("tcp"))
                        .withSources(Arrays.asList("source")).withDestinations(Arrays.asList("destination")))),
            null, com.azure.core.util.Context.NONE);
    }
}
