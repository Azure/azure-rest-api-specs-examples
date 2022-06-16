import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.ExpressRoutePortInner;
import com.azure.resourcemanager.network.models.ExpressRoutePortsEncapsulation;

/** Samples for ExpressRoutePorts CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ExpressRoutePortCreate.json
     */
    /**
     * Sample code: ExpressRoutePortCreate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void expressRoutePortCreate(com.azure.resourcemanager.AzureResourceManager azure) {
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
                    .withEncapsulation(ExpressRoutePortsEncapsulation.QINQ),
                Context.NONE);
    }
}
