
/**
 * Samples for VirtualApplianceSites List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkVirtualApplianceSiteList.json
     */
    /**
     * Sample code: List all Network Virtual Appliance sites for a given Network Virtual Appliance.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllNetworkVirtualApplianceSitesForAGivenNetworkVirtualAppliance(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualApplianceSites().list("rg1", "nva",
            com.azure.core.util.Context.NONE);
    }
}
