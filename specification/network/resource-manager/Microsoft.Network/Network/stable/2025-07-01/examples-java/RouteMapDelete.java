
/**
 * Samples for RouteMaps Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/RouteMapDelete.json
     */
    /**
     * Sample code: RouteMapDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void routeMapDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRouteMaps().delete("rg1", "virtualHub1", "routeMap1",
            com.azure.core.util.Context.NONE);
    }
}
