
import com.azure.resourcemanager.network.fluent.models.ConnectionSharedKeyInner;

/**
 * Samples for VirtualNetworkGatewayConnections SetSharedKey.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * VirtualNetworkGatewayConnectionSetSharedKey.json
     */
    /**
     * Sample code: SetVirtualNetworkGatewayConnectionSharedKey.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        setVirtualNetworkGatewayConnectionSharedKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGatewayConnections().setSharedKey("rg1", "connS2S",
            new ConnectionSharedKeyInner().withValue("AzureAbc123"), com.azure.core.util.Context.NONE);
    }
}
