
import com.azure.resourcemanager.network.fluent.models.ApplicationGatewayPrivateEndpointConnectionInner;
import com.azure.resourcemanager.network.models.PrivateLinkServiceConnectionState;

/**
 * Samples for ApplicationGatewayPrivateEndpointConnections Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/
     * ApplicationGatewayPrivateEndpointConnectionUpdate.json
     */
    /**
     * Sample code: Update Application Gateway Private Endpoint Connection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        updateApplicationGatewayPrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getApplicationGatewayPrivateEndpointConnections()
            .update("rg1", "appgw", "connection1",
                new ApplicationGatewayPrivateEndpointConnectionInner().withName("connection1")
                    .withPrivateLinkServiceConnectionState(new PrivateLinkServiceConnectionState()
                        .withStatus("Approved").withDescription("approved it for some reason.")),
                com.azure.core.util.Context.NONE);
    }
}
