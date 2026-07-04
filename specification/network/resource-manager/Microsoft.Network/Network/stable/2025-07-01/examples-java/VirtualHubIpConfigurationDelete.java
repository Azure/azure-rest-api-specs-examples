
/**
 * Samples for VirtualHubIpConfiguration Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualHubIpConfigurationDelete.json
     */
    /**
     * Sample code: VirtualHubIpConfigurationDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void virtualHubIpConfigurationDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualHubIpConfigurations().delete("rg1", "hub1", "ipconfig1",
            com.azure.core.util.Context.NONE);
    }
}
