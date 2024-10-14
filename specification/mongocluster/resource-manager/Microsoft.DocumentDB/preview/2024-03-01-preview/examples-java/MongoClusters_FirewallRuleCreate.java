
import com.azure.resourcemanager.mongocluster.models.FirewallRuleProperties;

/**
 * Samples for FirewallRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-01-preview/MongoClusters_FirewallRuleCreate.json
     */
    /**
     * Sample code: Creates a firewall rule on a Mongo Cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void createsAFirewallRuleOnAMongoClusterResource(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.firewallRules().define("rule1").withExistingMongoCluster("TestGroup", "myMongoCluster")
            .withProperties(
                new FirewallRuleProperties().withStartIpAddress("0.0.0.0").withEndIpAddress("255.255.255.255"))
            .create();
    }
}
