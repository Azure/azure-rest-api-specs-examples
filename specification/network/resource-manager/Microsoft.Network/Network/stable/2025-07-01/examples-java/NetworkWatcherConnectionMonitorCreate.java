
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
     * x-ms-original-file: 2025-07-01/NetworkWatcherConnectionMonitorCreate.json
     */
    /**
     * Sample code: Create connection monitor V1.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createConnectionMonitorV1(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getConnectionMonitors().createOrUpdate("rg1", "nw1", "cm1", new ConnectionMonitorInner()
            .withLocation("eastus")
            .withEndpoints(Arrays.asList(new ConnectionMonitorEndpoint().withName("source").withResourceId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/ct1"),
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
