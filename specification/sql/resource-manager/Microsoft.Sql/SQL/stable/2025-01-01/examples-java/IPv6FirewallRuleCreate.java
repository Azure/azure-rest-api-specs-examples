
import com.azure.resourcemanager.sql.fluent.models.IPv6FirewallRuleInner;

/**
 * Samples for IPv6FirewallRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/IPv6FirewallRuleCreate.json
     */
    /**
     * Sample code: Create an IPv6 firewall rule max/min.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createAnIPv6FirewallRuleMaxMin(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getIPv6FirewallRules().createOrUpdateWithResponse("firewallrulecrudtest-12",
            "firewallrulecrudtest-6285", "firewallrulecrudtest-5370",
            new IPv6FirewallRuleInner().withStartIPv6Address("0000:0000:0000:0000:0000:ffff:0000:0003")
                .withEndIPv6Address("0000:0000:0000:0000:0000:ffff:0000:0003"),
            com.azure.core.util.Context.NONE);
    }
}
