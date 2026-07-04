
/**
 * Samples for VirtualHubIpConfiguration Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualHubIpConfigurationGet.json
     */
    /**
     * Sample code: VirtualHubVirtualHubRouteTableV2Get.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void virtualHubVirtualHubRouteTableV2Get(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualHubIpConfigurations().getWithResponse("rg1", "hub1", "ipconfig1",
            com.azure.core.util.Context.NONE);
    }
}
