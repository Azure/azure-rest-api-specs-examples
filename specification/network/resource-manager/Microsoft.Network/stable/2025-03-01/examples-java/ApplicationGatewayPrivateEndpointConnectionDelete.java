
/**
 * Samples for ApplicationGatewayPrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * ApplicationGatewayPrivateEndpointConnectionDelete.json
     */
    /**
     * Sample code: Delete Application Gateway Private Endpoint Connection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        deleteApplicationGatewayPrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getApplicationGatewayPrivateEndpointConnections().delete("rg1",
            "appgw", "connection1", com.azure.core.util.Context.NONE);
    }
}
