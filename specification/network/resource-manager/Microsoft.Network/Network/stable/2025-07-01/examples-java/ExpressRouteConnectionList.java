
/**
 * Samples for ExpressRouteConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteConnectionList.json
     */
    /**
     * Sample code: ExpressRouteConnectionList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void expressRouteConnectionList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteConnections().listWithResponse("resourceGroupName",
            "expressRouteGatewayName", com.azure.core.util.Context.NONE);
    }
}
