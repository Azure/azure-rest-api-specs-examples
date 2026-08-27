
/**
 * Samples for HorizonDbFirewallRules List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/FirewallRules_List.json
     */
    /**
     * Sample code: List HorizonDB firewall rules in a pool.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void listHorizonDBFirewallRulesInAPool(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.horizonDbFirewallRules().list("exampleresourcegroup", "examplecluster", "examplepool",
            com.azure.core.util.Context.NONE);
    }
}
