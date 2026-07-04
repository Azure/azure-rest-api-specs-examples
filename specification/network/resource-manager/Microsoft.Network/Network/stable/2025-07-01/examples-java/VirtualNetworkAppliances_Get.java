
/**
 * Samples for VirtualNetworkAppliances GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkAppliances_Get.json
     */
    /**
     * Sample code: Get virtual network appliance.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getVirtualNetworkAppliance(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkAppliances().getByResourceGroupWithResponse("rg1", "test-vna",
            com.azure.core.util.Context.NONE);
    }
}
