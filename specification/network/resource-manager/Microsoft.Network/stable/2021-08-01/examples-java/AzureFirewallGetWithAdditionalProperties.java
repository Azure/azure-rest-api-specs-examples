import com.azure.core.util.Context;

/** Samples for AzureFirewalls GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/AzureFirewallGetWithAdditionalProperties.json
     */
    /**
     * Sample code: Get Azure Firewall With Additional Properties.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAzureFirewallWithAdditionalProperties(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getAzureFirewalls()
            .getByResourceGroupWithResponse("rg1", "azurefirewall", Context.NONE);
    }
}
