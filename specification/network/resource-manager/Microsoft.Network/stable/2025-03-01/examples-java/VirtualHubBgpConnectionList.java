
/**
 * Samples for VirtualHubBgpConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VirtualHubBgpConnectionList.
     * json
     */
    /**
     * Sample code: VirtualHubRouteTableV2List.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualHubRouteTableV2List(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualHubBgpConnections().list("rg1", "hub1",
            com.azure.core.util.Context.NONE);
    }
}
