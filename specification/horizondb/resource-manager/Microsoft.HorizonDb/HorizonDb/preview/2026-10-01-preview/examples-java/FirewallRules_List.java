
/**
 * Samples for FirewallRules List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/FirewallRules_List.json
     */
    /**
     * Sample code: List firewall rules in a HorizonDB pool.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void listFirewallRulesInAHorizonDBPool(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.firewallRules().list("example-resource-group", "example-cluster", "example-pool-01",
            com.azure.core.util.Context.NONE);
    }
}
