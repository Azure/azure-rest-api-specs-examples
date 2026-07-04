
import com.azure.resourcemanager.network.fluent.models.RouteTableInner;
import com.azure.resourcemanager.network.models.DisablePeeringRoute;

/**
 * Samples for RouteTables CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/RouteTableCreateWithDisablePeeringRoute.json
     */
    /**
     * Sample code: Create route table with disable peering route.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        createRouteTableWithDisablePeeringRoute(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRouteTables()
            .createOrUpdate("rg1", "testrt", new RouteTableInner().withLocation("westus")
                .withDisableBgpRoutePropagation(true).withDisablePeeringRoute(DisablePeeringRoute.ALL),
                com.azure.core.util.Context.NONE);
    }
}
