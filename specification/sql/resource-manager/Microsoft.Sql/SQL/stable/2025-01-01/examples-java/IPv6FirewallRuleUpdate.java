
import com.azure.resourcemanager.sql.fluent.models.IPv6FirewallRuleInner;

/**
 * Samples for IPv6FirewallRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/IPv6FirewallRuleUpdate.json
     */
    /**
     * Sample code: Update an IPv6 firewall rule max/min.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateAnIPv6FirewallRuleMaxMin(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getIPv6FirewallRules().createOrUpdateWithResponse("firewallrulecrudtest-12",
            "firewallrulecrudtest-6285", "firewallrulecrudtest-3927",
            new IPv6FirewallRuleInner().withStartIPv6Address("0000:0000:0000:0000:0000:ffff:0000:0001")
                .withEndIPv6Address("0000:0000:0000:0000:0000:ffff:0000:0001"),
            com.azure.core.util.Context.NONE);
    }
}
