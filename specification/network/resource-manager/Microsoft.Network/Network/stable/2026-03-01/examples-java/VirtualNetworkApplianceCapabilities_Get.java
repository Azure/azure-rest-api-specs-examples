
/**
 * Samples for VirtualNetworkApplianceCapabilities Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/VirtualNetworkApplianceCapabilities_Get.json
     */
    /**
     * Sample code: Get virtual network appliance capability.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getVirtualNetworkApplianceCapability(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkApplianceCapabilities().getWithResponse("rg1", "test-vna",
            "pl-fastpath", com.azure.core.util.Context.NONE);
    }
}
