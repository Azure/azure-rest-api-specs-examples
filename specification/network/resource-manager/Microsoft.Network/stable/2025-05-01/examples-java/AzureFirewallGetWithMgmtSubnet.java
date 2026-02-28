
/**
 * Samples for AzureFirewalls GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * AzureFirewallGetWithMgmtSubnet.json
     */
    /**
     * Sample code: Get Azure Firewall With management subnet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAzureFirewallWithManagementSubnet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getAzureFirewalls().getByResourceGroupWithResponse("rg1",
            "azurefirewall", com.azure.core.util.Context.NONE);
    }
}
