
/**
 * Samples for NetworkVirtualAppliances GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/NetworkVirtualApplianceVnetDualStackGet.json
     */
    /**
     * Sample code: Get NetworkVirtualAppliance in VNet with Dual Stack.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getNetworkVirtualApplianceInVNetWithDualStack(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkVirtualAppliances().getByResourceGroupWithResponse("rg1", "nva", null,
            com.azure.core.util.Context.NONE);
    }
}
