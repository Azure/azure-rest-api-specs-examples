
/**
 * Samples for PrivateLinkServices List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/PrivateLinkServiceListAll.
     * json
     */
    /**
     * Sample code: List all private list service.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllPrivateListService(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPrivateLinkServices().list(com.azure.core.util.Context.NONE);
    }
}
