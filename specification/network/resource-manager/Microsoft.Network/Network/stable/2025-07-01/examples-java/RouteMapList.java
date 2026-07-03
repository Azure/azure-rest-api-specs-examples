
/**
 * Samples for RouteMaps List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/RouteMapList.json
     */
    /**
     * Sample code: RouteMapList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void routeMapList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRouteMaps().list("rg1", "virtualHub1", com.azure.core.util.Context.NONE);
    }
}
