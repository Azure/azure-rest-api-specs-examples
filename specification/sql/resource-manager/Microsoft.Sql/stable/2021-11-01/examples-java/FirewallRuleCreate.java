
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.FirewallRuleInner;

/** Samples for FirewallRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/FirewallRuleCreate.json
     */
    /**
     * Sample code: Create a firewall rule max/min.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAFirewallRuleMaxMin(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getFirewallRules().createOrUpdateWithResponse(
            "firewallrulecrudtest-12", "firewallrulecrudtest-6285", "firewallrulecrudtest-5370",
            new FirewallRuleInner().withStartIpAddress("0.0.0.3").withEndIpAddress("0.0.0.3"), Context.NONE);
    }
}
