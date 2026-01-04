
/**
 * Samples for RouteMaps List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/RouteMapList.json
     */
    /**
     * Sample code: RouteMapList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void routeMapList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getRouteMaps().list("rg1", "virtualHub1",
            com.azure.core.util.Context.NONE);
    }
}
