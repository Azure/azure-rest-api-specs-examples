
import com.azure.resourcemanager.network.fluent.models.ConnectionMonitorInner;
import com.azure.resourcemanager.network.models.ConnectionMonitorEndpoint;
import com.azure.resourcemanager.network.models.ConnectionMonitorEndpointFilter;
import com.azure.resourcemanager.network.models.ConnectionMonitorEndpointFilterItem;
import com.azure.resourcemanager.network.models.ConnectionMonitorEndpointFilterItemType;
import com.azure.resourcemanager.network.models.ConnectionMonitorEndpointFilterType;
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
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/
     * NetworkWatcherConnectionMonitorV2Create.json
     */
    /**
     * Sample code: Create connection monitor V2.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createConnectionMonitorV2(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getConnectionMonitors().createOrUpdate("rg1", "nw1", "cm1",
            new ConnectionMonitorInner().withEndpoints(Arrays.asList(
                new ConnectionMonitorEndpoint().withName("vm1").withResourceId(
                    "/subscriptions/96e68903-0a56-4819-9987-8d08ad6a1f99/resourceGroups/NwRgIrinaCentralUSEUAP/providers/Microsoft.Compute/virtualMachines/vm1"),
                new ConnectionMonitorEndpoint().withName("CanaryWorkspaceVamshi").withResourceId(
                    "/subscriptions/96e68903-0a56-4819-9987-8d08ad6a1f99/resourceGroups/vasamudrRG/providers/Microsoft.OperationalInsights/workspaces/vasamudrWorkspace")
                    .withFilter(new ConnectionMonitorEndpointFilter()
                        .withType(ConnectionMonitorEndpointFilterType.INCLUDE)
                        .withItems(Arrays.asList(new ConnectionMonitorEndpointFilterItem()
                            .withType(ConnectionMonitorEndpointFilterItemType.AGENT_ADDRESS).withAddress("npmuser")))),
                new ConnectionMonitorEndpoint().withName("bing").withAddress("bing.com"),
                new ConnectionMonitorEndpoint().withName("google").withAddress("google.com")))
                .withTestConfigurations(Arrays
                    .asList(new ConnectionMonitorTestConfiguration().withName("testConfig1").withTestFrequencySec(60)
                        .withProtocol(ConnectionMonitorTestConfigurationProtocol.TCP).withTcpConfiguration(
                            new ConnectionMonitorTcpConfiguration().withPort(80).withDisableTraceRoute(false))))
                .withTestGroups(Arrays.asList(new ConnectionMonitorTestGroup().withName("test1").withDisable(false)
                    .withTestConfigurations(Arrays.asList("testConfig1"))
                    .withSources(Arrays.asList("vm1", "CanaryWorkspaceVamshi"))
                    .withDestinations(Arrays.asList("bing", "google"))))
                .withOutputs(Arrays.asList()),
            null, com.azure.core.util.Context.NONE);
    }
}
