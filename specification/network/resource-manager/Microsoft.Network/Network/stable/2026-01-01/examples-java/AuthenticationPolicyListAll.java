
/**
 * Samples for AuthenticationPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/AuthenticationPolicyListAll.json
     */
    /**
     * Sample code: Lists JWT validation and user sign-in authentication policies in a subscription.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listsJWTValidationAndUserSignInAuthenticationPoliciesInASubscription(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAuthenticationPolicies().list(com.azure.core.util.Context.NONE);
    }
}
