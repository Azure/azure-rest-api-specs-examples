/** Samples for IpFirewallRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateIpFirewallRule.json
     */
    /**
     * Sample code: Create an IP firewall rule.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void createAnIPFirewallRule(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .ipFirewallRules()
            .define("ExampleIpFirewallRule")
            .withExistingWorkspace("ExampleResourceGroup", "ExampleWorkspace")
            .withEndIpAddress("10.0.0.254")
            .withStartIpAddress("10.0.0.0")
            .create();
    }
}
