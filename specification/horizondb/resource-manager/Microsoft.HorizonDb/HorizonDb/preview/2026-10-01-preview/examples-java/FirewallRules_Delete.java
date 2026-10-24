
/**
 * Samples for FirewallRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/FirewallRules_Delete.json
     */
    /**
     * Sample code: Delete a firewall rule from a HorizonDB pool.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        deleteAFirewallRuleFromAHorizonDBPool(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.firewallRules().delete("example-resource-group", "example-cluster", "example-pool-01",
            "example-firewall-rule-01", com.azure.core.util.Context.NONE);
    }
}
