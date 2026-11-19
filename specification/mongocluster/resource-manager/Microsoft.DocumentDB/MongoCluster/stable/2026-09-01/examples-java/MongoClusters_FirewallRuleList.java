
/**
 * Samples for FirewallRules ListByMongoCluster.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/MongoClusters_FirewallRuleList.json
     */
    /**
     * Sample code: List the firewall rules on a DocumentDB cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void listTheFirewallRulesOnADocumentDBClusterResource(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.firewallRules().listByMongoCluster("TestGroup", "myMongoCluster", com.azure.core.util.Context.NONE);
    }
}
