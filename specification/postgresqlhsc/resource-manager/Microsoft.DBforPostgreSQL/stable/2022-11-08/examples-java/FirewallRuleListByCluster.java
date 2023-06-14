/** Samples for FirewallRules ListByCluster. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/FirewallRuleListByCluster.json
     */
    /**
     * Sample code: List firewall rules of the cluster.
     *
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void listFirewallRulesOfTheCluster(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager.firewallRules().listByCluster("TestGroup", "pgtestsvc4", com.azure.core.util.Context.NONE);
    }
}
