
/**
 * Samples for NetworkVirtualAppliances Reimage.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkVirtualApplianceSpecificReimage.json
     */
    /**
     * Sample code: Reimages Specific NetworkVirtualAppliance VMs in VM Scale Set.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void reimagesSpecificNetworkVirtualApplianceVMsInVMScaleSet(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkVirtualAppliances().reimage("rg1", "nva", null,
            com.azure.core.util.Context.NONE);
    }
}
