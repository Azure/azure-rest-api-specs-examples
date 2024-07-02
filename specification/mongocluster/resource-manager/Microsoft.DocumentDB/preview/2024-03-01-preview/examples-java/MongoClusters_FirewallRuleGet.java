
/**
 * Samples for FirewallRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/mongocluster/DocumentDB.MongoCluster.Management/examples/2024-03-01-preview/
     * MongoClusters_FirewallRuleGet.json
     */
    /**
     * Sample code: Gets a firewall rule on a Mongo Cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void
        getsAFirewallRuleOnAMongoClusterResource(com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.firewallRules().getWithResponse("TestGroup", "myMongoCluster", "rule1",
            com.azure.core.util.Context.NONE);
    }
}
