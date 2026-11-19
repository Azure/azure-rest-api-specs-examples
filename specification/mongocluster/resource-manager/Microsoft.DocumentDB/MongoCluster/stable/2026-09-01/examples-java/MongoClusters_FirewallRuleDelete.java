
/**
 * Samples for FirewallRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/MongoClusters_FirewallRuleDelete.json
     */
    /**
     * Sample code: Deletes a firewall rule on a DocumentDB cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void deletesAFirewallRuleOnADocumentDBClusterResource(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.firewallRules().delete("TestGroup", "myMongoCluster", "rule1", com.azure.core.util.Context.NONE);
    }
}
