
/**
 * Samples for FirewallPolicyKubeSelectorGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/FirewallPolicyKubeSelectorGroupGet.json
     */
    /**
     * Sample code: Get FirewallPolicyKubeSelectorGroup.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getFirewallPolicyKubeSelectorGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFirewallPolicyKubeSelectorGroups().getWithResponse("rg1", "firewallPolicy",
            "kubeSelectorGroup1", com.azure.core.util.Context.NONE);
    }
}
