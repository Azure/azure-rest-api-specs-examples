
import com.azure.core.util.Context;

/** Samples for IPv6FirewallRules ListByServer. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/IPv6FirewallRuleList.json
     */
    /**
     * Sample code: List IPv6 Firewall Rules.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listIPv6FirewallRules(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getIPv6FirewallRules().listByServer("firewallrulecrudtest-12",
            "firewallrulecrudtest-6285", Context.NONE);
    }
}
