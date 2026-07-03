
/**
 * Samples for RoutingIntent Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/RoutingIntentDelete.json
     */
    /**
     * Sample code: RouteTableDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void routeTableDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRoutingIntents().delete("rg1", "virtualHub1", "Intent1",
            com.azure.core.util.Context.NONE);
    }
}
