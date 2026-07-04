
/**
 * Samples for Routes List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/RouteTableRouteList.json
     */
    /**
     * Sample code: List routes.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listRoutes(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRoutes().list("rg1", "testrt", com.azure.core.util.Context.NONE);
    }
}
