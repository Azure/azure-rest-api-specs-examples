
import com.azure.resourcemanager.network.fluent.models.RouteInner;
import com.azure.resourcemanager.network.models.RouteNextHopType;

/**
 * Samples for Routes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/RouteTableRouteCreate.json
     */
    /**
     * Sample code: Create route.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createRoute(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRoutes().createOrUpdate("rg1", "testrt", "route1",
            new RouteInner().withAddressPrefix("10.0.3.0/24").withNextHopType(RouteNextHopType.VIRTUAL_NETWORK_GATEWAY),
            com.azure.core.util.Context.NONE);
    }
}
