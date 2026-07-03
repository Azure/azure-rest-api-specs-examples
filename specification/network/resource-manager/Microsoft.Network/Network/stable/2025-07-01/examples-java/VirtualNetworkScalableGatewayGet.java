
/**
 * Samples for VirtualNetworkGateways GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkScalableGatewayGet.json
     */
    /**
     * Sample code: GetVirtualNetworkScalableGateway.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getVirtualNetworkScalableGateway(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways().getByResourceGroupWithResponse("rg1", "ergw",
            com.azure.core.util.Context.NONE);
    }
}
