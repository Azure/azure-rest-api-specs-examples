/** Samples for VirtualNetworkRules ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/VirtualNetworkRulesList.json
     */
    /**
     * Sample code: List virtual network rules.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void listVirtualNetworkRules(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager.virtualNetworkRules().listByServer("TestGroup", "vnet-test-svr", com.azure.core.util.Context.NONE);
    }
}
