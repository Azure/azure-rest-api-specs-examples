
/**
 * Samples for VirtualNetworkRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/
     * VirtualNetworkRulesGet.json
     */
    /**
     * Sample code: Gets a virtual network rule.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getsAVirtualNetworkRule(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.virtualNetworkRules().getWithResponse("TestGroup", "vnet-test-svr", "vnet-firewall-rule",
            com.azure.core.util.Context.NONE);
    }
}
