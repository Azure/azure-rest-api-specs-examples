
/**
 * Samples for FirewallPolicyIdpsSignaturesOverrides Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/FirewallPolicySignatureOverridesGet.json
     */
    /**
     * Sample code: get signature overrides.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getSignatureOverrides(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFirewallPolicyIdpsSignaturesOverrides().getWithResponse("rg1", "firewallPolicy",
            com.azure.core.util.Context.NONE);
    }
}
