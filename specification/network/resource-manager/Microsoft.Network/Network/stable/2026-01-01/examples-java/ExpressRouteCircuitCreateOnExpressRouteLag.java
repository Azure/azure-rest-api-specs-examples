
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
     * x-ms-original-file: 2026-01-01/ExpressRouteCircuitCreateOnExpressRouteLag.json
     */
    /**
     * Sample code: Create ExpressRouteCircuit on ExpressRouteLag.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        createExpressRouteCircuitOnExpressRouteLag(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuits().createOrUpdate("rg1", "expressRouteCircuit1",
            new ExpressRouteCircuitInner().withLocation("eastus2euap")
                .withSku(new ExpressRouteCircuitSku().withName("Premium_MeteredData")
                    .withTier(ExpressRouteCircuitSkuTier.PREMIUM).withFamily(ExpressRouteCircuitSkuFamily.METERED_DATA))
                .withExpressRouteLag(new SubResource().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName"))
                .withBandwidthInGbps(5.0F).withEnableDirectPortRateLimit(true),
            com.azure.core.util.Context.NONE);
    }
}
