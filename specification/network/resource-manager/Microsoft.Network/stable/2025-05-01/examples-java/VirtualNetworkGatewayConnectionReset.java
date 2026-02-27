
/**
 * Samples for VirtualNetworkGatewayConnections ResetConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * VirtualNetworkGatewayConnectionReset.json
     */
    /**
     * Sample code: ResetVirtualNetworkGatewayConnection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void resetVirtualNetworkGatewayConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGatewayConnections().resetConnection("rg1", "conn1",
            com.azure.core.util.Context.NONE);
    }
}
