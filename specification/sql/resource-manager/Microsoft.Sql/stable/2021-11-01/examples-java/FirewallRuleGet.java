
/**
 * Samples for FirewallRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/FirewallRuleGet.json
     */
    /**
     * Sample code: Get Firewall Rule.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getFirewallRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getFirewallRules().getWithResponse("firewallrulecrudtest-12",
            "firewallrulecrudtest-6285", "firewallrulecrudtest-2304", com.azure.core.util.Context.NONE);
    }
}
