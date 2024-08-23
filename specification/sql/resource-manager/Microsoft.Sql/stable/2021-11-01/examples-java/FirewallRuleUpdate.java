
import com.azure.resourcemanager.sql.fluent.models.FirewallRuleInner;

/**
 * Samples for FirewallRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/FirewallRuleUpdate.json
     */
    /**
     * Sample code: Update a firewall rule max/min.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateAFirewallRuleMaxMin(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getFirewallRules().createOrUpdateWithResponse(
            "firewallrulecrudtest-12", "firewallrulecrudtest-6285", "firewallrulecrudtest-3927",
            new FirewallRuleInner().withStartIpAddress("0.0.0.1").withEndIpAddress("0.0.0.1"),
            com.azure.core.util.Context.NONE);
    }
}
