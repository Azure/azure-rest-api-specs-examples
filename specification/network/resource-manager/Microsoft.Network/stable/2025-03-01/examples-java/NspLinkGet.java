
/**
 * Samples for NetworkSecurityPerimeterLinks Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NspLinkGet.json
     */
    /**
     * Sample code: NspLinksGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nspLinksGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeterLinks().getWithResponse("rg1", "nsp1",
            "link1", com.azure.core.util.Context.NONE);
    }
}
