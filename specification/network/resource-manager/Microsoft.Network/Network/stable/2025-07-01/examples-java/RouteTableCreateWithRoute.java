
import com.azure.resourcemanager.network.fluent.models.RouteInner;
import com.azure.resourcemanager.network.fluent.models.RouteTableInner;
import com.azure.resourcemanager.network.models.RouteNextHopType;
import java.util.Arrays;

/**
 * Samples for RouteTables CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/RouteTableCreateWithRoute.json
     */
    /**
     * Sample code: Create route table with route.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createRouteTableWithRoute(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRouteTables().createOrUpdate("rg1", "testrt",
            new RouteTableInner().withLocation("westus")
                .withRoutes(Arrays.asList(new RouteInner().withName("route1").withAddressPrefix("10.0.3.0/24")
                    .withNextHopType(RouteNextHopType.VIRTUAL_NETWORK_GATEWAY)))
                .withDisableBgpRoutePropagation(true),
            com.azure.core.util.Context.NONE);
    }
}
