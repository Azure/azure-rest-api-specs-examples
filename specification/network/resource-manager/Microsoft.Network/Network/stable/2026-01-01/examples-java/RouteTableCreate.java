
import com.azure.resourcemanager.network.fluent.models.RouteTableInner;

/**
 * Samples for RouteTables CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/RouteTableCreate.json
     */
    /**
     * Sample code: Create route table.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createRouteTable(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRouteTables().createOrUpdate("rg1", "testrt",
            new RouteTableInner().withLocation("westus"), com.azure.core.util.Context.NONE);
    }
}
