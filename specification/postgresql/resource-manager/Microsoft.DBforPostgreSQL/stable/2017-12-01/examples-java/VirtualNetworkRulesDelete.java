import com.azure.core.util.Context;

/** Samples for VirtualNetworkRules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/VirtualNetworkRulesDelete.json
     */
    /**
     * Sample code: Delete a virtual network rule.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void deleteAVirtualNetworkRule(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.virtualNetworkRules().delete("TestGroup", "vnet-test-svr", "vnet-firewall-rule", Context.NONE);
    }
}
