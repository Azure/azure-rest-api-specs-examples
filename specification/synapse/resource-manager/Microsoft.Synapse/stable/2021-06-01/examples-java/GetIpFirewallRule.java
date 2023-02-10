/** Samples for IpFirewallRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetIpFirewallRule.json
     */
    /**
     * Sample code: Get IP firewall rule.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getIPFirewallRule(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .ipFirewallRules()
            .getWithResponse(
                "ExampleResourceGroup", "ExampleWorkspace", "ExampleIpFirewallRule", com.azure.core.util.Context.NONE);
    }
}
