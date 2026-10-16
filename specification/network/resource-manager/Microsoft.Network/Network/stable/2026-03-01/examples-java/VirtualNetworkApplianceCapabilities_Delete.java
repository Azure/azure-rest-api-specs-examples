
/**
 * Samples for VirtualNetworkApplianceCapabilities Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/VirtualNetworkApplianceCapabilities_Delete.json
     */
    /**
     * Sample code: Delete virtual network appliance capability.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        deleteVirtualNetworkApplianceCapability(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkApplianceCapabilities().delete("rg1", "test-vna", "pl-fastpath",
            com.azure.core.util.Context.NONE);
    }
}
