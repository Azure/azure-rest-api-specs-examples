
/**
 * Samples for VirtualHubIpConfiguration Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * VirtualHubIpConfigurationDelete.json
     */
    /**
     * Sample code: VirtualHubIpConfigurationDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualHubIpConfigurationDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualHubIpConfigurations().delete("rg1", "hub1", "ipconfig1",
            com.azure.core.util.Context.NONE);
    }
}
