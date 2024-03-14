
/**
 * Samples for JitNetworkAccessPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/
     * DeleteJitNetworkAccessPolicy_example.json
     */
    /**
     * Sample code: Delete a JIT network access policy.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void deleteAJITNetworkAccessPolicy(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.jitNetworkAccessPolicies().deleteWithResponse("myRg1", "westeurope", "default",
            com.azure.core.util.Context.NONE);
    }
}
