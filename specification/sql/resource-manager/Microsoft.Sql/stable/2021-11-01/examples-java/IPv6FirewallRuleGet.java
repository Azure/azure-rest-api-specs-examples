
import com.azure.core.util.Context;

/** Samples for IPv6FirewallRules Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/IPv6FirewallRuleGet.json
     */
    /**
     * Sample code: Get IPv6 Firewall Rule.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getIPv6FirewallRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getIPv6FirewallRules().getWithResponse("firewallrulecrudtest-12",
            "firewallrulecrudtest-6285", "firewallrulecrudtest-2304", Context.NONE);
    }
}
