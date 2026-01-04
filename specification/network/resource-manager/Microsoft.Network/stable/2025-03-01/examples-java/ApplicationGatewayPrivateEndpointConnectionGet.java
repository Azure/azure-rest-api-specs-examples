
/**
 * Samples for ApplicationGatewayPrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * ApplicationGatewayPrivateEndpointConnectionGet.json
     */
    /**
     * Sample code: Get Application Gateway Private Endpoint Connection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getApplicationGatewayPrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getApplicationGatewayPrivateEndpointConnections()
            .getWithResponse("rg1", "appgw", "connection1", com.azure.core.util.Context.NONE);
    }
}
