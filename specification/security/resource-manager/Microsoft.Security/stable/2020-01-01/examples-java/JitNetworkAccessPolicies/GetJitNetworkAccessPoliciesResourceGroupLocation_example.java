
/**
 * Samples for JitNetworkAccessPolicies ListByResourceGroupAndRegion.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/
     * GetJitNetworkAccessPoliciesResourceGroupLocation_example.json
     */
    /**
     * Sample code: Get JIT network access policies on a resource group from a security data location.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getJITNetworkAccessPoliciesOnAResourceGroupFromASecurityDataLocation(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.jitNetworkAccessPolicies().listByResourceGroupAndRegion("myRg1", "westeurope",
            com.azure.core.util.Context.NONE);
    }
}
