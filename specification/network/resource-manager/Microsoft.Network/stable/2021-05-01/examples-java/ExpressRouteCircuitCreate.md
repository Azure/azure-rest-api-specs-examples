Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.ExpressRouteCircuitInner;
import com.azure.resourcemanager.network.models.ExpressRouteCircuitServiceProviderProperties;
import com.azure.resourcemanager.network.models.ExpressRouteCircuitSku;
import com.azure.resourcemanager.network.models.ExpressRouteCircuitSkuFamily;
import com.azure.resourcemanager.network.models.ExpressRouteCircuitSkuTier;
import java.util.Arrays;

/** Samples for ExpressRouteCircuits CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ExpressRouteCircuitCreate.json
     */
    /**
     * Sample code: Create ExpressRouteCircuit.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createExpressRouteCircuit(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getExpressRouteCircuits()
            .createOrUpdate(
                "rg1",
                "circuitName",
                new ExpressRouteCircuitInner()
                    .withLocation("Brazil South")
                    .withSku(
                        new ExpressRouteCircuitSku()
                            .withName("Standard_MeteredData")
                            .withTier(ExpressRouteCircuitSkuTier.STANDARD)
                            .withFamily(ExpressRouteCircuitSkuFamily.METERED_DATA))
                    .withAllowClassicOperations(false)
                    .withAuthorizations(Arrays.asList())
                    .withPeerings(Arrays.asList())
                    .withServiceProviderProperties(
                        new ExpressRouteCircuitServiceProviderProperties()
                            .withServiceProviderName("Equinix")
                            .withPeeringLocation("Silicon Valley")
                            .withBandwidthInMbps(200)),
                Context.NONE);
    }
}
```
