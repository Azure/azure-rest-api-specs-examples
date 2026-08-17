
/**
 * Samples for FirewallPolicyKubeSelectorGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/FirewallPolicyKubeSelectorGroupList.json
     */
    /**
     * Sample code: List all FirewallPolicyKubeSelectorGroups for a given FirewallPolicy.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllFirewallPolicyKubeSelectorGroupsForAGivenFirewallPolicy(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFirewallPolicyKubeSelectorGroups().list("rg1", "firewallPolicy",
            com.azure.core.util.Context.NONE);
    }
}
