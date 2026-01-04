
/**
 * Samples for PrivateLinkServices ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/PrivateLinkServiceList.json
     */
    /**
     * Sample code: List private link service in resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listPrivateLinkServiceInResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPrivateLinkServices().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
