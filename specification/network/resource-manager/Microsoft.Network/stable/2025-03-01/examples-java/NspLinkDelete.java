
/**
 * Samples for NetworkSecurityPerimeterLinks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NspLinkDelete.json
     */
    /**
     * Sample code: NspLinkDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nspLinkDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeterLinks().delete("rg1", "nsp1", "link1",
            com.azure.core.util.Context.NONE);
    }
}
