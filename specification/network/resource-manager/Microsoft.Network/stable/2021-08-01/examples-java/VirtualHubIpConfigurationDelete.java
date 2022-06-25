import com.azure.core.util.Context;

/** Samples for VirtualHubIpConfiguration Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VirtualHubIpConfigurationDelete.json
     */
    /**
     * Sample code: VirtualHubIpConfigurationDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualHubIpConfigurationDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualHubIpConfigurations()
            .delete("rg1", "hub1", "ipconfig1", Context.NONE);
    }
}
