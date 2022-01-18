Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.ExpressRouteLinkInner;
import com.azure.resourcemanager.network.fluent.models.ExpressRoutePortInner;
import com.azure.resourcemanager.network.models.ExpressRouteLinkAdminState;
import com.azure.resourcemanager.network.models.ExpressRoutePortsEncapsulation;
import java.util.Arrays;

/** Samples for ExpressRoutePorts CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ExpressRoutePortUpdateLink.json
     */
    /**
     * Sample code: ExpressRoutePortUpdateLink.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void expressRoutePortUpdateLink(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getExpressRoutePorts()
            .createOrUpdate(
                "rg1",
                "portName",
                new ExpressRoutePortInner()
                    .withLocation("westus")
                    .withPeeringLocation("peeringLocationName")
                    .withBandwidthInGbps(100)
                    .withEncapsulation(ExpressRoutePortsEncapsulation.QINQ)
                    .withLinks(
                        Arrays
                            .asList(
                                new ExpressRouteLinkInner()
                                    .withName("link1")
                                    .withAdminState(ExpressRouteLinkAdminState.ENABLED))),
                Context.NONE);
    }
}
```
