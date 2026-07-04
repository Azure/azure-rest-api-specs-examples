
/**
 * Samples for NetworkVirtualAppliances GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkVirtualApplianceGet.json
     */
    /**
     * Sample code: Get NetworkVirtualAppliance.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getNetworkVirtualAppliance(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkVirtualAppliances().getByResourceGroupWithResponse("rg1", "nva", null,
            com.azure.core.util.Context.NONE);
    }
}
