
/**
 * Samples for VirtualHubBgpConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualHubBgpConnectionDelete.json
     */
    /**
     * Sample code: VirtualHubRouteTableV2Delete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void virtualHubRouteTableV2Delete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualHubBgpConnections().delete("rg1", "hub1", "conn1",
            com.azure.core.util.Context.NONE);
    }
}
