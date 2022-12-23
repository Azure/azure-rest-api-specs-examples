import com.azure.core.util.Context;

/** Samples for ExpressRouteConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/ExpressRouteConnectionDelete.json
     */
    /**
     * Sample code: ExpressRouteConnectionDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void expressRouteConnectionDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getExpressRouteConnections()
            .delete("resourceGroupName", "expressRouteGatewayName", "connectionName", Context.NONE);
    }
}
