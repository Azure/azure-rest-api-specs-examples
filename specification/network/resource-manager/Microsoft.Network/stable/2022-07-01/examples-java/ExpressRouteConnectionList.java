import com.azure.core.util.Context;

/** Samples for ExpressRouteConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/ExpressRouteConnectionList.json
     */
    /**
     * Sample code: ExpressRouteConnectionList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void expressRouteConnectionList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getExpressRouteConnections()
            .listWithResponse("resourceGroupName", "expressRouteGatewayName", Context.NONE);
    }
}
