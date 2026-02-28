
/**
 * Samples for AzureFirewalls GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/AzureFirewallGetWithZones.
     * json
     */
    /**
     * Sample code: Get Azure Firewall With Zones.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAzureFirewallWithZones(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getAzureFirewalls().getByResourceGroupWithResponse("rg1",
            "azurefirewall", com.azure.core.util.Context.NONE);
    }
}
