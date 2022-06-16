import com.azure.core.util.Context;

/** Samples for IpFirewallRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetIpFirewallRule.json
     */
    /**
     * Sample code: Create an IP firewall rule.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void createAnIPFirewallRule(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .ipFirewallRules()
            .getWithResponse("ExampleResourceGroup", "ExampleWorkspace", "ExampleIpFirewallRule", Context.NONE);
    }
}
