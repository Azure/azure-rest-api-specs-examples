
/**
 * Samples for AzureFirewalls GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/AzureFirewallGet.json
     */
    /**
     * Sample code: Get Azure Firewall.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAzureFirewall(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getAzureFirewalls().getByResourceGroupWithResponse("rg1",
            "azurefirewall", com.azure.core.util.Context.NONE);
    }
}
