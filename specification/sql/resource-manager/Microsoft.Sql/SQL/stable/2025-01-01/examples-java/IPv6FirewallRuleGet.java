
/**
 * Samples for IPv6FirewallRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/IPv6FirewallRuleGet.json
     */
    /**
     * Sample code: Get IPv6 Firewall Rule.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getIPv6FirewallRule(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getIPv6FirewallRules().getWithResponse("firewallrulecrudtest-12",
            "firewallrulecrudtest-6285", "firewallrulecrudtest-2304", com.azure.core.util.Context.NONE);
    }
}
