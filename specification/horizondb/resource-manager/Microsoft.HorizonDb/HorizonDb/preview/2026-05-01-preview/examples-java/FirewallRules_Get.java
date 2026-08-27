
/**
 * Samples for HorizonDbFirewallRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/FirewallRules_Get.json
     */
    /**
     * Sample code: Get a HorizonDB firewall rule.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void getAHorizonDBFirewallRule(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.horizonDbFirewallRules().getWithResponse("exampleresourcegroup", "examplecluster", "examplepool",
            "examplefirewallrule", com.azure.core.util.Context.NONE);
    }
}
