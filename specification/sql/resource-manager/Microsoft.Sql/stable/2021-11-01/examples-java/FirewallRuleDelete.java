
import com.azure.core.util.Context;

/** Samples for FirewallRules Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/FirewallRuleDelete.json
     */
    /**
     * Sample code: Delete a firewall rule.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAFirewallRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getFirewallRules().deleteWithResponse("firewallrulecrudtest-9886",
            "firewallrulecrudtest-2368", "firewallrulecrudtest-7011", Context.NONE);
    }
}
