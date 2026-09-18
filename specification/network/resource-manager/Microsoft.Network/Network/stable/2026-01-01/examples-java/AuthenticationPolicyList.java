
/**
 * Samples for AuthenticationPolicies ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/AuthenticationPolicyList.json
     */
    /**
     * Sample code: Lists authentication policies in a resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listsAuthenticationPoliciesInAResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAuthenticationPolicies().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
