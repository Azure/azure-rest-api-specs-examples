
import com.azure.core.util.Context;

/** Samples for IPv6FirewallRules Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/IPv6FirewallRuleDelete.json
     */
    /**
     * Sample code: Delete an IPv6 firewall rule.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAnIPv6FirewallRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getIPv6FirewallRules().deleteWithResponse(
            "firewallrulecrudtest-9886", "firewallrulecrudtest-2368", "firewallrulecrudtest-7011", Context.NONE);
    }
}
