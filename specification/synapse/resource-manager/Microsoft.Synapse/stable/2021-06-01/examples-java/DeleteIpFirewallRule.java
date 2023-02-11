/** Samples for IpFirewallRules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DeleteIpFirewallRule.json
     */
    /**
     * Sample code: Delete an IP firewall rule from a workspace.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void deleteAnIPFirewallRuleFromAWorkspace(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .ipFirewallRules()
            .delete(
                "ExampleResourceGroup", "ExampleWorkspace", "ExampleIpFirewallRule", com.azure.core.util.Context.NONE);
    }
}
