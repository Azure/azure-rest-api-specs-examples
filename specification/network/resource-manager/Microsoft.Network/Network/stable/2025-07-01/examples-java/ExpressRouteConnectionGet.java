
/**
 * Samples for ExpressRouteConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteConnectionGet.json
     */
    /**
     * Sample code: ExpressRouteConnectionGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void expressRouteConnectionGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteConnections().getWithResponse("resourceGroupName",
            "expressRouteGatewayName", "connectionName", com.azure.core.util.Context.NONE);
    }
}
