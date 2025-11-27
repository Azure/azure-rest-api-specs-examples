
/**
 * Samples for FirewallRules ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/
     * FirewallRulesListByServer.json
     */
    /**
     * Sample code: List information about all firewall rules in a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void listInformationAboutAllFirewallRulesInAServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.firewallRules().listByServer("exampleresourcegroup", "exampleserver", com.azure.core.util.Context.NONE);
    }
}
