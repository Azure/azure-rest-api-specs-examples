
/**
 * Samples for JitNetworkAccessPolicies ListByRegion.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/
     * GetJitNetworkAccessPoliciesSubscriptionLocation_example.json
     */
    /**
     * Sample code: Get JIT network access policies on a subscription from a security data location.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getJITNetworkAccessPoliciesOnASubscriptionFromASecurityDataLocation(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.jitNetworkAccessPolicies().listByRegion("westeurope", com.azure.core.util.Context.NONE);
    }
}
