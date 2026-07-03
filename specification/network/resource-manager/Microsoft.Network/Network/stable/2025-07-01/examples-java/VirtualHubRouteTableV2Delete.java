
/**
 * Samples for VirtualHubRouteTableV2S Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualHubRouteTableV2Delete.json
     */
    /**
     * Sample code: VirtualHubRouteTableV2Delete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void virtualHubRouteTableV2Delete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualHubRouteTableV2S().delete("rg1", "virtualHub1", "virtualHubRouteTable1a",
            com.azure.core.util.Context.NONE);
    }
}
