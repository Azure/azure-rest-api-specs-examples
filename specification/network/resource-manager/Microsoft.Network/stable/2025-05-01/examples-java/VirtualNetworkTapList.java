
/**
 * Samples for VirtualNetworkTaps ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VirtualNetworkTapList.json
     */
    /**
     * Sample code: List virtual network taps in resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listVirtualNetworkTapsInResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkTaps().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
