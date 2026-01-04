
/**
 * Samples for ExpressRouteConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/ExpressRouteConnectionGet.
     * json
     */
    /**
     * Sample code: ExpressRouteConnectionGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void expressRouteConnectionGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRouteConnections().getWithResponse("resourceGroupName",
            "expressRouteGatewayName", "connectionName", com.azure.core.util.Context.NONE);
    }
}
