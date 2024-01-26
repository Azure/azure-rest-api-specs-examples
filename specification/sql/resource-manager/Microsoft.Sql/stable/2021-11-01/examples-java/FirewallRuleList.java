
import com.azure.core.util.Context;

/** Samples for FirewallRules ListByServer. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/FirewallRuleList.json
     */
    /**
     * Sample code: List Firewall Rules.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listFirewallRules(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getFirewallRules().listByServer("firewallrulecrudtest-12",
            "firewallrulecrudtest-6285", Context.NONE);
    }
}
