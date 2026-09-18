
/**
 * Samples for AuthenticationPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/AuthenticationPolicyDelete.json
     */
    /**
     * Sample code: Deletes an authentication policy within a resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        deletesAnAuthenticationPolicyWithinAResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAuthenticationPolicies().deleteWithResponse("rg1", "authPolicy1",
            com.azure.core.util.Context.NONE);
    }
}
