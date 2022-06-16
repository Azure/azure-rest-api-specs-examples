import com.azure.core.util.Context;

/** Samples for VirtualNetworkRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/VirtualNetworkRulesGet.json
     */
    /**
     * Sample code: Gets a virtual network rule.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void getsAVirtualNetworkRule(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.virtualNetworkRules().getWithResponse("TestGroup", "vnet-test-svr", "vnet-firewall-rule", Context.NONE);
    }
}
