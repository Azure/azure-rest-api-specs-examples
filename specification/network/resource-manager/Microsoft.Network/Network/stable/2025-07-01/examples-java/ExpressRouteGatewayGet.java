
/**
 * Samples for ExpressRouteGateways GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteGatewayGet.json
     */
    /**
     * Sample code: ExpressRouteGatewayGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void expressRouteGatewayGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteGateways().getByResourceGroupWithResponse("resourceGroupName",
            "expressRouteGatewayName", com.azure.core.util.Context.NONE);
    }
}
