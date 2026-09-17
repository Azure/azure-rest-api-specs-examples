
/**
 * Samples for VirtualNetworks GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/VirtualNetworkGet.json
     */
    /**
     * Sample code: Get virtual network.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getVirtualNetwork(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworks().getByResourceGroupWithResponse("rg1", "test-vnet", null,
            com.azure.core.util.Context.NONE);
    }
}
