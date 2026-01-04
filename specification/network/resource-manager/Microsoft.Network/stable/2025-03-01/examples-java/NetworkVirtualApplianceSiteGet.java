
/**
 * Samples for VirtualApplianceSites Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * NetworkVirtualApplianceSiteGet.json
     */
    /**
     * Sample code: GetNetwork Virtual Appliance Site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getNetworkVirtualApplianceSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualApplianceSites().getWithResponse("rg1", "nva", "site1",
            com.azure.core.util.Context.NONE);
    }
}
