
/**
 * Samples for RouteMaps Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/RouteMapDelete.json
     */
    /**
     * Sample code: RouteMapDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void routeMapDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getRouteMaps().delete("rg1", "virtualHub1", "routeMap1",
            com.azure.core.util.Context.NONE);
    }
}
