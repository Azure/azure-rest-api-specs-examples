
import com.azure.resourcemanager.network.fluent.models.ExpressRoutePortInner;
import com.azure.resourcemanager.network.models.ExpressRoutePortsBillingType;
import com.azure.resourcemanager.network.models.ExpressRoutePortsEncapsulation;

/**
 * Samples for ExpressRoutePorts CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRoutePortCreate.json
     */
    /**
     * Sample code: ExpressRoutePortCreate.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void expressRoutePortCreate(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRoutePorts().createOrUpdate("rg1", "portName",
            new ExpressRoutePortInner().withLocation("westus").withPeeringLocation("peeringLocationName")
                .withBandwidthInGbps(100).withEncapsulation(ExpressRoutePortsEncapsulation.QINQ)
                .withBillingType(ExpressRoutePortsBillingType.UNLIMITED_DATA),
            com.azure.core.util.Context.NONE);
    }
}
