
/**
 * Samples for VirtualNetworkGateways Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayDelete.json
     */
    /**
     * Sample code: DeleteVirtualNetworkGateway.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteVirtualNetworkGateway(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways().delete("rg1", "vpngw", com.azure.core.util.Context.NONE);
    }
}
