
/**
 * Samples for AzureFirewalls GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/AzureFirewallGetWithAdditionalProperties.json
     */
    /**
     * Sample code: Get Azure Firewall With Additional Properties.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getAzureFirewallWithAdditionalProperties(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAzureFirewalls().getByResourceGroupWithResponse("rg1", "azurefirewall",
            com.azure.core.util.Context.NONE);
    }
}
