
/**
 * Samples for VirtualNetworkGateways ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VirtualNetworkGatewayList.
     * json
     */
    /**
     * Sample code: ListVirtualNetworkGatewaysinResourceGroup.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listVirtualNetworkGatewaysinResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGateways().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
