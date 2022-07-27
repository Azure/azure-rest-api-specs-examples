import com.azure.core.util.Context;

/** Samples for VirtualNetworkRules ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2015-05-01-preview/examples/VirtualNetworkRulesList.json
     */
    /**
     * Sample code: List virtual network rules.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listVirtualNetworkRules(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getVirtualNetworkRules()
            .listByServer("Default", "vnet-test-svr", Context.NONE);
    }
}
