
import com.azure.resourcemanager.network.fluent.models.ConnectionResetSharedKeyInner;

/**
 * Samples for VirtualNetworkGatewayConnections ResetSharedKey.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayConnectionResetSharedKey.json
     */
    /**
     * Sample code: ResetVirtualNetworkGatewayConnectionSharedKey.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        resetVirtualNetworkGatewayConnectionSharedKey(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGatewayConnections().resetSharedKey("rg1", "conn1",
            new ConnectionResetSharedKeyInner().withKeyLength(128), com.azure.core.util.Context.NONE);
    }
}
