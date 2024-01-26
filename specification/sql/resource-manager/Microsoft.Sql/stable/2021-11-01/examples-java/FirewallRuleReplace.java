
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.FirewallRuleInner;
import com.azure.resourcemanager.sql.models.FirewallRuleList;
import java.util.Arrays;

/** Samples for FirewallRules Replace. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/FirewallRuleReplace.json
     */
    /**
     * Sample code: Replace firewall rules.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void replaceFirewallRules(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getFirewallRules().replaceWithResponse("firewallrulecrudtest-12",
            "firewallrulecrudtest-6285",
            new FirewallRuleList().withValues(Arrays.asList(new FirewallRuleInner()
                .withName("firewallrulecrudtest-5370 ").withStartIpAddress("0.0.0.0").withEndIpAddress("100.0.0.0"))),
            Context.NONE);
    }
}
