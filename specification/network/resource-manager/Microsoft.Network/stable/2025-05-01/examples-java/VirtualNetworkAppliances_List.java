
/**
 * Samples for VirtualNetworkAppliances ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VirtualNetworkAppliances_List
     * .json
     */
    /**
     * Sample code: List virtual network appliances in resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listVirtualNetworkAppliancesInResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkAppliances().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
