
/**
 * Samples for NetworkVirtualAppliances GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/NetworkVirtualApplianceVhubDualStackGet.json
     */
    /**
     * Sample code: Get NetworkVirtualAppliance in Virtual Hub with Dual Stack.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getNetworkVirtualApplianceInVirtualHubWithDualStack(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkVirtualAppliances().getByResourceGroupWithResponse("rg1", "nva", null,
            com.azure.core.util.Context.NONE);
    }
}
