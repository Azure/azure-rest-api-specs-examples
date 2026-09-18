
/**
 * Samples for AzureFirewalls GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/AzureFirewallGetWithMgmtSubnet.json
     */
    /**
     * Sample code: Get Azure Firewall With management subnet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getAzureFirewallWithManagementSubnet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAzureFirewalls().getByResourceGroupWithResponse("rg1", "azurefirewall",
            com.azure.core.util.Context.NONE);
    }
}
