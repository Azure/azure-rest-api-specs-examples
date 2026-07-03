
/**
 * Samples for VirtualApplianceSites List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkVirtualApplianceSiteList.json
     */
    /**
     * Sample code: List all Network Virtual Appliance sites for a given Network Virtual Appliance.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllNetworkVirtualApplianceSitesForAGivenNetworkVirtualAppliance(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualApplianceSites().list("rg1", "nva", com.azure.core.util.Context.NONE);
    }
}
