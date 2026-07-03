
/**
 * Samples for FirewallPolicyIdpsSignaturesOverrides List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/FirewallPolicySignatureOverridesList.json
     */
    /**
     * Sample code: get signature overrides.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getSignatureOverrides(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFirewallPolicyIdpsSignaturesOverrides().listWithResponse("rg1", "firewallPolicy",
            com.azure.core.util.Context.NONE);
    }
}
