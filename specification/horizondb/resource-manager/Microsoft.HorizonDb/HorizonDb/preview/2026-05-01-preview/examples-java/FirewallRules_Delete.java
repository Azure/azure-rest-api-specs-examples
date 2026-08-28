
/**
 * Samples for HorizonDbFirewallRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/FirewallRules_Delete.json
     */
    /**
     * Sample code: Delete a HorizonDB firewall rule.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void deleteAHorizonDBFirewallRule(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.horizonDbFirewallRules().delete("exampleresourcegroup", "examplecluster", "examplepool",
            "examplefirewallrule", com.azure.core.util.Context.NONE);
    }
}
