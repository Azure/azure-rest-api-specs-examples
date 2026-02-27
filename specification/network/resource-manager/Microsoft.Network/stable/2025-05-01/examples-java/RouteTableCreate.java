
import com.azure.resourcemanager.network.fluent.models.RouteTableInner;

/**
 * Samples for RouteTables CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/RouteTableCreate.json
     */
    /**
     * Sample code: Create route table.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createRouteTable(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getRouteTables().createOrUpdate("rg1", "testrt",
            new RouteTableInner().withLocation("westus"), com.azure.core.util.Context.NONE);
    }
}
