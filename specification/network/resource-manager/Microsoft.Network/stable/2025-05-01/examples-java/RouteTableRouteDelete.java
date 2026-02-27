
/**
 * Samples for Routes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/RouteTableRouteDelete.json
     */
    /**
     * Sample code: Delete route.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteRoute(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getRoutes().delete("rg1", "testrt", "route1",
            com.azure.core.util.Context.NONE);
    }
}
