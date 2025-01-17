
/**
 * Samples for VirtualNetworkRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/
     * VirtualNetworkRulesDelete.json
     */
    /**
     * Sample code: Delete a virtual network rule.
     * 
     * @param manager Entry point to MariaDBManager.
     */
    public static void deleteAVirtualNetworkRule(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager.virtualNetworkRules().delete("TestGroup", "vnet-test-svr", "vnet-firewall-rule",
            com.azure.core.util.Context.NONE);
    }
}
