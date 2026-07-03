
/**
 * Samples for VirtualNetworks List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkListAll.json
     */
    /**
     * Sample code: List all virtual networks.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllVirtualNetworks(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworks().list(com.azure.core.util.Context.NONE);
    }
}
