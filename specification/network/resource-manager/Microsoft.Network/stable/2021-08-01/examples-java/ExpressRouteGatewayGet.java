import com.azure.core.util.Context;

/** Samples for ExpressRouteGateways GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/ExpressRouteGatewayGet.json
     */
    /**
     * Sample code: ExpressRouteGatewayGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void expressRouteGatewayGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getExpressRouteGateways()
            .getByResourceGroupWithResponse("resourceGroupName", "expressRouteGatewayName", Context.NONE);
    }
}
