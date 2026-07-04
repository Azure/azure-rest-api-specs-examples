
/**
 * Samples for Routes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/RouteTableRouteDelete.json
     */
    /**
     * Sample code: Delete route.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteRoute(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRoutes().delete("rg1", "testrt", "route1", com.azure.core.util.Context.NONE);
    }
}
