
/**
 * Samples for RoutingIntent Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/RoutingIntentDelete.json
     */
    /**
     * Sample code: RouteTableDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void routeTableDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getRoutingIntents().delete("rg1", "virtualHub1", "Intent1",
            com.azure.core.util.Context.NONE);
    }
}
