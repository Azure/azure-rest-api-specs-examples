
import com.azure.core.util.Context;

/** Samples for OutboundFirewallRules Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/OutboundFirewallRuleDelete.json
     */
    /**
     * Sample code: Deletes a outbound firewall rule with a given name.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        deletesAOutboundFirewallRuleWithAGivenName(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getOutboundFirewallRules().delete("sqlcrudtest-7398",
            "sqlcrudtest-6661", "server.database.windows.net", Context.NONE);
    }
}
