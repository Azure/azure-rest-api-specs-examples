
/**
 * Samples for RouteMaps Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/RouteMapGet.json
     */
    /**
     * Sample code: RouteMapGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void routeMapGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getRouteMaps().getWithResponse("rg1", "virtualHub1", "routeMap1",
            com.azure.core.util.Context.NONE);
    }
}
