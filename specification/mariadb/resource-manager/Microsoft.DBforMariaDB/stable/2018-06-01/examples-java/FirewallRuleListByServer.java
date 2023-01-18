/** Samples for FirewallRules ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/FirewallRuleListByServer.json
     */
    /**
     * Sample code: FirewallRuleList.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void firewallRuleList(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager.firewallRules().listByServer("TestGroup", "testserver", com.azure.core.util.Context.NONE);
    }
}
