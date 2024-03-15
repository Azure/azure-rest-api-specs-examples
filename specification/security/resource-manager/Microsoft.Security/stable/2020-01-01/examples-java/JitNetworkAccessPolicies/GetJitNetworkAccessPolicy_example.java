
/**
 * Samples for JitNetworkAccessPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/
     * GetJitNetworkAccessPolicy_example.json
     */
    /**
     * Sample code: Get JIT network access policy.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getJITNetworkAccessPolicy(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.jitNetworkAccessPolicies().getWithResponse("myRg1", "westeurope", "default",
            com.azure.core.util.Context.NONE);
    }
}
