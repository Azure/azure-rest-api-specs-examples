
/**
 * Samples for AuthenticationPolicies GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/AuthenticationPolicyGet.json
     */
    /**
     * Sample code: Gets a JWT validation authentication policy within a resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getsAJWTValidationAuthenticationPolicyWithinAResourceGroup(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAuthenticationPolicies().getByResourceGroupWithResponse("rg1", "authPolicy1",
            com.azure.core.util.Context.NONE);
    }
}
