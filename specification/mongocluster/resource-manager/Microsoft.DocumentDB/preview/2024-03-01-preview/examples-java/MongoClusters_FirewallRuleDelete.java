
/**
 * Samples for FirewallRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/mongocluster/DocumentDB.MongoCluster.Management/examples/2024-03-01-preview/
     * MongoClusters_FirewallRuleDelete.json
     */
    /**
     * Sample code: Deletes a firewall rule on a Mongo Cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void deletesAFirewallRuleOnAMongoClusterResource(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.firewallRules().delete("TestGroup", "myMongoCluster", "rule1", com.azure.core.util.Context.NONE);
    }
}
