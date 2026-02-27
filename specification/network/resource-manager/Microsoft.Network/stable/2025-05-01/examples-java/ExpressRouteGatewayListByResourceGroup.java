
/**
 * Samples for ExpressRouteGateways ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * ExpressRouteGatewayListByResourceGroup.json
     */
    /**
     * Sample code: ExpressRouteGatewayListByResourceGroup.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void expressRouteGatewayListByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRouteGateways()
            .listByResourceGroupWithResponse("resourceGroupName", com.azure.core.util.Context.NONE);
    }
}
