
/**
 * Samples for ExpressRouteConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteConnectionDelete.json
     */
    /**
     * Sample code: ExpressRouteConnectionDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void expressRouteConnectionDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteConnections().delete("resourceGroupName", "expressRouteGatewayName",
            "connectionName", com.azure.core.util.Context.NONE);
    }
}
