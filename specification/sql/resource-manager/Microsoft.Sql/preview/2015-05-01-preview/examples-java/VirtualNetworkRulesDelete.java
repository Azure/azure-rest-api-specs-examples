import com.azure.core.util.Context;

/** Samples for VirtualNetworkRules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2015-05-01-preview/examples/VirtualNetworkRulesDelete.json
     */
    /**
     * Sample code: Delete a virtual network rule.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAVirtualNetworkRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getVirtualNetworkRules()
            .delete("Default", "vnet-test-svr", "vnet-firewall-rule", Context.NONE);
    }
}
