
/**
 * Samples for VirtualHubBgpConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VirtualHubBgpConnectionDelete
     * .json
     */
    /**
     * Sample code: VirtualHubRouteTableV2Delete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualHubRouteTableV2Delete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualHubBgpConnections().delete("rg1", "hub1", "conn1",
            com.azure.core.util.Context.NONE);
    }
}
