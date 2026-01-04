
/**
 * Samples for PrivateLinkServices Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/PrivateLinkServiceDelete.json
     */
    /**
     * Sample code: Delete private link service.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deletePrivateLinkService(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPrivateLinkServices().delete("rg1", "testPls",
            com.azure.core.util.Context.NONE);
    }
}
