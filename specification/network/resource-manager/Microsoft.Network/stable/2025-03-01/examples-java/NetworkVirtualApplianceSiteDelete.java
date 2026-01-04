
/**
 * Samples for VirtualApplianceSites Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * NetworkVirtualApplianceSiteDelete.json
     */
    /**
     * Sample code: Delete Network Virtual Appliance Site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteNetworkVirtualApplianceSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualApplianceSites().delete("rg1", "nva", "site1",
            com.azure.core.util.Context.NONE);
    }
}
