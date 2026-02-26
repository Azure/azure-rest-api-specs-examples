
/**
 * Samples for FirewallRules ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/FirewallRulesListByServer.json
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
