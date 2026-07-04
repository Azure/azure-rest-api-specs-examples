
/**
 * Samples for VirtualNetworks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkDelete.json
     */
    /**
     * Sample code: Delete virtual network.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteVirtualNetwork(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworks().delete("rg1", "test-vnet", com.azure.core.util.Context.NONE);
    }
}
