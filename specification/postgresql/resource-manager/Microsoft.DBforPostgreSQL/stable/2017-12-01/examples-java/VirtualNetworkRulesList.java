import com.azure.core.util.Context;

/** Samples for VirtualNetworkRules ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/VirtualNetworkRulesList.json
     */
    /**
     * Sample code: List virtual network rules.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void listVirtualNetworkRules(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.virtualNetworkRules().listByServer("TestGroup", "vnet-test-svr", Context.NONE);
    }
}
