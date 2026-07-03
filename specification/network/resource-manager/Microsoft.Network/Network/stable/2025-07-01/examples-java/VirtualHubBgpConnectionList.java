
/**
 * Samples for VirtualHubBgpConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualHubBgpConnectionList.json
     */
    /**
     * Sample code: VirtualHubRouteTableV2List.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void virtualHubRouteTableV2List(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualHubBgpConnections().list("rg1", "hub1", com.azure.core.util.Context.NONE);
    }
}
