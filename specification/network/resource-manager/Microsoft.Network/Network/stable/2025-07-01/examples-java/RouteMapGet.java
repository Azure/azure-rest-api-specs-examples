
/**
 * Samples for RouteMaps Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/RouteMapGet.json
     */
    /**
     * Sample code: RouteMapGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void routeMapGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRouteMaps().getWithResponse("rg1", "virtualHub1", "routeMap1",
            com.azure.core.util.Context.NONE);
    }
}
