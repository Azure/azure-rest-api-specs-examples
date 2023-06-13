/** Samples for FirewallRules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/FirewallRuleDelete.json
     */
    /**
     * Sample code: Delete the firewall rule of the cluster.
     *
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void deleteTheFirewallRuleOfTheCluster(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager.firewallRules().delete("TestGroup", "pgtestsvc4", "rule1", com.azure.core.util.Context.NONE);
    }
}
