import com.azure.resourcemanager.datalakestore.models.FirewallRule;

/** Samples for FirewallRules Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/FirewallRules_Update.json
     */
    /**
     * Sample code: Updates the specified firewall rule.
     *
     * @param manager Entry point to DataLakeStoreManager.
     */
    public static void updatesTheSpecifiedFirewallRule(
        com.azure.resourcemanager.datalakestore.DataLakeStoreManager manager) {
        FirewallRule resource =
            manager
                .firewallRules()
                .getWithResponse("contosorg", "contosoadla", "test_rule", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withStartIpAddress("1.1.1.1").withEndIpAddress("2.2.2.2").apply();
    }
}
