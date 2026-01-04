
/**
 * Samples for VirtualHubIpConfiguration Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VirtualHubIpConfigurationGet.
     * json
     */
    /**
     * Sample code: VirtualHubVirtualHubRouteTableV2Get.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualHubVirtualHubRouteTableV2Get(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualHubIpConfigurations().getWithResponse("rg1", "hub1",
            "ipconfig1", com.azure.core.util.Context.NONE);
    }
}
