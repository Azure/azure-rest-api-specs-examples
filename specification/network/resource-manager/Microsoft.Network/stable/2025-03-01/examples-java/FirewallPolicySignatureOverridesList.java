
/**
 * Samples for FirewallPolicyIdpsSignaturesOverrides List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * FirewallPolicySignatureOverridesList.json
     */
    /**
     * Sample code: get signature overrides.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getSignatureOverrides(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getFirewallPolicyIdpsSignaturesOverrides().listWithResponse("rg1",
            "firewallPolicy", com.azure.core.util.Context.NONE);
    }
}
