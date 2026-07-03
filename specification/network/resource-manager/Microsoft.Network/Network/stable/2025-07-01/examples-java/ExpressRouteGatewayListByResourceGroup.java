
/**
 * Samples for ExpressRouteGateways ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteGatewayListByResourceGroup.json
     */
    /**
     * Sample code: ExpressRouteGatewayListByResourceGroup.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        expressRouteGatewayListByResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteGateways().listByResourceGroupWithResponse("resourceGroupName",
            com.azure.core.util.Context.NONE);
    }
}
