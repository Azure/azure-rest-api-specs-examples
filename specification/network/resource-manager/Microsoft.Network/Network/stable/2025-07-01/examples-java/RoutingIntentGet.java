
/**
 * Samples for RoutingIntent Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/RoutingIntentGet.json
     */
    /**
     * Sample code: RouteTableGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void routeTableGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRoutingIntents().getWithResponse("rg1", "virtualHub1", "Intent1",
            com.azure.core.util.Context.NONE);
    }
}
