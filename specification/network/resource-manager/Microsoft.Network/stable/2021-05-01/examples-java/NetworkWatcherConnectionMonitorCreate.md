Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.ConnectionMonitorInner;
import com.azure.resourcemanager.network.models.ConnectionMonitorEndpoint;
import com.azure.resourcemanager.network.models.ConnectionMonitorTcpConfiguration;
import com.azure.resourcemanager.network.models.ConnectionMonitorTestConfiguration;
import com.azure.resourcemanager.network.models.ConnectionMonitorTestConfigurationProtocol;
import com.azure.resourcemanager.network.models.ConnectionMonitorTestGroup;
import java.util.Arrays;

/** Samples for ConnectionMonitors CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkWatcherConnectionMonitorCreate.json
     */
    /**
     * Sample code: Create connection monitor V1.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createConnectionMonitorV1(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getConnectionMonitors()
            .createOrUpdate(
                "rg1",
                "nw1",
                "cm1",
                new ConnectionMonitorInner()
                    .withLocation("eastus")
                    .withEndpoints(
                        Arrays
                            .asList(
                                new ConnectionMonitorEndpoint()
                                    .withName("source")
                                    .withResourceId(
                                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/ct1"),
                                new ConnectionMonitorEndpoint().withName("destination").withAddress("bing.com")))
                    .withTestConfigurations(
                        Arrays
                            .asList(
                                new ConnectionMonitorTestConfiguration()
                                    .withName("tcp")
                                    .withTestFrequencySec(60)
                                    .withProtocol(ConnectionMonitorTestConfigurationProtocol.TCP)
                                    .withTcpConfiguration(new ConnectionMonitorTcpConfiguration().withPort(80))))
                    .withTestGroups(
                        Arrays
                            .asList(
                                new ConnectionMonitorTestGroup()
                                    .withName("tg")
                                    .withTestConfigurations(Arrays.asList("tcp"))
                                    .withSources(Arrays.asList("source"))
                                    .withDestinations(Arrays.asList("destination")))),
                null,
                Context.NONE);
    }
}
```
