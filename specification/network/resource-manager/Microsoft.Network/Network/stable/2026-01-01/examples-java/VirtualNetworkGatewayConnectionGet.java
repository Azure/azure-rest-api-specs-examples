
/**
 * Samples for VirtualNetworkGatewayConnections GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/VirtualNetworkGatewayConnectionGet.json
     */
    /**
     * Sample code: GetVirtualNetworkGatewayConnection.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getVirtualNetworkGatewayConnection(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGatewayConnections().getByResourceGroupWithResponse("rg1", "connS2S",
            com.azure.core.util.Context.NONE);
    }
}
