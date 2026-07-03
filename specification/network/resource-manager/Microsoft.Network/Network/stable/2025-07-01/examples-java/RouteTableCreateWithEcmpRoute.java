
import com.azure.resourcemanager.network.fluent.models.RouteInner;
import com.azure.resourcemanager.network.fluent.models.RouteTableInner;
import com.azure.resourcemanager.network.models.RouteNextHopEcmp;
import com.azure.resourcemanager.network.models.RouteNextHopType;
import java.util.Arrays;

/**
 * Samples for RouteTables CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/RouteTableCreateWithEcmpRoute.json
     */
    /**
     * Sample code: Create route table with ECMP route.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createRouteTableWithECMPRoute(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRouteTables().createOrUpdate("rg1", "testrt-ecmp", new RouteTableInner()
            .withLocation("westus")
            .withRoutes(Arrays.asList(new RouteInner().withName("ecmp-route").withAddressPrefix("10.1.0.0/16")
                .withNextHopType(RouteNextHopType.VIRTUAL_APPLIANCE_ECMP).withNextHop(
                    new RouteNextHopEcmp().withNextHopIpAddresses(Arrays.asList("10.0.0.4", "10.0.0.5", "10.0.0.6")))))
            .withDisableBgpRoutePropagation(false), com.azure.core.util.Context.NONE);
    }
}
