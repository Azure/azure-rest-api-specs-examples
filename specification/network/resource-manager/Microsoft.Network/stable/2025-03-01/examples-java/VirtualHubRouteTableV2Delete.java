
/**
 * Samples for VirtualHubRouteTableV2S Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VirtualHubRouteTableV2Delete.
     * json
     */
    /**
     * Sample code: VirtualHubRouteTableV2Delete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualHubRouteTableV2Delete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualHubRouteTableV2S().delete("rg1", "virtualHub1",
            "virtualHubRouteTable1a", com.azure.core.util.Context.NONE);
    }
}
