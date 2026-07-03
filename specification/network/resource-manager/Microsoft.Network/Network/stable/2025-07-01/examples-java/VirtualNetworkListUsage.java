
/**
 * Samples for VirtualNetworks ListUsage.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkListUsage.json
     */
    /**
     * Sample code: VnetGetUsage.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void vnetGetUsage(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworks().listUsage("rg1", "vnetName", com.azure.core.util.Context.NONE);
    }
}
