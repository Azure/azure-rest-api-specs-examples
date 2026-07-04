
/**
 * Samples for NetworkVirtualAppliances Restart.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkVirtualApplianceSpecificRestart.json
     */
    /**
     * Sample code: Restart Specific NetworkVirtualAppliance VMs in VM Scale Set.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void restartSpecificNetworkVirtualApplianceVMsInVMScaleSet(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkVirtualAppliances().restart("rg1", "nva", null,
            com.azure.core.util.Context.NONE);
    }
}
