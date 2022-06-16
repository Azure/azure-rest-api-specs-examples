import com.azure.core.util.Context;

/** Samples for VirtualNetworkRules ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/VirtualNetworkRulesList.json
     */
    /**
     * Sample code: List virtual network rules.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void listVirtualNetworkRules(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.virtualNetworkRules().listByServer("TestGroup", "vnet-test-svr", Context.NONE);
    }
}
