
import com.azure.resourcemanager.sql.fluent.models.VirtualNetworkRuleInner;

/**
 * Samples for VirtualNetworkRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/VirtualNetworkRulesCreateOrUpdate.
     * json
     */
    /**
     * Sample code: Create or update a virtual network rule.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateAVirtualNetworkRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getVirtualNetworkRules().createOrUpdate("Default", "vnet-test-svr",
            "vnet-firewall-rule",
            new VirtualNetworkRuleInner().withVirtualNetworkSubnetId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Network/virtualNetworks/testvnet/subnets/testsubnet")
                .withIgnoreMissingVnetServiceEndpoint(false),
            com.azure.core.util.Context.NONE);
    }
}
