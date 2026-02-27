
/**
 * Samples for NetworkVirtualAppliances Reimage.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkVirtualApplianceSpecificReimage.json
     */
    /**
     * Sample code: Reimages Specific NetworkVirtualAppliance VMs in VM Scale Set.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        reimagesSpecificNetworkVirtualApplianceVMsInVMScaleSet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkVirtualAppliances().reimage("rg1", "nva", null,
            com.azure.core.util.Context.NONE);
    }
}
