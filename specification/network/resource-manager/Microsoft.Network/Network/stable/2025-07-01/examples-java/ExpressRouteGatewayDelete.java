
/**
 * Samples for ExpressRouteGateways Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteGatewayDelete.json
     */
    /**
     * Sample code: ExpressRouteGatewayDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void expressRouteGatewayDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteGateways().delete("resourceGroupName", "expressRouteGatewayName",
            com.azure.core.util.Context.NONE);
    }
}
