
import com.azure.resourcemanager.horizondb.models.FirewallRuleProperties;

/**
 * Samples for FirewallRules CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/FirewallRules_CreateOrReplace.json
     */
    /**
     * Sample code: Create or replace a firewall rule on a HorizonDB pool.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        createOrReplaceAFirewallRuleOnAHorizonDBPool(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.firewallRules().define("example-firewall-rule-01")
            .withExistingPool("example-resource-group", "example-cluster", "example-pool-01")
            .withProperties(new FirewallRuleProperties().withStartIpAddress("10.0.0.1").withEndIpAddress("10.0.0.254")
                .withDescription("Allow access from the corporate network"))
            .create();
    }
}
