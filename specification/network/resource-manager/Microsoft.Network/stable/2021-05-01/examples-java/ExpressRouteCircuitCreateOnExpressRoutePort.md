Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.SubResource;
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.ExpressRouteCircuitInner;
import com.azure.resourcemanager.network.models.ExpressRouteCircuitSku;
import com.azure.resourcemanager.network.models.ExpressRouteCircuitSkuFamily;
import com.azure.resourcemanager.network.models.ExpressRouteCircuitSkuTier;

/** Samples for ExpressRouteCircuits CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ExpressRouteCircuitCreateOnExpressRoutePort.json
     */
    /**
     * Sample code: Create ExpressRouteCircuit on ExpressRoutePort.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createExpressRouteCircuitOnExpressRoutePort(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getExpressRouteCircuits()
            .createOrUpdate(
                "rg1",
                "expressRouteCircuit1",
                new ExpressRouteCircuitInner()
                    .withLocation("westus")
                    .withSku(
                        new ExpressRouteCircuitSku()
                            .withName("Premium_MeteredData")
                            .withTier(ExpressRouteCircuitSkuTier.PREMIUM)
                            .withFamily(ExpressRouteCircuitSkuFamily.METERED_DATA))
                    .withExpressRoutePort(
                        new SubResource()
                            .withId(
                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName"))
                    .withBandwidthInGbps(10.0f),
                Context.NONE);
    }
}
```
