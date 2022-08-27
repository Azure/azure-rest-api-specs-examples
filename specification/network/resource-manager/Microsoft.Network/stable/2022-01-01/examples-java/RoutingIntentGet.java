import com.azure.core.util.Context;

/** Samples for RoutingIntent Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/RoutingIntentGet.json
     */
    /**
     * Sample code: RouteTableGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void routeTableGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getRoutingIntents()
            .getWithResponse("rg1", "virtualHub1", "Intent1", Context.NONE);
    }
}
