
import com.azure.resourcemanager.network.fluent.models.ConnectionSharedKeyInner;

/**
 * Samples for VirtualNetworkGatewayConnections SetSharedKey.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayConnectionSetSharedKey.json
     */
    /**
     * Sample code: SetVirtualNetworkGatewayConnectionSharedKey.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        setVirtualNetworkGatewayConnectionSharedKey(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGatewayConnections().setSharedKey("rg1", "connS2S",
            new ConnectionSharedKeyInner().withValue("AzureAbc123"), com.azure.core.util.Context.NONE);
    }
}
