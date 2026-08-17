
/**
 * Samples for NetworkVirtualAppliances ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/NetworkVirtualApplianceListByResourceGroup.json
     */
    /**
     * Sample code: List all Network Virtual Appliance for a given resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listAllNetworkVirtualApplianceForAGivenResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkVirtualAppliances().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
