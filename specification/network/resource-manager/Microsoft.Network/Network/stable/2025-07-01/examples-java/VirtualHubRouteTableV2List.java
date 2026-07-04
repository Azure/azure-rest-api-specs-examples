
/**
 * Samples for VirtualHubRouteTableV2S List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualHubRouteTableV2List.json
     */
    /**
     * Sample code: VirtualHubRouteTableV2List.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void virtualHubRouteTableV2List(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualHubRouteTableV2S().list("rg1", "virtualHub1",
            com.azure.core.util.Context.NONE);
    }
}
