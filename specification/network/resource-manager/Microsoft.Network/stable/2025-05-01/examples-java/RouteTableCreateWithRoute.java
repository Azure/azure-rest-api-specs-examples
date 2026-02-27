
import com.azure.resourcemanager.network.fluent.models.RouteInner;
import com.azure.resourcemanager.network.fluent.models.RouteTableInner;
import com.azure.resourcemanager.network.models.RouteNextHopType;
import java.util.Arrays;

/**
 * Samples for RouteTables CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/RouteTableCreateWithRoute.
     * json
     */
    /**
     * Sample code: Create route table with route.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createRouteTableWithRoute(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getRouteTables().createOrUpdate("rg1", "testrt",
            new RouteTableInner().withLocation("westus")
                .withRoutes(Arrays.asList(new RouteInner().withName("route1").withAddressPrefix("10.0.3.0/24")
                    .withNextHopType(RouteNextHopType.VIRTUAL_NETWORK_GATEWAY)))
                .withDisableBgpRoutePropagation(true),
            com.azure.core.util.Context.NONE);
    }
}
