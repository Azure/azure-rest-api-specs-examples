
/**
 * Samples for VirtualNetworkGateways ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/VirtualNetworkGatewayList.json
     */
    /**
     * Sample code: ListVirtualNetworkGatewaysinResourceGroup.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listVirtualNetworkGatewaysinResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
