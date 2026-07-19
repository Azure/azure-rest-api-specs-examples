
import com.azure.resourcemanager.sql.fluent.models.FirewallRuleInner;

/**
 * Samples for FirewallRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/FirewallRuleCreate.json
     */
    /**
     * Sample code: Create a firewall rule max/min.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createAFirewallRuleMaxMin(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getFirewallRules().createOrUpdateWithResponse("firewallrulecrudtest-12",
            "firewallrulecrudtest-6285", "firewallrulecrudtest-5370",
            new FirewallRuleInner().withStartIpAddress("0.0.0.3").withEndIpAddress("0.0.0.3"),
            com.azure.core.util.Context.NONE);
    }
}
