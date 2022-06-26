import com.azure.core.util.Context;

/** Samples for AzureFirewalls GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/AzureFirewallGet.json
     */
    /**
     * Sample code: Get Azure Firewall.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAzureFirewall(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getAzureFirewalls()
            .getByResourceGroupWithResponse("rg1", "azurefirewall", Context.NONE);
    }
}
