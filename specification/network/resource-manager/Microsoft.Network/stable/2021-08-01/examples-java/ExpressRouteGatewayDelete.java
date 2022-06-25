import com.azure.core.util.Context;

/** Samples for ExpressRouteGateways Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/ExpressRouteGatewayDelete.json
     */
    /**
     * Sample code: ExpressRouteGatewayDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void expressRouteGatewayDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getExpressRouteGateways()
            .delete("resourceGroupName", "expressRouteGatewayName", Context.NONE);
    }
}
