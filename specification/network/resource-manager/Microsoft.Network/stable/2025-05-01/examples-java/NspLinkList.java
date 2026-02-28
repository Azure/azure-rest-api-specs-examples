
/**
 * Samples for NetworkSecurityPerimeterLinks List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NspLinkList.json
     */
    /**
     * Sample code: NspLinkList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nspLinkList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeterLinks().list("rg1", "nsp1", null, null,
            com.azure.core.util.Context.NONE);
    }
}
