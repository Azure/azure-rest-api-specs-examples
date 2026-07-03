
import com.azure.resourcemanager.network.fluent.models.RouteInner;
import com.azure.resourcemanager.network.models.RouteNextHopEcmp;
import com.azure.resourcemanager.network.models.RouteNextHopType;
import java.util.Arrays;

/**
 * Samples for Routes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/RouteTableRouteCreateEcmp.json
     */
    /**
     * Sample code: Create ECMP route.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createECMPRoute(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRoutes().createOrUpdate("rg1", "testrt", "ecmp-route",
            new RouteInner().withAddressPrefix("10.1.0.0/16").withNextHopType(RouteNextHopType.VIRTUAL_APPLIANCE_ECMP)
                .withNextHop(
                    new RouteNextHopEcmp().withNextHopIpAddresses(Arrays.asList("10.0.0.4", "10.0.0.5", "10.0.0.6"))),
            com.azure.core.util.Context.NONE);
    }
}
