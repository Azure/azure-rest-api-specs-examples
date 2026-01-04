
import com.azure.resourcemanager.network.fluent.models.ConnectionMonitorInner;
import com.azure.resourcemanager.network.models.ConnectionMonitorEndpoint;
import com.azure.resourcemanager.network.models.ConnectionMonitorEndpointLocationDetails;
import com.azure.resourcemanager.network.models.ConnectionMonitorEndpointScope;
import com.azure.resourcemanager.network.models.ConnectionMonitorEndpointScopeItem;
import com.azure.resourcemanager.network.models.ConnectionMonitorTcpConfiguration;
import com.azure.resourcemanager.network.models.ConnectionMonitorTestConfiguration;
import com.azure.resourcemanager.network.models.ConnectionMonitorTestConfigurationProtocol;
import com.azure.resourcemanager.network.models.ConnectionMonitorTestGroup;
import com.azure.resourcemanager.network.models.EndpointType;
import java.util.Arrays;

/**
 * Samples for ConnectionMonitors CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * NetworkWatcherConnectionMonitorCreateWithArcNetwork.json
     */
    /**
     * Sample code: Create connection monitor with Arc Network.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createConnectionMonitorWithArcNetwork(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getConnectionMonitors().createOrUpdate("rg1", "nw1", "cm1",
            new ConnectionMonitorInner().withEndpoints(Arrays.asList(
                new ConnectionMonitorEndpoint().withName("vm1").withType(EndpointType.AZURE_VM).withResourceId(
                    "/subscriptions/9cece3e3-0f7d-47ca-af0e-9772773f90b7/resourceGroups/testRG/providers/Microsoft.Compute/virtualMachines/TESTVM"),
                new ConnectionMonitorEndpoint().withName("bing").withType(EndpointType.EXTERNAL_ADDRESS)
                    .withAddress("bing.com"),
                new ConnectionMonitorEndpoint().withName("google").withType(EndpointType.EXTERNAL_ADDRESS)
                    .withAddress("google.com"),
                new ConnectionMonitorEndpoint().withName("ArcBasedNetwork").withType(EndpointType.AZURE_ARC_NETWORK)
                    .withScope(new ConnectionMonitorEndpointScope().withInclude(
                        Arrays.asList(new ConnectionMonitorEndpointScopeItem().withAddress("172.21.128.0/20"))))
                    .withLocationDetails(new ConnectionMonitorEndpointLocationDetails().withRegion("eastus"))
                    .withSubscriptionId("9cece3e3-0f7d-47ca-af0e-9772773f90b7")))
                .withTestConfigurations(Arrays
                    .asList(new ConnectionMonitorTestConfiguration().withName("testConfig1").withTestFrequencySec(60)
                        .withProtocol(ConnectionMonitorTestConfigurationProtocol.TCP).withTcpConfiguration(
                            new ConnectionMonitorTcpConfiguration().withPort(80).withDisableTraceRoute(false))))
                .withTestGroups(Arrays.asList(new ConnectionMonitorTestGroup().withName("test1").withDisable(false)
                    .withTestConfigurations(Arrays.asList("testConfig1"))
                    .withSources(Arrays.asList("vm1", "ArcBasedNetwork"))
                    .withDestinations(Arrays.asList("bing", "google"))))
                .withOutputs(Arrays.asList()),
            null, com.azure.core.util.Context.NONE);
    }
}
