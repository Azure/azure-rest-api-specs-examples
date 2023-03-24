import com.azure.core.util.Context;

/** Samples for FirewallRules ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/FirewallRulesListByServer.json
     */
    /**
     * Sample code: List all firewall rules in a server.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void listAllFirewallRulesInAServer(
        com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.firewallRules().listByServer("TestGroup", "testserver", Context.NONE);
    }
}
