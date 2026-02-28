
/**
 * Samples for NetworkVirtualAppliances Restart.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkVirtualApplianceEmptyRestart.json
     */
    /**
     * Sample code: Restart All NetworkVirtualAppliance VMs in VM Scale Set.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        restartAllNetworkVirtualApplianceVMsInVMScaleSet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkVirtualAppliances().restart("rg1", "nva", null,
            com.azure.core.util.Context.NONE);
    }
}
