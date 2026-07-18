
import com.azure.resourcemanager.sql.fluent.models.FirewallRuleInner;

/**
 * Samples for FirewallRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/FirewallRuleUpdate.json
     */
    /**
     * Sample code: Update a firewall rule max/min.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateAFirewallRuleMaxMin(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getFirewallRules().createOrUpdateWithResponse("firewallrulecrudtest-12",
            "firewallrulecrudtest-6285", "firewallrulecrudtest-3927",
            new FirewallRuleInner().withStartIpAddress("0.0.0.1").withEndIpAddress("0.0.0.1"),
            com.azure.core.util.Context.NONE);
    }
}
