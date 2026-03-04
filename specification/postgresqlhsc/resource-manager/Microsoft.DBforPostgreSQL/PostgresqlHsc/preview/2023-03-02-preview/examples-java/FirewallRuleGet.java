
/**
 * Samples for FirewallRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-03-02-preview/FirewallRuleGet.json
     */
    /**
     * Sample code: Get the firewall rule of the cluster.
     * 
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void getTheFirewallRuleOfTheCluster(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager.firewallRules().getWithResponse("TestGroup", "pgtestsvc4", "rule1", com.azure.core.util.Context.NONE);
    }
}
