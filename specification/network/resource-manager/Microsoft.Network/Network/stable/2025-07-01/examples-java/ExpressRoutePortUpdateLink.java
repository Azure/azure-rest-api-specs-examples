
import com.azure.resourcemanager.network.fluent.models.ExpressRouteLinkInner;
import com.azure.resourcemanager.network.fluent.models.ExpressRoutePortInner;
import com.azure.resourcemanager.network.models.ExpressRouteLinkAdminState;
import com.azure.resourcemanager.network.models.ExpressRoutePortsBillingType;
import com.azure.resourcemanager.network.models.ExpressRoutePortsEncapsulation;
import java.util.Arrays;

/**
 * Samples for ExpressRoutePorts CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRoutePortUpdateLink.json
     */
    /**
     * Sample code: ExpressRoutePortUpdateLink.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void expressRoutePortUpdateLink(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRoutePorts().createOrUpdate("rg1", "portName",
            new ExpressRoutePortInner().withLocation("westus").withPeeringLocation("peeringLocationName")
                .withBandwidthInGbps(100).withEncapsulation(ExpressRoutePortsEncapsulation.QINQ)
                .withLinks(Arrays.asList(
                    new ExpressRouteLinkInner().withName("link1").withAdminState(ExpressRouteLinkAdminState.ENABLED)))
                .withBillingType(ExpressRoutePortsBillingType.UNLIMITED_DATA),
            com.azure.core.util.Context.NONE);
    }
}
