
import com.azure.resourcemanager.sql.fluent.models.FirewallRuleInner;
import com.azure.resourcemanager.sql.models.FirewallRuleList;
import java.util.Arrays;

/**
 * Samples for FirewallRules Replace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/FirewallRuleReplace.json
     */
    /**
     * Sample code: Replace firewall rules.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void replaceFirewallRules(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getFirewallRules().replaceWithResponse("firewallrulecrudtest-12",
            "firewallrulecrudtest-6285",
            new FirewallRuleList().withValues(Arrays.asList(new FirewallRuleInner()
                .withName("firewallrulecrudtest-5370 ").withStartIpAddress("0.0.0.0").withEndIpAddress("100.0.0.0"))),
            com.azure.core.util.Context.NONE);
    }
}
