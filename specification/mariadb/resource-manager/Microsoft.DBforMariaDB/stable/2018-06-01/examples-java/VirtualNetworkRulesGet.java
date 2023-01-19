/** Samples for VirtualNetworkRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/VirtualNetworkRulesGet.json
     */
    /**
     * Sample code: Gets a virtual network rule.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void getsAVirtualNetworkRule(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager
            .virtualNetworkRules()
            .getWithResponse("TestGroup", "vnet-test-svr", "vnet-firewall-rule", com.azure.core.util.Context.NONE);
    }
}
