
/**
 * Samples for VirtualNetworkApplianceCapabilities List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/VirtualNetworkApplianceCapabilities_List.json
     */
    /**
     * Sample code: List virtual network appliance capabilities.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listVirtualNetworkApplianceCapabilities(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkApplianceCapabilities().list("rg1", "test-vna",
            com.azure.core.util.Context.NONE);
    }
}
