
import com.azure.resourcemanager.network.fluent.models.ExpressRouteCircuitInner;
import com.azure.resourcemanager.network.models.ExpressRouteCircuitServiceProviderProperties;
import com.azure.resourcemanager.network.models.ExpressRouteCircuitSku;
import com.azure.resourcemanager.network.models.ExpressRouteCircuitSkuFamily;
import com.azure.resourcemanager.network.models.ExpressRouteCircuitSkuTier;

/**
 * Samples for ExpressRouteCircuits CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ExpressRouteMultiCloudCircuitCreateWithPartnerAccountId.json
     */
    /**
     * Sample code: Create MultiCloud ExpressRouteCircuit with PartnerAccountId.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createMultiCloudExpressRouteCircuitWithPartnerAccountId(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuits().createOrUpdate("rg1", "circuitName",
            new ExpressRouteCircuitInner().withLocation("eastus2euap")
                .withSku(new ExpressRouteCircuitSku().withName("MultiCloud_MeteredData")
                    .withTier(ExpressRouteCircuitSkuTier.MULTI_CLOUD)
                    .withFamily(ExpressRouteCircuitSkuFamily.METERED_DATA))
                .withServiceProviderProperties(new ExpressRouteCircuitServiceProviderProperties()
                    .withServiceProviderName("AWS").withPeeringLocation("uswest2").withBandwidthInMbps(200))
                .withPartnerAccountId("123456789"),
            com.azure.core.util.Context.NONE);
    }
}
