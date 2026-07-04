
/**
 * Samples for VirtualHubBgpConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualHubBgpConnectionGet.json
     */
    /**
     * Sample code: VirtualHubVirtualHubRouteTableV2Get.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void virtualHubVirtualHubRouteTableV2Get(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualHubBgpConnections().getWithResponse("rg1", "hub1", "conn1",
            com.azure.core.util.Context.NONE);
    }
}
