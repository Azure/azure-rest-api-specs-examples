
/**
 * Samples for FirewallRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/MongoClusters_FirewallRuleGet.json
     */
    /**
     * Sample code: Gets a firewall rule on a DocumentDB cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void getsAFirewallRuleOnADocumentDBClusterResource(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.firewallRules().getWithResponse("TestGroup", "myMongoCluster", "rule1",
            com.azure.core.util.Context.NONE);
    }
}
