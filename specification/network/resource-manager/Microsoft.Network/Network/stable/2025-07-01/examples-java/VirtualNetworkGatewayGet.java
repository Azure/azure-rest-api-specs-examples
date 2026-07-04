
/**
 * Samples for VirtualNetworkGateways GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayGet.json
     */
    /**
     * Sample code: GetVirtualNetworkGateway.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getVirtualNetworkGateway(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways().getByResourceGroupWithResponse("rg1", "vpngw",
            com.azure.core.util.Context.NONE);
    }
}
