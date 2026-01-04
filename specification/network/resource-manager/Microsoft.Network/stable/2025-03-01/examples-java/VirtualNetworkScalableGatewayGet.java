
/**
 * Samples for VirtualNetworkGateways GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * VirtualNetworkScalableGatewayGet.json
     */
    /**
     * Sample code: GetVirtualNetworkScalableGateway.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVirtualNetworkScalableGateway(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGateways().getByResourceGroupWithResponse("rg1",
            "ergw", com.azure.core.util.Context.NONE);
    }
}
