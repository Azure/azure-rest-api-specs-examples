import com.azure.core.util.Context;

/** Samples for JitNetworkAccessPolicies List. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/GetJitNetworkAccessPoliciesSubscription_example.json
     */
    /**
     * Sample code: Get JIT network access policies on a subscription.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getJITNetworkAccessPoliciesOnASubscription(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.jitNetworkAccessPolicies().list(Context.NONE);
    }
}
