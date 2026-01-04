
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.ExpressRouteCircuitInner;
import com.azure.resourcemanager.network.models.ExpressRouteCircuitSku;
import com.azure.resourcemanager.network.models.ExpressRouteCircuitSkuFamily;
import com.azure.resourcemanager.network.models.ExpressRouteCircuitSkuTier;

/**
 * Samples for ExpressRouteCircuits CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * ExpressRouteCircuitCreateOnExpressRoutePort.json
     */
    /**
     * Sample code: Create ExpressRouteCircuit on ExpressRoutePort.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createExpressRouteCircuitOnExpressRoutePort(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRouteCircuits()
            .createOrUpdate("rg1", "expressRouteCircuit1", new ExpressRouteCircuitInner().withLocation("westus")
                .withSku(new ExpressRouteCircuitSku().withName("Premium_MeteredData")
                    .withTier(ExpressRouteCircuitSkuTier.PREMIUM).withFamily(ExpressRouteCircuitSkuFamily.METERED_DATA))
                .withExpressRoutePort(new SubResource().withId(
                    "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName"))
                .withBandwidthInGbps(10.0F).withAuthorizationKey("fakeTokenPlaceholder")
                .withEnableDirectPortRateLimit(false), com.azure.core.util.Context.NONE);
    }
}
