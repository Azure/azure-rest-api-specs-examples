
/**
 * Samples for ResourceNavigationLinks List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * VirtualNetworkGetResourceNavigationLinks.json
     */
    /**
     * Sample code: Get Resource Navigation Links.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getResourceNavigationLinks(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getResourceNavigationLinks().listWithResponse("rg1", "vnet",
            "subnet", com.azure.core.util.Context.NONE);
    }
}
