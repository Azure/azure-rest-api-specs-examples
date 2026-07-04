
/**
 * Samples for VirtualNetworkGatewayConnections ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayConnectionsList.json
     */
    /**
     * Sample code: ListVirtualNetworkGatewayConnectionsinResourceGroup.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listVirtualNetworkGatewayConnectionsinResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGatewayConnections().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
