
/**
 * Samples for Routes Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/RouteTableRouteGet.json
     */
    /**
     * Sample code: Get route.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getRoute(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRoutes().getWithResponse("rg1", "testrt", "route1",
            com.azure.core.util.Context.NONE);
    }
}
